package wdt

import "context"

type RawOrder struct {
	Tid             string         `json:"tid,omitempty"`
	ShopNo          string         `json:"shop_no,omitempty"`
	PlatformID      string         `json:"platform_id,omitempty"`
	BuyerNick       string         `json:"buyer_nick,omitempty"`
	ReceiverName    string         `json:"receiver_name,omitempty"`
	ReceiverMobile  string         `json:"receiver_mobile,omitempty"`
	ReceiverAddress string         `json:"receiver_address,omitempty"`
	PayTime         string         `json:"pay_time,omitempty"`
	Remark          string         `json:"remark,omitempty"`
	Goods           []RawOrderItem `json:"goods_list,omitempty"`
}

type RawOrderItem struct {
	Oid      string  `json:"oid,omitempty"`
	GoodsNo  string  `json:"goods_no,omitempty"`
	SpecNo   string  `json:"spec_no,omitempty"`
	Num      float64 `json:"num,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Discount float64 `json:"discount,omitempty"`
}

type CreateRawOrderRequest []RawOrder

type ListRawOrderRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	ShopNo string `json:"shop_no,omitempty"`
	Tid    string `json:"tid,omitempty"`
	Status string `json:"status,omitempty"`
}

type ListRawOrderResponse struct {
	PageInfo
	Orders []RawOrder `json:"trade_list,omitempty"`
}

type SaleOrder struct {
	TradeID         string          `json:"trade_id,omitempty"`
	TradeNo         string          `json:"trade_no,omitempty"`
	SrcTid          string          `json:"src_tid,omitempty"`
	ShopNo          string          `json:"shop_no,omitempty"`
	WarehouseNo     string          `json:"warehouse_no,omitempty"`
	TradeStatus     string          `json:"trade_status,omitempty"`
	BuyerNick       string          `json:"buyer_nick,omitempty"`
	ReceiverName    string          `json:"receiver_name,omitempty"`
	ReceiverMobile  string          `json:"receiver_mobile,omitempty"`
	ReceiverAddress string          `json:"receiver_address,omitempty"`
	Paid            float64         `json:"paid,omitempty"`
	PostAmount      float64         `json:"post_amount,omitempty"`
	Created         Time            `json:"created,omitempty"`
	Modified        Time            `json:"modified,omitempty"`
	Goods           []SaleOrderItem `json:"goods_list,omitempty"`
}

type SaleOrderItem struct {
	RecID     string  `json:"rec_id,omitempty"`
	GoodsNo   string  `json:"goods_no,omitempty"`
	SpecNo    string  `json:"spec_no,omitempty"`
	GoodsName string  `json:"goods_name,omitempty"`
	SpecName  string  `json:"spec_name,omitempty"`
	Num       float64 `json:"num,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

type ListSaleOrderRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	TradeNo     string `json:"trade_no,omitempty"`
	SrcTid      string `json:"src_tid,omitempty"`
	ShopNo      string `json:"shop_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	Status      string `json:"status,omitempty"`
}

type ListSaleOrderResponse struct {
	PageInfo
	Orders []SaleOrder `json:"trade_list,omitempty"`
}

type ReturnOrder struct {
	RefundID string          `json:"refund_id,omitempty"`
	RefundNo string          `json:"refund_no,omitempty"`
	TradeNo  string          `json:"trade_no,omitempty"`
	ShopNo   string          `json:"shop_no,omitempty"`
	Status   string          `json:"status,omitempty"`
	Reason   string          `json:"reason,omitempty"`
	Created  Time            `json:"created,omitempty"`
	Modified Time            `json:"modified,omitempty"`
	Goods    []SaleOrderItem `json:"goods_list,omitempty"`
}

type ListReturnOrderRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	RefundNo string `json:"refund_no,omitempty"`
	TradeNo  string `json:"trade_no,omitempty"`
	ShopNo   string `json:"shop_no,omitempty"`
	Status   string `json:"status,omitempty"`
}

type ListReturnOrderResponse struct {
	PageInfo
	Returns []ReturnOrder `json:"refund_list,omitempty"`
}

type SaleStockout struct {
	StockoutID  string          `json:"stockout_id,omitempty"`
	StockoutNo  string          `json:"stockout_no,omitempty"`
	TradeNo     string          `json:"trade_no,omitempty"`
	WarehouseNo string          `json:"warehouse_no,omitempty"`
	LogisticsNo string          `json:"logistics_no,omitempty"`
	Status      string          `json:"status,omitempty"`
	ConsignTime Time            `json:"consign_time,omitempty"`
	Goods       []SaleOrderItem `json:"goods_list,omitempty"`
}

type ListSaleStockoutRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	StockoutNo  string `json:"stockout_no,omitempty"`
	TradeNo     string `json:"trade_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
}

