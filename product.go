package wdt

import "context"

type Product struct {
	GoodsID     string  `json:"goods_id,omitempty"`
	GoodsNo     string  `json:"goods_no,omitempty"`
	GoodsName   string  `json:"goods_name,omitempty"`
	SpecID      string  `json:"spec_id,omitempty"`
	SpecNo      string  `json:"spec_no,omitempty"`
	SpecName    string  `json:"spec_name,omitempty"`
	Barcode     string  `json:"barcode,omitempty"`
	ClassID     string  `json:"class_id,omitempty"`
	ClassName   string  `json:"class_name,omitempty"`
	BrandNo     string  `json:"brand_no,omitempty"`
	BrandName   string  `json:"brand_name,omitempty"`
	UnitName    string  `json:"unit_name,omitempty"`
	RetailPrice float64 `json:"retail_price,omitempty"`
	Modified    Time    `json:"modified,omitempty"`
}

type ProductPayload struct {
	GoodsNo     string  `json:"goods_no,omitempty"`
	GoodsName   string  `json:"goods_name,omitempty"`
	SpecNo      string  `json:"spec_no,omitempty"`
	SpecName    string  `json:"spec_name,omitempty"`
	Barcode     string  `json:"barcode,omitempty"`
	ClassName   string  `json:"class_name,omitempty"`
	BrandName   string  `json:"brand_name,omitempty"`
	UnitName    string  `json:"unit_name,omitempty"`
	RetailPrice float64 `json:"retail_price,omitempty"`
}

type ListProductRequest struct {
	Pager   *Pager `json:"-"`
	GoodsNo string `json:"goods_no,omitempty"`
	SpecNo  string `json:"spec_no,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	ClassID string `json:"class_id,omitempty"`
	BrandNo string `json:"brand_no,omitempty"`
	TimeRange
}

type ListProductResponse struct {
	PageInfo
	Products []Product `json:"goods_list,omitempty"`
}

type GetProductRequest struct {
	GoodsNo string `json:"goods_no,omitempty"`
	SpecNo  string `json:"spec_no,omitempty"`
	Barcode string `json:"barcode,omitempty"`
}

type GetProductResponse struct {
	Product  Product   `json:"goods,omitempty"`
	Products []Product `json:"goods_list,omitempty"`
}

type CreateProductRequest struct {
	Products []ProductPayload `json:"goods_list,omitempty"`
}

type ListCategoryRequest struct {
	Pager     *Pager `json:"-"`
	ClassID   string `json:"class_id,omitempty"`
	ClassName string `json:"class_name,omitempty"`
}

type Category struct {
	ClassID   string `json:"class_id,omitempty"`
	ClassNo   string `json:"class_no,omitempty"`
	ClassName string `json:"class_name,omitempty"`
	ParentID  string `json:"parent_id,omitempty"`
}

type ListCategoryResponse struct {
	PageInfo
	Categories []Category `json:"class_list,omitempty"`
}

type ListPlatformProductRequest struct {
	Pager *Pager `json:"-"`
	TimeRange
	ShopNo  string `json:"shop_no,omitempty"`
	GoodsNo string `json:"goods_no,omitempty"`
	SpecNo  string `json:"spec_no,omitempty"`
}

type ListPlatformProductResponse struct {
	PageInfo
	Products []Product `json:"goods_list,omitempty"`
}

type SKU struct {
	SuiteNo   string `json:"suite_no,omitempty"`
	SuiteName string `json:"suite_name,omitempty"`
	GoodsNo   string `json:"goods_no,omitempty"`
	SpecNo    string `json:"spec_no,omitempty"`
	SpecName  string `json:"spec_name,omitempty"`
}

type ListSkuRequest struct {
	Pager   *Pager `json:"-"`
	SuiteNo string `json:"suite_no,omitempty"`
	GoodsNo string `json:"goods_no,omitempty"`
	SpecNo  string `json:"spec_no,omitempty"`
	TimeRange
}

type ListSkuResponse struct {
	PageInfo
	Skus []SKU `json:"suite_list,omitempty"`
}

func (c *Client) ListProduct(ctx context.Context, req ListProductRequest, opts ...CallOption) (*ListProductResponse, error) {
	return callTyped[ListProductResponse](c, ctx, OpListProduct, req.Pager, req, opts...)
}

func (c *Client) GetProduct(ctx context.Context, req GetProductRequest, opts ...CallOption) (*GetProductResponse, error) {
	return callTyped[GetProductResponse](c, ctx, OpGetProduct, nil, req, opts...)
}

func (c *Client) CreateProduct(ctx context.Context, req CreateProductRequest, opts ...CallOption) (*MutationResponse, error) {
	return callTyped[MutationResponse](c, ctx, OpCreateProduct, nil, req, opts...)
}

func (c *Client) ListCategory(ctx context.Context, req ListCategoryRequest, opts ...CallOption) (*ListCategoryResponse, error) {
	return callTyped[ListCategoryResponse](c, ctx, OpListCategory, req.Pager, req, opts...)
}

func (c *Client) ListPlatformProduct(ctx context.Context, req ListPlatformProductRequest, opts ...CallOption) (*ListPlatformProductResponse, error) {
	return callTyped[ListPlatformProductResponse](c, ctx, OpListPlatformProduct, req.Pager, req, opts...)
}

func (c *Client) ListSku(ctx context.Context, req ListSkuRequest, opts ...CallOption) (*ListSkuResponse, error) {
	return callTyped[ListSkuResponse](c, ctx, OpListSku, req.Pager, req, opts...)
}
