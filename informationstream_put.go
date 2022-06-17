package dyflexis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/omniboost/go-dyflexis/utils"
)

func (c *Client) NewInformationstreamPutRequest() InformationstreamPutRequest {
	r := InformationstreamPutRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InformationstreamPutRequest struct {
	client      *Client
	queryParams *InformationstreamPutRequestQueryParams
	pathParams  *InformationstreamPutRequestPathParams
	method      string
	headers     http.Header
	requestBody InformationstreamPutRequestBody
}

func (r InformationstreamPutRequest) NewQueryParams() *InformationstreamPutRequestQueryParams {
	return &InformationstreamPutRequestQueryParams{}
}

type InformationstreamPutRequestQueryParams struct {
}

func (p InformationstreamPutRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InformationstreamPutRequest) QueryParams() *InformationstreamPutRequestQueryParams {
	return r.queryParams
}

func (r InformationstreamPutRequest) NewPathParams() *InformationstreamPutRequestPathParams {
	return &InformationstreamPutRequestPathParams{}
}

type InformationstreamPutRequestPathParams struct {
	Key  string `schema:"key"`
	Date Date   `schema:"date"`
}

func (p *InformationstreamPutRequestPathParams) Params() map[string]string {
	return map[string]string{
		"key":  p.Key,
		"date": p.Date.String(),
	}
}

func (r *InformationstreamPutRequest) PathParams() *InformationstreamPutRequestPathParams {
	return r.pathParams
}

func (r *InformationstreamPutRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InformationstreamPutRequest) SetMethod(method string) {
	r.method = method
}

func (r *InformationstreamPutRequest) Method() string {
	return r.method
}

func (r InformationstreamPutRequest) NewRequestBody() InformationstreamPutRequestBody {
	return InformationstreamPutRequestBody{}
}

type InformationstreamPutRequestBody Informationstream

func (r *InformationstreamPutRequest) RequestBody() *InformationstreamPutRequestBody {
	return &r.requestBody
}

func (r *InformationstreamPutRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *InformationstreamPutRequest) SetRequestBody(body InformationstreamPutRequestBody) {
	r.requestBody = body
}

func (r *InformationstreamPutRequest) NewResponseBody() *InformationStreamPutResponseBody {
	return &InformationStreamPutResponseBody{}
}

type InformationStreamPutResponseBody struct {
	Date  string `json:"date"`
	Value int    `json:"value"`
}

func (r *InformationstreamPutRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/informationstream/{{.key}}/{{.date}}", r.PathParams())
	return &u, err
}

func (r *InformationstreamPutRequest) Do() (InformationStreamPutResponseBody, error) {
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
