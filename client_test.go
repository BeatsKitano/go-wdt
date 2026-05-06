package wdt

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestSignDirect(t *testing.T) {
	sign := SignDirect("secret", map[string]string{
		"sid":       "wdtapi3",
		"key":       "appkey",
		"salt":      "salt",
		"method":    "goods.Goods.queryWithSpec",
		"timestamp": "100",
		"v":         "1.0",
		"body":      `[{"goods_no":"G001"}]`,
	})
	if sign != "703b2c4b5b16565883e6b58d18e254a2" {
		t.Fatalf("unexpected direct sign: %s", sign)
	}
}

func TestSignQimenCustomMatchesReference(t *testing.T) {
	sign := SignQimenCustom("secret", "3ldsmu02o9.wdt.wms.stockspec.querychangehistory", map[string]any{
		"datetime":         "2020-09-17 19:22:26",
		"wdt_salt":         "salt",
		"wdt3_customer_id": "wdtapi3",
		"pager": map[string]any{
			"page_no":   1,
			"page_size": 10,
		},
		"wdt_appkey": "wdt_appkey",
		"params": map[string]any{
			"end_date":     "2020-09-20 00:00:00",
			"spec_no":      nil,
			"start_date":   "2020-08-23 00:00:00",
			"warehouse_no": "1001",
		},
	})
	if sign != "1f5e87bbbc8bc68f4f0af70f79620b7d" {
		t.Fatalf("unexpected qimen sign: %s", sign)
	}
}

