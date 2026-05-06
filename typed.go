package wdt

import "context"

type PageInfo struct {
	TotalCount int `json:"total_count,omitempty"`
	PageNo     int `json:"page_no,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

type TimeRange struct {
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

type ModifiedRange struct {
	StartModified string `json:"startModified,omitempty"`
	EndModified   string `json:"endModified,omitempty"`
}

type DateRange struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type MutationResponse struct {
	Code       string   `json:"code,omitempty"`
	Message    string   `json:"message,omitempty"`
	ID         string   `json:"id,omitempty"`
	IDs        []string `json:"ids,omitempty"`
	OrderNo    string   `json:"order_no,omitempty"`
	StockinNo  string   `json:"stockin_no,omitempty"`
	StockoutNo string   `json:"stockout_no,omitempty"`
}

func callTyped[T any](c *Client, ctx context.Context, op Operation, pager *Pager, params any, opts ...CallOption) (*T, error) {
	var out T
	if err := c.Call(ctx, Request{Operation: op, Pager: pager, Params: params}, &out, opts...); err != nil {
		return nil, err
	}
	return &out, nil
}
