# go-wdt

`go-wdt` is a Go SDK for WangDianTong APIs. It exposes one package and one client while supporting both direct WDT OpenAPI calls and Alibaba Qimen custom WDT calls.

The caller chooses the route globally or per request:

```go
client, err := wdt.NewClient(wdt.Config{
    DefaultRoute: wdt.RouteDirect,
    Direct: &wdt.DirectConfig{
        GatewayURL: "https://wdt.example.com/openapi",
        SID:        "sid",
        AppKey:     "app-key",
        AppSecret:  "app-secret",
        Salt:       "salt",
    },
    Qimen: &wdt.QimenConfig{
        GatewayURL:   "https://qimen.api.taobao.com/router/qm",
        AppKey:       "top-app-key",
        AppSecret:    "top-app-secret",
        TargetAppKey: "target-app-key",
        CustomerID:   "wdt-customer-id",
        WDTAppKey:    "wdt-app-key",
        WDTAppSecret: "wdt-secret",
        WDTSalt:      "wdt-salt",
    },
})
if err != nil {
    return err
}

resp, err := client.ListPlatformProduct(ctx, wdt.ListPlatformProductRequest{
    Pager: &wdt.Pager{PageNo: 1, PageSize: 50},
    TimeRange: wdt.TimeRange{
        StartTime: "2026-05-06 00:00:00",
        EndTime:   "2026-05-06 23:59:59",
    },
}, wdt.WithRoute(wdt.RouteQimen))
if err != nil {
    return err
}
for _, product := range resp.Products {
    fmt.Println(product.GoodsNo, product.SpecNo)
}
```

Each business API has a typed request and response, for example `ListProductRequest`, `ListProductResponse`, `CreatePurchaseOrderRequest`, and `ListStockResponse`. The low-level operation registry remains available as an escape hatch for custom fields or newly released WDT APIs:

```go
var raw struct {
    Total int `json:"total_count"`
}
err = client.Call(ctx, wdt.Request{
    Operation: wdt.OpListPurchaseOrder,
    Pager:     &wdt.Pager{PageNo: 1, PageSize: 20},
    Params:    map[string]string{"purchase_no": "PO001"},
}, &raw, wdt.WithRoute(wdt.RouteDirect))
```

Time fields can use `wdt.Time`, which unmarshals WDT string timestamps, RFC3339 timestamps, Unix seconds, Unix milliseconds, empty strings, and `null` values into `carbon.Carbon`.
