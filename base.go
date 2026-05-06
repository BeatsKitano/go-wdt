package wdt

import "context"

type Supplier struct {
	ProviderID   string `json:"provider_id,omitempty"`
	ProviderNo   string `json:"provider_no,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	Contact      string `json:"contact,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Telno        string `json:"telno,omitempty"`
	Address      string `json:"address,omitempty"`
	Modified     Time   `json:"modified,omitempty"`
}

type ListSupplierRequest struct {
	Pager        *Pager `json:"-"`
	ProviderNo   string `json:"provider_no,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	TimeRange
}

type ListSupplierResponse struct {
	PageInfo
	Suppliers []Supplier `json:"provider_list,omitempty"`
}

type CreateSupplierRequest struct {
	ProviderNo   string `json:"provider_no,omitempty"`
	ProviderName string `json:"provider_name,omitempty"`
	Contact      string `json:"contact,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	Telno        string `json:"telno,omitempty"`
	Address      string `json:"address,omitempty"`
}

type ProductSupplier struct {
	ProviderNo string  `json:"provider_no,omitempty"`
	GoodsNo    string  `json:"goods_no,omitempty"`
	SpecNo     string  `json:"spec_no,omitempty"`
	Price      float64 `json:"price,omitempty"`
}

type CreateProductSupplierRequest []ProductSupplier

type Warehouse struct {
	WarehouseID   string `json:"warehouse_id,omitempty"`
	WarehouseNo   string `json:"warehouse_no,omitempty"`
	WarehouseName string `json:"warehouse_name,omitempty"`
	Type          string `json:"type,omitempty"`
	Province      string `json:"province,omitempty"`
	City          string `json:"city,omitempty"`
	District      string `json:"district,omitempty"`
	Address       string `json:"address,omitempty"`
}

type ListWarehouseRequest struct {
	Pager         *Pager `json:"-"`
	WarehouseNo   string `json:"warehouse_no,omitempty"`
	WarehouseName string `json:"warehouse_name,omitempty"`
}

type ListWarehouseResponse struct {
	PageInfo
	Warehouses []Warehouse `json:"warehouse_list,omitempty"`
}

type Shop struct {
	ShopID   string `json:"shop_id,omitempty"`
	ShopNo   string `json:"shop_no,omitempty"`
	ShopName string `json:"shop_name,omitempty"`
	Platform string `json:"platform,omitempty"`
}

type ListShopRequest struct {
	Pager    *Pager `json:"-"`
	ShopNo   string `json:"shop_no,omitempty"`
	ShopName string `json:"shop_name,omitempty"`
	Platform string `json:"platform,omitempty"`
}

type ListShopResponse struct {
	PageInfo
	Shops []Shop `json:"shop_list,omitempty"`
}

type LogisticsCompany struct {
	LogisticsID   string `json:"logistics_id,omitempty"`
	LogisticsNo   string `json:"logistics_no,omitempty"`
	LogisticsCode string `json:"logistics_code,omitempty"`
	LogisticsName string `json:"logistics_name,omitempty"`
}

type ListLogisticsCompanyRequest struct {
	Pager         *Pager `json:"-"`
	LogisticsNo   string `json:"logistics_no,omitempty"`
	LogisticsCode string `json:"logistics_code,omitempty"`
	LogisticsName string `json:"logistics_name,omitempty"`
}

type ListLogisticsCompanyResponse struct {
	PageInfo
	LogisticsCompanies []LogisticsCompany `json:"logistics_list,omitempty"`
}

func (c *Client) ListSupplier(ctx context.Context, req ListSupplierRequest, opts ...CallOption) (*ListSupplierResponse, error) {
	return callTyped[ListSupplierResponse](c, ctx, OpListSupplier, req.Pager, req, opts...)
}

func (c *Client) CreateSupplier(ctx context.Context, req CreateSupplierRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateSupplier, nil, req, opts...)
}

func (c *Client) CreateProductSupplier(ctx context.Context, req CreateProductSupplierRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateProductSupplier, nil, req, opts...)
}

func (c *Client) ListWarehouse(ctx context.Context, req ListWarehouseRequest, opts ...CallOption) (*ListWarehouseResponse, error) {
	return callTyped[ListWarehouseResponse](c, ctx, OpListWarehouse, req.Pager, req, opts...)
}

func (c *Client) ListShop(ctx context.Context, req ListShopRequest, opts ...CallOption) (*ListShopResponse, error) {
	return callTyped[ListShopResponse](c, ctx, OpListShop, req.Pager, req, opts...)
}

func (c *Client) ListLogisticsCompany(ctx context.Context, req ListLogisticsCompanyRequest, opts ...CallOption) (*ListLogisticsCompanyResponse, error) {
	return callTyped[ListLogisticsCompanyResponse](c, ctx, OpListLogisticsCompany, req.Pager, req, opts...)
}
