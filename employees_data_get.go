package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewEmployeeDataGetRequest() EmployeeDataGetRequest {
	r := EmployeeDataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeeDataGetRequest struct {
	client      *Client
	queryParams *EmployeeDataGetRequestQueryParams
	pathParams  *EmployeeDataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeeDataGetRequestBody
}

func (r EmployeeDataGetRequest) NewQueryParams() *EmployeeDataGetRequestQueryParams {
	return &EmployeeDataGetRequestQueryParams{}
}

type EmployeeDataGetRequestQueryParams struct {
}

func (p EmployeeDataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeeDataGetRequest) QueryParams() *EmployeeDataGetRequestQueryParams {
	return r.queryParams
}

func (r EmployeeDataGetRequest) NewPathParams() *EmployeeDataGetRequestPathParams {
	return &EmployeeDataGetRequestPathParams{}
}

type EmployeeDataGetRequestPathParams struct {
}

func (p *EmployeeDataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeeDataGetRequest) PathParams() *EmployeeDataGetRequestPathParams {
	return r.pathParams
}

func (r *EmployeeDataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeeDataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeeDataGetRequest) Method() string {
	return r.method
}

func (r EmployeeDataGetRequest) NewRequestBody() EmployeeDataGetRequestBody {
	return EmployeeDataGetRequestBody{}
}

type EmployeeDataGetRequestBody struct {
}

func (r *EmployeeDataGetRequest) RequestBody() *EmployeeDataGetRequestBody {
	return nil
}

func (r *EmployeeDataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeeDataGetRequest) SetRequestBody(body EmployeeDataGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeeDataGetRequest) NewResponseBody() *EmployeesResponseBody {
	return &EmployeesResponseBody{}
}

type EmployeesResponseBody struct {
}

func (r *EmployeeDataGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/employee-data", r.PathParams())
	return &u, err
}

func (r *EmployeeDataGetRequest) Do() (EmployeesResponseBody, error) {
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
