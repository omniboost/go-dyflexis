package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewEmployeeMappingsGetRequest() EmployeeMappingsGetRequest {
	r := EmployeeMappingsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeeMappingsGetRequest struct {
	client      *Client
	queryParams *EmployeeMappingsGetRequestQueryParams
	pathParams  *EmployeeMappingsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeeMappingsGetRequestBody
}

func (r EmployeeMappingsGetRequest) NewQueryParams() *EmployeeMappingsGetRequestQueryParams {
	return &EmployeeMappingsGetRequestQueryParams{}
}

type EmployeeMappingsGetRequestQueryParams struct {
}

func (p EmployeeMappingsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeeMappingsGetRequest) QueryParams() *EmployeeMappingsGetRequestQueryParams {
	return r.queryParams
}

func (r EmployeeMappingsGetRequest) NewPathParams() *EmployeeMappingsGetRequestPathParams {
	return &EmployeeMappingsGetRequestPathParams{}
}

type EmployeeMappingsGetRequestPathParams struct {
}

func (p *EmployeeMappingsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeMappingsGetRequest) PathParams() *EmployeeMappingsGetRequestPathParams {
	return r.pathParams
}

func (r *EmployeeMappingsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeeMappingsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeMappingsGetRequest) Method() string {
	return r.method
}

func (r EmployeeMappingsGetRequest) NewRequestBody() EmployeeMappingsGetRequestBody {
	return EmployeeMappingsGetRequestBody{}
}

type EmployeeMappingsGetRequestBody struct {
}

func (r *EmployeeMappingsGetRequest) RequestBody() *EmployeeMappingsGetRequestBody {
	return nil
}

func (r *EmployeeMappingsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeeMappingsGetRequest) SetRequestBody(body EmployeeMappingsGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeeMappingsGetRequest) NewResponseBody() *EmployeeMappingsResponseBody {
	return &EmployeeMappingsResponseBody{}
}

type EmployeeMappingsResponseBody struct {
}

func (r *EmployeeMappingsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/employee-mappings", r.PathParams())
	return &u, err
}

func (r *EmployeeMappingsGetRequest) Do() (EmployeeMappingsResponseBody, error) {
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
