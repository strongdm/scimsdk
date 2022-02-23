package sdmscim

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultAPIURL      = "https://app.strongdm.com/provisioning/generic/v2"
	defaultAPIPageSize = 5
)

func apiList(token string, pathname string, opts *apiListOptions, ctx context.Context) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname)
	request, err := createHTTPRequest("GET", url, nil, ctx)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = prepareRequestQueryParams(request, opts)
	response, err := executeHTTPRequest(request, token)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func createHTTPRequest(method string, url string, body io.Reader, ctx context.Context) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	requestWithCTX := request.WithContext(ctx)
	return requestWithCTX, err
}

func prepareRequestQueryParams(request *http.Request, opts *apiListOptions) string {
	query := url.Values{}
	query.Set("startIndex", fmt.Sprint(opts.Offset))
	query.Set("count", fmt.Sprint(opts.PageSize))
	if opts.Filter != "" {
		query.Set("filter", fmt.Sprint(opts.Filter))
	}
	return query.Encode()
}
