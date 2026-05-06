package wdt

import "context"

type Stock struct {
	StockID       string  `json:"stock_id,omitempty"`
	GoodsNo       string  `json:"goods_no,omitempty"`
	SpecNo        string  `json:"spec_no,omitempty"`
	GoodsName     string  `json:"goods_name,omitempty"`
	SpecName      string  `json:"spec_name,omitempty"`
	WarehouseNo   string  `json:"warehouse_no,omitempty"`
	WarehouseName string  `json:"warehouse_name,omitempty"`
	PositionNo    string  `json:"position_no,omitempty"`
	StockNum      float64 `json:"stock_num,omitempty"`
	AvailableNum  float64 `json:"available_num,omitempty"`
	LockNum       float64 `json:"lock_num,omitempty"`
	Modified      Time    `json:"modified,omitempty"`
}

type ListStockRequest struct {
	Pager       *Pager `json:"-"`
	GoodsNo     string `json:"goods_no,omitempty"`
	SpecNo      string `json:"spec_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	PositionNo  string `json:"position_no,omitempty"`
	TimeRange
}

type ListStockResponse struct {
	PageInfo
	Stocks []Stock `json:"stock_list,omitempty"`
}

type GetStockDetailRequest struct {
	GoodsNo     string `json:"goods_no,omitempty"`
	SpecNo      string `json:"spec_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
}

type GetStockDetailResponse struct {
	Stock  Stock   `json:"stock,omitempty"`
	Stocks []Stock `json:"stock_list,omitempty"`
}

type PositionStock struct {
	WarehouseNo string  `json:"warehouse_no,omitempty"`
	PositionNo  string  `json:"position_no,omitempty"`
	GoodsNo     string  `json:"goods_no,omitempty"`
	SpecNo      string  `json:"spec_no,omitempty"`
	StockNum    float64 `json:"stock_num,omitempty"`
}

type ListPositionStockRequest struct {
	Pager       *Pager `json:"-"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	PositionNo  string `json:"position_no,omitempty"`
	GoodsNo     string `json:"goods_no,omitempty"`
	SpecNo      string `json:"spec_no,omitempty"`
}

type ListPositionStockResponse struct {
	PageInfo
	Positions []PositionStock `json:"position_list,omitempty"`
}

type VirtualStock struct {
	WarehouseNo string  `json:"warehouse_no,omitempty"`
	GoodsNo     string  `json:"goods_no,omitempty"`
	SpecNo      string  `json:"spec_no,omitempty"`
	StockNum    float64 `json:"stock_num,omitempty"`
}

type ListVirtualStockRequest struct {
	Pager       *Pager `json:"-"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
	GoodsNo     string `json:"goods_no,omitempty"`
	SpecNo      string `json:"spec_no,omitempty"`
}

type ListVirtualStockResponse struct {
	PageInfo
	Stocks []VirtualStock `json:"stock_list,omitempty"`
}

type StockDefect struct {
	DefectID    string  `json:"defect_id,omitempty"`
	GoodsNo     string  `json:"goods_no,omitempty"`
	SpecNo      string  `json:"spec_no,omitempty"`
	WarehouseNo string  `json:"warehouse_no,omitempty"`
	ChangeNum   float64 `json:"change_num,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	Created     Time    `json:"created,omitempty"`
}

type ListStockDefectRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	GoodsNo     string `json:"goods_no,omitempty"`
	SpecNo      string `json:"spec_no,omitempty"`
	WarehouseNo string `json:"warehouse_no,omitempty"`
}

type ListStockDefectResponse struct {
	PageInfo
	Defects []StockDefect `json:"defect_list,omitempty"`
}

type StockTransferByPositionNoRequest []struct {
	FromWarehouseNo string  `json:"from_warehouse_no,omitempty"`
	FromPositionNo  string  `json:"from_position_no,omitempty"`
	ToWarehouseNo   string  `json:"to_warehouse_no,omitempty"`
	ToPositionNo    string  `json:"to_position_no,omitempty"`
	GoodsNo         string  `json:"goods_no,omitempty"`
	SpecNo          string  `json:"spec_no,omitempty"`
	Num             float64 `json:"num,omitempty"`
}

func (c *Client) CreateStockTransferByPositionNo(ctx context.Context, req StockTransferByPositionNoRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateStockTransferByPosition, nil, req, opts...)
}

func (c *Client) ListStockByPage(ctx context.Context, req ListStockRequest, opts ...CallOption) (*ListStockResponse, error) {
	return callTyped[ListStockResponse](c, ctx, OpListStockByPage, req.Pager, req, opts...)
}

func (c *Client) ListStock(ctx context.Context, req ListStockRequest, opts ...CallOption) (*ListStockResponse, error) {
	return callTyped[ListStockResponse](c, ctx, OpListStock, req.Pager, req, opts...)
}

func (c *Client) GetStockDetail(ctx context.Context, req GetStockDetailRequest, opts ...CallOption) (*GetStockDetailResponse, error) {
	return callTyped[GetStockDetailResponse](c, ctx, OpGetStockDetail, nil, req, opts...)
}

func (c *Client) ListDefaultPosition(ctx context.Context, req ListPositionStockRequest, opts ...CallOption) (*ListPositionStockResponse, error) {
	return callTyped[ListPositionStockResponse](c, ctx, OpListDefaultPosition, req.Pager, req, opts...)
}

func (c *Client) ListPositionStock(ctx context.Context, req ListPositionStockRequest, opts ...CallOption) (*ListPositionStockResponse, error) {
	return callTyped[ListPositionStockResponse](c, ctx, OpListPositionStock, req.Pager, req, opts...)
}

func (c *Client) ListVirtualStock(ctx context.Context, req ListVirtualStockRequest, opts ...CallOption) (*ListVirtualStockResponse, error) {
	return callTyped[ListVirtualStockResponse](c, ctx, OpListVirtualStock, req.Pager, req, opts...)
}

func (c *Client) ListVirtualWarehouse(ctx context.Context, req ListWarehouseRequest, opts ...CallOption) (*ListWarehouseResponse, error) {
	return callTyped[ListWarehouseResponse](c, ctx, OpListVirtualWarehouse, req.Pager, req, opts...)
}

func (c *Client) ListStockDefect(ctx context.Context, req ListStockDefectRequest, opts ...CallOption) (*ListStockDefectResponse, error) {
	return callTyped[ListStockDefectResponse](c, ctx, OpListStockDefect, req.Pager, req, opts...)
}

func (c *Client) ListStockSpec(ctx context.Context, req ListStockRequest, opts ...CallOption) (*ListStockResponse, error) {
	return callTyped[ListStockResponse](c, ctx, OpListStockSpec, req.Pager, req, opts...)
}
