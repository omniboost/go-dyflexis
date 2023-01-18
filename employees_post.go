package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewEmployeesPostRequest() EmployeesPostRequest {
	r := EmployeesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesPostRequest struct {
	client      *Client
	queryParams *EmployeesPostRequestQueryParams
	pathParams  *EmployeesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeesPostRequestBody
}

func (r EmployeesPostRequest) NewQueryParams() *EmployeesPostRequestQueryParams {
	return &EmployeesPostRequestQueryParams{}
}

type EmployeesPostRequestQueryParams struct {
}

func (p EmployeesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesPostRequest) QueryParams() *EmployeesPostRequestQueryParams {
	return r.queryParams
}

func (r EmployeesPostRequest) NewPathParams() *EmployeesPostRequestPathParams {
	return &EmployeesPostRequestPathParams{}
}

type EmployeesPostRequestPathParams struct {
	EmployeeReference string `schema:"employeeReference"`
}

func (p *EmployeesPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"employeeReference": p.EmployeeReference,
	}
}

func (r *EmployeesPostRequest) PathParams() *EmployeesPostRequestPathParams {
	return r.pathParams
}

func (r *EmployeesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesPostRequest) Method() string {
	return r.method
}

func (r EmployeesPostRequest) NewRequestBody() EmployeesPostRequestBody {
	return EmployeesPostRequestBody{}
}

type EmployeesPostRequestBody Employee

func (r *EmployeesPostRequest) RequestBody() *EmployeesPostRequestBody {
	return &r.requestBody
}

func (r *EmployeesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *EmployeesPostRequest) SetRequestBody(body EmployeesPostRequestBody) {
	r.requestBody = body
}

func (r *EmployeesPostRequest) NewResponseBody() *EmployeesPostResponseBody {
	return &EmployeesPostResponseBody{}
}

type EmployeesPostResponseBody []struct {
}

func (r *EmployeesPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/external-employees/{{.employeeReference}}/", r.PathParams())
	return &u, err
}

func (r *EmployeesPostRequest) Do() (EmployeesPostResponseBody, error) {
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
