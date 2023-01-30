package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewContractTypesGetRequest() ContractTypesGetRequest {
	r := ContractTypesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ContractTypesGetRequest struct {
	client      *Client
	queryParams *ContractTypesGetRequestQueryParams
	pathParams  *ContractTypesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ContractTypesGetRequestBody
}

func (r ContractTypesGetRequest) NewQueryParams() *ContractTypesGetRequestQueryParams {
	return &ContractTypesGetRequestQueryParams{}
}

type ContractTypesGetRequestQueryParams struct {
}

func (p ContractTypesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ContractTypesGetRequest) QueryParams() *ContractTypesGetRequestQueryParams {
	return r.queryParams
}

func (r ContractTypesGetRequest) NewPathParams() *ContractTypesGetRequestPathParams {
	return &ContractTypesGetRequestPathParams{}
}

type ContractTypesGetRequestPathParams struct {
}

func (p *ContractTypesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ContractTypesGetRequest) PathParams() *ContractTypesGetRequestPathParams {
	return r.pathParams
}

func (r *ContractTypesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ContractTypesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ContractTypesGetRequest) Method() string {
	return r.method
}

func (r ContractTypesGetRequest) NewRequestBody() ContractTypesGetRequestBody {
	return ContractTypesGetRequestBody{}
}

type ContractTypesGetRequestBody struct {
}

func (r *ContractTypesGetRequest) RequestBody() *ContractTypesGetRequestBody {
	return nil
}

func (r *ContractTypesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ContractTypesGetRequest) SetRequestBody(body ContractTypesGetRequestBody) {
	r.requestBody = body
}

func (r *ContractTypesGetRequest) NewResponseBody() *ContractTypesResponseBody {
	return &ContractTypesResponseBody{}
}

type ContractTypesResponseBody struct {
	ContractTypes ContractTypes
}

func (r *ContractTypesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/contract-types", r.PathParams())
	return &u, err
}

func (r *ContractTypesGetRequest) Do() (ContractTypesResponseBody, error) {
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
