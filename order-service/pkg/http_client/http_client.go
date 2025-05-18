package http_client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
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
	return r.call(ctx, http.MethodGet, path, header, nil)
}

func (r *restyClient) Post(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	return r.call(ctx, http.MethodPost, path, header, requestBody)
}

func (r *restyClient) Put(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	return r.call(ctx, http.MethodPut, path, header, requestBody)
}

func (r *restyClient) Patch(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	return r.call(ctx, http.MethodPatch, path, header, requestBody)
}

func (r *restyClient) Delete(ctx context.Context, path string, header http.Header, requestBody []byte) ([]byte, int, error) {
	return r.call(ctx, http.MethodDelete, path, header, requestBody)
}

func (r *restyClient) call(ctx context.Context, method, path string, header http.Header, requestBody []byte,
) (body []byte, status int, err error) {
	ctx, span := tracer.Define().Start(ctx, fmt.Sprintf("%s %s", method, path))
	defer span.End()

	resp, err := r.client.R().
		SetContext(ctx).
		SetHeaders(headerToMap(header)).
		SetHeaders(PopulateTraceparentHeadersFromOtelContext(span.SpanContext())).
		SetBody(requestBody).
		Execute(method, path)
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
