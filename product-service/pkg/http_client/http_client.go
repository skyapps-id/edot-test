package http_client

import (
	"context"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type RestClient interface {
	Get(ctx context.Context, path string, header http.Header) (body []byte, statusCode int, err error)
	Post(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Put(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Patch(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
	Delete(ctx context.Context, path string, header http.Header, requestBody []byte) (body []byte, statusCode int, err error)
}

type restyClient struct {
	client *resty.Client
}

func NewRestClient(baseURL, token string) RestClient {
	c := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Static-Token", token).
		SetHeader("Content-Type", "application/json")
	return &restyClient{client: c}
}

func (r *restyClient) Get(ctx context.Context, path string, header http.Header) ([]byte, int, error) {
	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		Get(path)
	return responseValues(resp, err)
}

func (r *restyClient) Post(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		SetBody(requestBody).
		Post(path)
	return responseValues(resp, err)
}

func (r *restyClient) Put(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		SetBody(requestBody).
		Put(path)
	return responseValues(resp, err)
}

func (r *restyClient) Patch(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		SetBody(requestBody).
		Patch(path)
	return responseValues(resp, err)
}

func (r *restyClient) Delete(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		SetBody(requestBody).
		Delete(path)
	return responseValues(resp, err)
}

func headerToMap(header http.Header) map[string]string {
	result := make(map[string]string)
	for key, values := range header {
		if len(values) > 0 {
			result[key] = values[0]
		}
	}
	return result
}

func responseValues(resp *resty.Response, err error) ([]byte, int, error) {
	if err != nil {
		return nil, 0, err
	}
	return resp.Body(), resp.StatusCode(), nil
}