type ListSaleStockoutResponse struct {
	PageInfo
	Stockouts []SaleStockout `json:"stockout_list,omitempty"`
}

type SalesPayment struct {
	PayID      string  `json:"pay_id,omitempty"`
	TradeNo    string  `json:"trade_no,omitempty"`
	PayAccount string  `json:"pay_account,omitempty"`
	PayAmount  float64 `json:"pay_amount,omitempty"`
	PayTime    Time    `json:"pay_time,omitempty"`
}

type ListSalesPaymentRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	TradeNo string `json:"trade_no,omitempty"`
	ShopNo  string `json:"shop_no,omitempty"`
}

type ListSalesPaymentResponse struct {
	PageInfo
	Payments []SalesPayment `json:"payment_list,omitempty"`
}

type ReissueOrderRequest []struct {
	TradeNo string `json:"trade_no,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

func (c *Client) CreateRawOrder(ctx context.Context, req CreateRawOrderRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateRawOrder, nil, req, opts...)
}

func (c *Client) ListRawOrder(ctx context.Context, req ListRawOrderRequest, opts ...CallOption) (*ListRawOrderResponse, error) {
	return callTyped[ListRawOrderResponse](c, ctx, OpListRawOrder, req.Pager, req, opts...)
}

func (c *Client) ListHistoryRawOrder(ctx context.Context, req ListRawOrderRequest, opts ...CallOption) (*ListRawOrderResponse, error) {
	return callTyped[ListRawOrderResponse](c, ctx, OpListHistoryRawOrder, req.Pager, req, opts...)
}

func (c *Client) ListSaleOrder(ctx context.Context, req ListSaleOrderRequest, opts ...CallOption) (*ListSaleOrderResponse, error) {
	return callTyped[ListSaleOrderResponse](c, ctx, OpListSaleOrder, req.Pager, req, opts...)
}

func (c *Client) ListArchivedSaleOrder(ctx context.Context, req ListSaleOrderRequest, opts ...CallOption) (*ListSaleOrderResponse, error) {
	return callTyped[ListSaleOrderResponse](c, ctx, OpListArchivedSaleOrder, req.Pager, req, opts...)
}

func (c *Client) ListReturnOrder(ctx context.Context, req ListReturnOrderRequest, opts ...CallOption) (*ListReturnOrderResponse, error) {
	return callTyped[ListReturnOrderResponse](c, ctx, OpListReturnOrder, req.Pager, req, opts...)
}

func (c *Client) ListArchivedReturnOrder(ctx context.Context, req ListReturnOrderRequest, opts ...CallOption) (*ListReturnOrderResponse, error) {
	return callTyped[ListReturnOrderResponse](c, ctx, OpListArchivedReturn, req.Pager, req, opts...)
}

func (c *Client) ListSaleStockout(ctx context.Context, req ListSaleStockoutRequest, opts ...CallOption) (*ListSaleStockoutResponse, error) {
	return callTyped[ListSaleStockoutResponse](c, ctx, OpListSaleStockout, req.Pager, req, opts...)
}

func (c *Client) ListSalesPayment(ctx context.Context, req ListSalesPaymentRequest, opts ...CallOption) (*ListSalesPaymentResponse, error) {
	return callTyped[ListSalesPaymentResponse](c, ctx, OpListSalesPayment, req.Pager, req, opts...)
}

func (c *Client) ReissueOrder(ctx context.Context, req ReissueOrderRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpReissueOrder, nil, req, opts...)
}
