package wdt

import "context"

type StockinOrder struct {
	StockinID   string        `json:"stockin_id,omitempty"`
	StockinNo   string        `json:"stockin_no,omitempty"`
	WarehouseNo string        `json:"warehouse_no,omitempty"`
	Status      string        `json:"status,omitempty"`
	Reason      string        `json:"reason,omitempty"`
	Created     Time          `json:"created,omitempty"`
	Modified    Time          `json:"modified,omitempty"`
	Goods       []StockinItem `json:"goods_list,omitempty"`
}

type StockinItem struct {
	GoodsNo string  `json:"goods_no,omitempty"`
	SpecNo  string  `json:"spec_no,omitempty"`
	Num     float64 `json:"num,omitempty"`
	Price   float64 `json:"price,omitempty"`
	BatchNo string  `json:"batch_no,omitempty"`
}

type GetStockinDocRequest struct {
	StockinNo string `json:"stockin_no,omitempty"`
	StockinID string `json:"stockin_id,omitempty"`
}

type GetStockinDocResponse struct {
	Stockin  StockinOrder   `json:"stockin,omitempty"`
	Stockins []StockinOrder `json:"stockin_list,omitempty"`
}

type CreateStockinRecordRequest []struct {
	StockinNo   string        `json:"stockin_no,omitempty"`
	WarehouseNo string        `json:"warehouse_no,omitempty"`
	Reason      string        `json:"reason,omitempty"`
	Remark      string        `json:"remark,omitempty"`
	Goods       []StockinItem `json:"goods_list,omitempty"`
}

type ListPrestockinRequest struct {
	Pager *Pager `json:"-"`
	DateRange
	StockinNo   string `json:"stockin_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	Status      string `json:"status,omitempty"`
}

type ListPrestockinResponse struct {
	PageInfo
	Stockins []StockinOrder `json:"stockin_list,omitempty"`
}

type ListReturnStockinRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	StockinNo   string `json:"stockin_no,omitempty"`
	RefundNo    string `json:"refund_no,omitempty"`
	ShopNo      string `json:"shop_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
}

type ListReturnStockinResponse struct {
	PageInfo
	Stockins []StockinOrder `json:"stockin_list,omitempty"`
}

type StockoutOrder struct {
	StockoutID  string         `json:"stockout_id,omitempty"`
	StockoutNo  string         `json:"stockout_no,omitempty"`
	WarehouseNo string         `json:"warehouse_no,omitempty"`
	Status      string         `json:"status,omitempty"`
	Reason      string         `json:"reason,omitempty"`
	Created     Time           `json:"created,omitempty"`
	Modified    Time           `json:"modified,omitempty"`
	Goods       []StockoutItem `json:"goods_list,omitempty"`
}

type StockoutItem struct {
	GoodsNo string  `json:"goods_no,omitempty"`
	SpecNo  string  `json:"spec_no,omitempty"`
	Num     float64 `json:"num,omitempty"`
	Price   float64 `json:"price,omitempty"`
}

type ListOtherStockoutRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	StockoutNo  string `json:"stockout_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	Status      string `json:"status,omitempty"`
}

type ListOtherStockoutResponse struct {
	PageInfo
	Stockouts []StockoutOrder `json:"stockout_list,omitempty"`
}

type LogisticsTrace struct {
	LogisticsNo   string `json:"logistics_no,omitempty"`
	LogisticsCode string `json:"logistics_code,omitempty"`
	LogisticsName string `json:"logistics_name,omitempty"`
	Status        string `json:"status,omitempty"`
	Content       string `json:"content,omitempty"`
	Modified      Time   `json:"modified,omitempty"`
}

type ListLogisticsRequest struct {
	Pager       *Pager `json:"-"`
	LogisticsNo string `json:"logistics_no,omitempty"`
	StockoutNo  string `json:"stockout_no,omitempty"`
	TradeNo     string `json:"trade_no,omitempty"`
}

type ListLogisticsResponse struct {
	PageInfo
	Logistics []LogisticsTrace `json:"logistics_list,omitempty"`
}

func (c *Client) GetStockinDoc(ctx context.Context, req GetStockinDocRequest, opts ...CallOption) (*GetStockinDocResponse, error) {
	return callTyped[GetStockinDocResponse](c, ctx, OpGetStockinDoc, nil, req, opts...)
}

func (c *Client) CreateStockinRecord(ctx context.Context, req CreateStockinRecordRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateStockinRecord, nil, req, opts...)
}

func (c *Client) ListPrestockin(ctx context.Context, req ListPrestockinRequest, opts ...CallOption) (*ListPrestockinResponse, error) {
	return callTyped[ListPrestockinResponse](c, ctx, OpListPrestockin, req.Pager, req, opts...)
}

func (c *Client) ListReturnStockin(ctx context.Context, req ListReturnStockinRequest, opts ...CallOption) (*ListReturnStockinResponse, error) {
	return callTyped[ListReturnStockinResponse](c, ctx, OpListReturnStockin, req.Pager, req, opts...)
}

func (c *Client) ListReturnPrestockin(ctx context.Context, req ListPrestockinRequest, opts ...CallOption) (*ListPrestockinResponse, error) {
	return callTyped[ListPrestockinResponse](c, ctx, OpListReturnPrestockin, req.Pager, req, opts...)
}

func (c *Client) ListLogistics(ctx context.Context, req ListLogisticsRequest, opts ...CallOption) (*ListLogisticsResponse, error) {
	return callTyped[ListLogisticsResponse](c, ctx, OpListLogistics, req.Pager, req, opts...)
}

func (c *Client) ListOtherStockout(ctx context.Context, req ListOtherStockoutRequest, opts ...CallOption) (*ListOtherStockoutResponse, error) {
	return callTyped[ListOtherStockoutResponse](c, ctx, OpListOtherStockout, req.Pager, req, opts...)
}
