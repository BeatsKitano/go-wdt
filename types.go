package wdt

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Route string

const (
	RouteAuto   Route = "auto"
	RouteDirect Route = "direct"
	RouteQimen  Route = "qimen"
)

type DirectConfig struct {
	GatewayURL string
	SID        string
	AppKey     string
	AppSecret  string
	Salt       string
	Version    string
}

type QimenConfig struct {
	GatewayURL   string
	AppKey       string
	AppSecret    string
	TargetAppKey string
	CustomerID   string
	WDTAppKey    string
	WDTAppSecret string
	WDTSalt      string

	Format     string
	SignMethod string
	Version    string
	PartnerID  string
}

type Config struct {
	DefaultRoute Route
	Direct       *DirectConfig
	Qimen        *QimenConfig
	HTTPClient   *http.Client
	Now          func() time.Time
}

type Client struct {
	defaultRoute Route
	direct       *directTransport
	qimen        *qimenTransport
}

func NewClient(cfg Config) (*Client, error) {
	now := cfg.Now
	if now == nil {
		now = time.Now
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 30 * time.Second}
	}

	client := &Client{defaultRoute: cfg.DefaultRoute}
	if cfg.Direct != nil {
		direct, err := newDirectTransport(*cfg.Direct, httpClient, now)
		if err != nil {
			return nil, err
		}
		client.direct = direct
	}
	if cfg.Qimen != nil {
		qimen, err := newQimenTransport(*cfg.Qimen, httpClient, now)
		if err != nil {
			return nil, err
		}
		client.qimen = qimen
	}
	if client.direct == nil && client.qimen == nil {
		return nil, fmt.Errorf("wdt: at least one transport config is required")
	}
	if client.defaultRoute == "" || client.defaultRoute == RouteAuto {
		if client.direct != nil {
			client.defaultRoute = RouteDirect
		} else {
			client.defaultRoute = RouteQimen
		}
	}
	return client, nil
}

type Pager struct {
	PageNo    int  `json:"page_no,omitempty"`
	PageSize  int  `json:"page_size,omitempty"`
	CalcTotal bool `json:"calc_total,omitempty"`
}

type DirectBodyStyle int

const (
	DirectBodyObject DirectBodyStyle = iota
	DirectBodyArray
)

type Operation struct {
	Name         string
	DirectMethod string
	QimenMethod  string
	DirectBody   DirectBodyStyle
}

func (op Operation) method(route Route) (string, error) {
	switch route {
	case RouteDirect:
		if op.DirectMethod == "" {
			return "", fmt.Errorf("wdt: operation %s does not support direct route", op.Name)
		}
		return op.DirectMethod, nil
	case RouteQimen:
		if op.QimenMethod == "" {
			return "", fmt.Errorf("wdt: operation %s does not support qimen route", op.Name)
		}
		return op.QimenMethod, nil
	default:
		return "", fmt.Errorf("wdt: unsupported route %q", route)
	}
}

type Request struct {
	Operation Operation
	Pager     *Pager
	Params    any
	Route     Route
}

type CallOption func(*Request)

func WithRoute(route Route) CallOption {
	return func(req *Request) { req.Route = route }
}

func WithPager(pager *Pager) CallOption {
	return func(req *Request) { req.Pager = pager }
}

func (c *Client) Call(ctx context.Context, req Request, out any, opts ...CallOption) error {
	for _, opt := range opts {
		opt(&req)
	}
	route := req.Route
	if route == "" || route == RouteAuto {
		route = c.defaultRoute
	}
	method, err := req.Operation.method(route)
	if err != nil {
		return err
	}

	switch route {
	case RouteDirect:
		if c.direct == nil {
			return fmt.Errorf("wdt: direct transport is not configured")
		}
		return c.direct.call(ctx, method, req.Operation.DirectBody, req.Pager, req.Params, out)
	case RouteQimen:
		if c.qimen == nil {
			return fmt.Errorf("wdt: qimen transport is not configured")
		}
		return c.qimen.call(ctx, method, req.Pager, req.Params, out)
	default:
		return fmt.Errorf("wdt: unsupported route %q", route)
	}
}
