package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewOfficesGetRequest() OfficesGetRequest {
	r := OfficesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OfficesGetRequest struct {
	client      *Client
	queryParams *OfficesGetRequestQueryParams
	pathParams  *OfficesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody OfficesGetRequestBody
}

func (r OfficesGetRequest) NewQueryParams() *OfficesGetRequestQueryParams {
	return &OfficesGetRequestQueryParams{}
}

type OfficesGetRequestQueryParams struct {
}

func (p OfficesGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(Time{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *OfficesGetRequest) QueryParams() *OfficesGetRequestQueryParams {
	return r.queryParams
}

func (r OfficesGetRequest) NewPathParams() *OfficesGetRequestPathParams {
	return &OfficesGetRequestPathParams{}
}

type OfficesGetRequestPathParams struct {
}

func (p *OfficesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OfficesGetRequest) PathParams() *OfficesGetRequestPathParams {
	return r.pathParams
}

func (r *OfficesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OfficesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *OfficesGetRequest) Method() string {
	return r.method
}

func (r OfficesGetRequest) NewRequestBody() OfficesGetRequestBody {
	return OfficesGetRequestBody{}
}

type OfficesGetRequestBody struct {
}

func (r *OfficesGetRequest) RequestBody() *OfficesGetRequestBody {
	return nil
}

func (r *OfficesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *OfficesGetRequest) SetRequestBody(body OfficesGetRequestBody) {
	r.requestBody = body
}

func (r *OfficesGetRequest) NewResponseBody() *OfficesResponseBody {
	return &OfficesResponseBody{}
}

type OfficesResponseBody Offices

func (r *OfficesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/offices", r.PathParams())
	return &u, err
}

func (r *OfficesGetRequest) Do() (OfficesResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", r.client.API2Token()))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
