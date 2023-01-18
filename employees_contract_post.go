package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewEmployeesContractPostRequest() EmployeesContractPostRequest {
	r := EmployeesContractPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesContractPostRequest struct {
	client      *Client
	queryParams *EmployeesContractPostRequestQueryParams
	pathParams  *EmployeesContractPostRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeesContractPostRequestBody
}

func (r EmployeesContractPostRequest) NewQueryParams() *EmployeesContractPostRequestQueryParams {
	return &EmployeesContractPostRequestQueryParams{}
}

type EmployeesContractPostRequestQueryParams struct {
}

func (p EmployeesContractPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesContractPostRequest) QueryParams() *EmployeesContractPostRequestQueryParams {
	return r.queryParams
}

func (r EmployeesContractPostRequest) NewPathParams() *EmployeesContractPostRequestPathParams {
	return &EmployeesContractPostRequestPathParams{}
}

type EmployeesContractPostRequestPathParams struct {
	EmployeeReference string `schema:"employeeReference"`
}

func (p *EmployeesContractPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"employeeReference": p.EmployeeReference,
	}
}

func (r *EmployeesContractPostRequest) PathParams() *EmployeesContractPostRequestPathParams {
	return r.pathParams
}

func (r *EmployeesContractPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesContractPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesContractPostRequest) Method() string {
	return r.method
}

func (r EmployeesContractPostRequest) NewRequestBody() EmployeesContractPostRequestBody {
	return EmployeesContractPostRequestBody{}
}

type EmployeesContractPostRequestBody Contract

func (r *EmployeesContractPostRequest) RequestBody() *EmployeesContractPostRequestBody {
	return &r.requestBody
}

func (r *EmployeesContractPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *EmployeesContractPostRequest) SetRequestBody(body EmployeesContractPostRequestBody) {
	r.requestBody = body
}

func (r *EmployeesContractPostRequest) NewResponseBody() *EmployeesContractPostResponseBody {
	return &EmployeesContractPostResponseBody{}
}

type EmployeesContractPostResponseBody []struct {
}

func (r *EmployeesContractPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/external-employees/{{.employeeReference}}/contracts", r.PathParams())
	return &u, err
}

func (r *EmployeesContractPostRequest) Do() (EmployeesContractPostResponseBody, error) {
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