func TestDirectCallBuildsSignedRequestAndDecodesData(t *testing.T) {
	var capturedQuery url.Values
	var capturedBody []map[string]string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.Query()
		if err := json.NewDecoder(r.Body).Decode(&capturedBody); err != nil {
			t.Fatalf("decode direct body: %v", err)
		}
		_, _ = w.Write([]byte(`{"status":0,"message":"ok","data":{"total_count":1,"goods_list":[{"goods_no":"G001","modified":"2026-05-06 12:00:00"}]}}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Direct: &DirectConfig{GatewayURL: server.URL, SID: "sid", AppKey: "key", AppSecret: "secret", Salt: "salt"},
		Now: func() time.Time {
			return time.Unix(directEpochOffset+100, 0)
		},
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	got, err := client.ListProduct(t.Context(), ListProductRequest{Pager: &Pager{PageNo: 1, PageSize: 20}, GoodsNo: "G001"}, WithRoute(RouteDirect))
	if err != nil {
		t.Fatalf("list product: %v", err)
	}
	if capturedQuery.Get("method") != "goods.Goods.queryWithSpec" || capturedQuery.Get("sign") == "" {
		t.Fatalf("unexpected direct query: %v", capturedQuery)
	}
	if len(capturedBody) != 1 || capturedBody[0]["goods_no"] != "G001" {
		t.Fatalf("unexpected direct body: %+v", capturedBody)
	}
	if got.TotalCount != 1 || got.Products[0].GoodsNo != "G001" || got.Products[0].Modified.ToDateTimeString() != "2026-05-06 12:00:00" {
		t.Fatalf("unexpected direct response: %+v", got)
	}
}

func TestQimenRouteCanBeChosenPerCall(t *testing.T) {
	var capturedQuery url.Values
	var capturedForm url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.Query()
		if err := r.ParseForm(); err != nil {
			t.Fatalf("parse form: %v", err)
		}
		capturedForm = r.PostForm
		_, _ = w.Write([]byte(`{"wdt_goods_apigoods_search_response":{"status":0,"message":"ok","data":{"total_count":1,"goods_list":[{"goods_no":"P001"}]}}}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		DefaultRoute: RouteDirect,
		Direct:       &DirectConfig{GatewayURL: "http://direct.invalid", SID: "sid", AppKey: "key", AppSecret: "secret", Salt: "salt"},
		Qimen: &QimenConfig{
			GatewayURL: server.URL, AppKey: "top-key", AppSecret: "top-secret", TargetAppKey: "target",
			CustomerID: "customer", WDTAppKey: "wdt-key", WDTAppSecret: "wdt-secret", WDTSalt: "salt",
		},
		Now: func() time.Time { return time.Date(2026, 5, 6, 12, 0, 0, 0, time.Local) },
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	got, err := client.ListPlatformProduct(t.Context(), ListPlatformProductRequest{Pager: &Pager{PageNo: 1, PageSize: 10}, TimeRange: TimeRange{StartTime: "2026-05-06 00:00:00"}}, WithRoute(RouteQimen))
	if err != nil {
		t.Fatalf("list platform product: %v", err)
	}
	if capturedQuery.Get("method") != "wdt.goods.apigoods.search" || capturedQuery.Get("target_app_key") != "target" || capturedQuery.Get("sign") == "" {
		t.Fatalf("unexpected qimen query: %v", capturedQuery)
	}
	if capturedForm.Get("wdt_appkey") != "wdt-key" || capturedForm.Get("wdt_sign") == "" {
		t.Fatalf("unexpected qimen form: %v", capturedForm)
	}
	if got.TotalCount != 1 || got.Products[0].GoodsNo != "P001" {
		t.Fatalf("unexpected qimen response: %+v", got)
	}
}

func TestOperationRegistryHasNoEmptyNames(t *testing.T) {
	seen := map[string]bool{}
	for _, op := range Operations {
		if op.Name == "" {
			t.Fatal("operation name is empty")
		}
		if seen[op.Name] {
			t.Fatalf("duplicate operation: %s", op.Name)
		}
		seen[op.Name] = true
		if op.DirectMethod == "" && op.QimenMethod == "" {
			t.Fatalf("operation %s has no route", op.Name)
		}
	}
}

func TestEveryDirectOperationDeserializesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"status":0,"message":"ok","data":{"marker":"direct"}}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Direct: &DirectConfig{GatewayURL: server.URL, SID: "sid", AppKey: "key", AppSecret: "secret", Salt: "salt"},
		Now:    func() time.Time { return time.Unix(directEpochOffset+100, 0) },
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	for _, op := range Operations {
		if op.DirectMethod == "" {
			continue
		}
		t.Run(op.Name, func(t *testing.T) {
			var got struct {
				Marker string `json:"marker"`
			}
			err := client.Call(t.Context(), Request{Operation: op, Pager: &Pager{PageNo: 1, PageSize: 10}, Params: map[string]string{"probe": op.Name}}, &got, WithRoute(RouteDirect))
			if err != nil {
				t.Fatalf("call direct: %v", err)
			}
			if got.Marker != "direct" {
				t.Fatalf("unexpected direct response: %+v", got)
			}
		})
	}
}

func TestEveryQimenOperationDeserializesResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"response":{"status":0,"message":"ok","data":{"marker":"qimen"}}}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{
		Qimen: &QimenConfig{
			GatewayURL: server.URL, AppKey: "top-key", AppSecret: "top-secret", TargetAppKey: "target",
			CustomerID: "customer", WDTAppKey: "wdt-key", WDTAppSecret: "wdt-secret", WDTSalt: "salt",
		},
		Now: func() time.Time { return time.Date(2026, 5, 6, 12, 0, 0, 0, time.Local) },
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	for _, op := range Operations {
		if op.QimenMethod == "" {
			continue
		}
		t.Run(op.Name, func(t *testing.T) {
			var got struct {
				Marker string `json:"marker"`
			}
			err := client.Call(t.Context(), Request{Operation: op, Pager: &Pager{PageNo: 1, PageSize: 10}, Params: map[string]string{"probe": op.Name}}, &got, WithRoute(RouteQimen))
			if err != nil {
				t.Fatalf("call qimen: %v", err)
			}
			if got.Marker != "qimen" {
				t.Fatalf("unexpected qimen response: %+v", got)
			}
		})
	}
}
