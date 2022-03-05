// Package productPricing provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package productPricing

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/mikeshimura/mwssp/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetCompetitivePricing request
	GetCompetitivePricing(ctx context.Context, params *GetCompetitivePricingParams) (*http.Response, error)

	// GetItemOffers request
	GetItemOffers(ctx context.Context, asin string, params *GetItemOffersParams) (*http.Response, error)

	// GetListingOffers request
	GetListingOffers(ctx context.Context, sellerSKU string, params *GetListingOffersParams) (*http.Response, error)

	// GetPricing request
	GetPricing(ctx context.Context, params *GetPricingParams) (*http.Response, error)
}

func (c *Client) GetCompetitivePricing(ctx context.Context, params *GetCompetitivePricingParams) (*http.Response, error) {
	req, err := NewGetCompetitivePricingRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetItemOffers(ctx context.Context, asin string, params *GetItemOffersParams) (*http.Response, error) {
	req, err := NewGetItemOffersRequest(c.Endpoint, asin, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetListingOffers(ctx context.Context, sellerSKU string, params *GetListingOffersParams) (*http.Response, error) {
	req, err := NewGetListingOffersRequest(c.Endpoint, sellerSKU, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) GetPricing(ctx context.Context, params *GetPricingParams) (*http.Response, error) {
	req, err := NewGetPricingRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewGetCompetitivePricingRequest generates requests for GetCompetitivePricing
func NewGetCompetitivePricingRequest(endpoint string, params *GetCompetitivePricingParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/products/pricing/v0/competitivePrice")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Asins != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "Asins", *params.Asins); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				vv:=""
				for _, v2 := range v {
					//queryValues.Add(k, v2)
					if vv==""{
						vv=v2
					} else {
						vv+=","+v2
					}
				}
				queryValues.Add(k, vv)
			}
		}

	}

	if params.Skus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "Skus", *params.Skus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParam("form", true, "ItemType", params.ItemType); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetItemOffersRequest generates requests for GetItemOffers
func NewGetItemOffersRequest(endpoint string, asin string, params *GetItemOffersParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "Asin", asin)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/products/pricing/v0/items/%s/offers", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParam("form", true, "ItemCondition", params.ItemCondition); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetListingOffersRequest generates requests for GetListingOffers
func NewGetListingOffersRequest(endpoint string, sellerSKU string, params *GetListingOffersParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "SellerSKU", sellerSKU)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/products/pricing/v0/listings/%s/offers", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if queryFrag, err := runtime.StyleParam("form", true, "ItemCondition", params.ItemCondition); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPricingRequest generates requests for GetPricing
func NewGetPricingRequest(endpoint string, params *GetPricingParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/products/pricing/v0/price")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if queryFrag, err := runtime.StyleParam("form", true, "MarketplaceId", params.MarketplaceId); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.Asins != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "Asins", *params.Asins); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if params.Skus != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "Skus", *params.Skus); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	if queryFrag, err := runtime.StyleParam("form", true, "ItemType", params.ItemType); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	if params.ItemCondition != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "ItemCondition", *params.ItemCondition); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetCompetitivePricing request
	GetCompetitivePricingWithResponse(ctx context.Context, params *GetCompetitivePricingParams) (*GetCompetitivePricingResp, error)

	// GetItemOffers request
	GetItemOffersWithResponse(ctx context.Context, asin string, params *GetItemOffersParams) (*GetItemOffersResp, error)

	// GetListingOffers request
	GetListingOffersWithResponse(ctx context.Context, sellerSKU string, params *GetListingOffersParams) (*GetListingOffersResp, error)

	// GetPricing request
	GetPricingWithResponse(ctx context.Context, params *GetPricingParams) (*GetPricingResp, error)
}

type GetCompetitivePricingResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetPricingResponse
}

// Status returns HTTPResponse.Status
func (r GetCompetitivePricingResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetCompetitivePricingResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetItemOffersResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetOffersResponse
}

// Status returns HTTPResponse.Status
func (r GetItemOffersResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetItemOffersResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetListingOffersResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetOffersResponse
}

// Status returns HTTPResponse.Status
func (r GetListingOffersResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetListingOffersResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPricingResp struct {
	Body         []byte
	HTTPResponse *http.Response
	Model        *GetPricingResponse
}

// Status returns HTTPResponse.Status
func (r GetPricingResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPricingResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetCompetitivePricingWithResponse request returning *GetCompetitivePricingResponse
func (c *ClientWithResponses) GetCompetitivePricingWithResponse(ctx context.Context, params *GetCompetitivePricingParams) (*GetCompetitivePricingResp, error) {
	rsp, err := c.GetCompetitivePricing(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetCompetitivePricingResp(rsp)
}

// GetItemOffersWithResponse request returning *GetItemOffersResponse
func (c *ClientWithResponses) GetItemOffersWithResponse(ctx context.Context, asin string, params *GetItemOffersParams) (*GetItemOffersResp, error) {
	rsp, err := c.GetItemOffers(ctx, asin, params)
	if err != nil {
		return nil, err
	}
	return ParseGetItemOffersResp(rsp)
}

// GetListingOffersWithResponse request returning *GetListingOffersResponse
func (c *ClientWithResponses) GetListingOffersWithResponse(ctx context.Context, sellerSKU string, params *GetListingOffersParams) (*GetListingOffersResp, error) {
	rsp, err := c.GetListingOffers(ctx, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	return ParseGetListingOffersResp(rsp)
}

// GetPricingWithResponse request returning *GetPricingResponse
func (c *ClientWithResponses) GetPricingWithResponse(ctx context.Context, params *GetPricingParams) (*GetPricingResp, error) {
	rsp, err := c.GetPricing(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseGetPricingResp(rsp)
}

// ParseGetCompetitivePricingResp parses an HTTP response from a GetCompetitivePricingWithResponse call
func ParseGetCompetitivePricingResp(rsp *http.Response) (*GetCompetitivePricingResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetCompetitivePricingResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetPricingResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetItemOffersResp parses an HTTP response from a GetItemOffersWithResponse call
func ParseGetItemOffersResp(rsp *http.Response) (*GetItemOffersResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetItemOffersResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetOffersResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetListingOffersResp parses an HTTP response from a GetListingOffersWithResponse call
func ParseGetListingOffersResp(rsp *http.Response) (*GetListingOffersResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetListingOffersResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetOffersResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}

// ParseGetPricingResp parses an HTTP response from a GetPricingWithResponse call
func ParseGetPricingResp(rsp *http.Response) (*GetPricingResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetPricingResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	var dest GetPricingResponse
	if err := json.Unmarshal(bodyBytes, &dest); err != nil {
		return nil, err
	}

	response.Model = &dest

	if rsp.StatusCode >= 300 {
		err = fmt.Errorf(rsp.Status)
	}

	return response, err
}
