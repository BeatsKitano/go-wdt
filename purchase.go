package wdt

import "context"

type PurchaseOrder struct {
	PurchaseID   string              `json:"purchase_id,omitempty"`
	PurchaseNo   string              `json:"purchase_no,omitempty"`
	ProviderNo   string              `json:"provider_no,omitempty"`
	ProviderName string              `json:"provider_name,omitempty"`
	WarehouseNo  string              `json:"warehouse_no,omitempty"`
	Status       string              `json:"status,omitempty"`
	Remark       string              `json:"remark,omitempty"`
	Created      Time                `json:"created,omitempty"`
	Modified     Time                `json:"modified,omitempty"`
	Details      []PurchaseOrderItem `json:"details_list,omitempty"`
}

type PurchaseOrderItem struct {
	GoodsNo string  `json:"goods_no,omitempty"`
	SpecNo  string  `json:"spec_no,omitempty"`
	Num     float64 `json:"num,omitempty"`
	Price   float64 `json:"price,omitempty"`
	Remark  string  `json:"remark,omitempty"`
}

type CreatePurchaseOrderRequest struct {
	PurchaseNo  string              `json:"purchase_no,omitempty"`
	ProviderNo  string              `json:"provider_no,omitempty"`
	WarehouseNo string              `json:"warehouse_no,omitempty"`
	Remark      string              `json:"remark,omitempty"`
	Details     []PurchaseOrderItem `json:"details_list,omitempty"`
}

type ListPurchaseOrderRequest struct {
	Pager *Pager `json:"-"`
	DateRange
	PurchaseNo  string `json:"purchase_no,omitempty"`
	ProviderNo  string `json:"provider_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	Status      string `json:"status,omitempty"`
}

type GetPurchaseOrderRequest struct {
	PurchaseNo string `json:"purchase_no,omitempty"`
	PurchaseID string `json:"purchase_id,omitempty"`
}

type ListPurchaseOrderResponse struct {
	PageInfo
	PurchaseOrders []PurchaseOrder `json:"purchase_list,omitempty"`
}

type PurchaseOrderActionRequest struct {
	PurchaseNo string `json:"purchase_no,omitempty"`
	Reason     string `json:"reason,omitempty"`
	Action     string `json:"action,omitempty"`
}

type PurchaseStockinOrder struct {
	StockinID   string              `json:"stockin_id,omitempty"`
	StockinNo   string              `json:"stockin_no,omitempty"`
	PurchaseNo  string              `json:"purchase_no,omitempty"`
	WarehouseNo string              `json:"warehouse_no,omitempty"`
	Status      string              `json:"status,omitempty"`
	Created     Time                `json:"created,omitempty"`
	Details     []PurchaseOrderItem `json:"details_list,omitempty"`
}

type ListPurchaseStockinOrderRequest struct {
	Pager *Pager `json:"-"`
	DateRange
	StockinNo   string `json:"stockin_no,omitempty"`
	PurchaseNo  string `json:"purchase_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
}

type ListPurchaseStockinOrderResponse struct {
	PageInfo
	Stockins []PurchaseStockinOrder `json:"stockin_list,omitempty"`
}

type PurchaseReturn struct {
	ReturnID    string              `json:"return_id,omitempty"`
	ReturnNo    string              `json:"return_no,omitempty"`
	ProviderNo  string              `json:"provider_no,omitempty"`
	WarehouseNo string              `json:"warehouse_no,omitempty"`
	Status      string              `json:"status,omitempty"`
	Created     Time                `json:"created,omitempty"`
	Details     []PurchaseOrderItem `json:"details_list,omitempty"`
}

type CreatePurchaseReturnRequest []struct {
	ReturnNo    string              `json:"return_no,omitempty"`
	ProviderNo  string              `json:"provider_no,omitempty"`
	WarehouseNo string              `json:"warehouse_no,omitempty"`
	Remark      string              `json:"remark,omitempty"`
	Details     []PurchaseOrderItem `json:"details_list,omitempty"`
}

type ListPurchaseReturnRequest struct {
	Pager *Pager `json:"-"`
	DateRange
	ReturnNo    string `json:"return_no,omitempty"`
	ProviderNo  string `json:"provider_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	Status      string `json:"status,omitempty"`
}

type ListPurchaseReturnResponse struct {
	PageInfo
	Returns []PurchaseReturn `json:"return_list,omitempty"`
}

type PurchaseReturnActionRequest struct {
	ReturnNo string `json:"return_no,omitempty"`
	Reason   string `json:"reason,omitempty"`
}

func (c *Client) CreatePurchaseOrder(ctx context.Context, req CreatePurchaseOrderRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreatePurchaseOrder, nil, req, opts...)
}

func (c *Client) ListPurchaseOrder(ctx context.Context, req ListPurchaseOrderRequest, opts ...CallOption) (*ListPurchaseOrderResponse, error) {
	return callTyped[ListPurchaseOrderResponse](c, ctx, OpListPurchaseOrder, req.Pager, req, opts...)
}

func (c *Client) GetPurchaseOrder(ctx context.Context, req GetPurchaseOrderRequest, opts ...CallOption) (*ListPurchaseOrderResponse, error) {
	return callTyped[ListPurchaseOrderResponse](c, ctx, OpGetPurchaseOrder, nil, req, opts...)
}

func (c *Client) CancelOrStopPurchaseOrder(ctx context.Context, req PurchaseOrderActionRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCancelOrStopPurchaseOrder, nil, req, opts...)
}

func (c *Client) ListPurchaseStockinOrder(ctx context.Context, req ListPurchaseStockinOrderRequest, opts ...CallOption) (*ListPurchaseStockinOrderResponse, error) {
	return callTyped[ListPurchaseStockinOrderResponse](c, ctx, OpListPurchaseStockinOrder, req.Pager, req, opts...)
}

func (c *Client) CreatePurchaseReturn(ctx context.Context, req CreatePurchaseReturnRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreatePurchaseReturn, nil, req, opts...)
}

func (c *Client) ListPurchaseReturn(ctx context.Context, req ListPurchaseReturnRequest, opts ...CallOption) (*ListPurchaseReturnResponse, error) {
	return callTyped[ListPurchaseReturnResponse](c, ctx, OpListPurchaseReturn, req.Pager, req, opts...)
}

func (c *Client) CancelPurchaseReturn(ctx context.Context, req PurchaseReturnActionRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCancelPurchaseReturn, nil, req, opts...)
}

func (c *Client) StopPurchaseReturnStockout(ctx context.Context, req PurchaseReturnActionRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpStopPurchaseReturnStockout, nil, req, opts...)
}
