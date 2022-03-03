package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultAPIURL        = "https://app.strongdm.com/provisioning/generic/v2"
	defaultAPIPageSize   = 5
	defaultAPIPageOffset = 1
)

func Create(ctx context.Context, pathname string, token string, opts *CreateOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname)
	body, err := json.Marshal(opts.Body)
	if err != nil {
		return nil, err
	}
	reader := ioutil.NopCloser(bytes.NewReader(body))
	request, err := createHTTPRequest(ctx, "POST", url, reader)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func List(ctx context.Context, pathname string, token string, opts *ListOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = prepareRequestQueryParams(request, opts)
	return executeSafeHTTPRequest(request, token)
}

func Find(ctx context.Context, pathname string, token string, opts *FindOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func Replace(ctx context.Context, pathname string, token string, opts *ReplaceOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	body, err := json.Marshal(opts.Body)
	if err != nil {
		return nil, err
	}
	reader := ioutil.NopCloser(bytes.NewReader(body))
	request, err := createHTTPRequest(ctx, "PUT", url, reader)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func Update(ctx context.Context, pathname string, token string, opts *UpdateOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	body, err := json.Marshal(opts.Body)
	if err != nil {
		return nil, err
	}
	reader := ioutil.NopCloser(bytes.NewReader(body))
	request, err := createHTTPRequest(ctx, "PATCH", url, reader)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func Delete(ctx context.Context, pathname string, token string, opts *DeleteOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	request, err := createHTTPRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func createHTTPRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	requestWithCTX := request.WithContext(ctx)
	return requestWithCTX, err
}

func prepareRequestQueryParams(request *http.Request, opts *ListOptions) string {
	query := url.Values{}
	query.Set("startIndex", fmt.Sprint(getPageOffset(opts.Offset)))
	query.Set("count", fmt.Sprint(getPageSize(opts.PageSize)))
	if opts.Filter != "" {
		query.Set("filter", fmt.Sprint(opts.Filter))
	}
	return query.Encode()
}

func getResponseDetails(body io.Reader) string {
	buff, err := io.ReadAll(body)
	if err != nil {
		return err.Error()
	}
	mappedResponse := make(map[string]interface{})
	err = json.Unmarshal(buff, &mappedResponse)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprint(mappedResponse["detail"])
}

func getPageOffset(customOffset int) int {
	if customOffset > 0 {
		return customOffset
	}
	return defaultAPIPageOffset
}

func getPageSize(customPageSize int) int {
	if customPageSize > 0 {
		return customPageSize
	}
	return defaultAPIPageSize
}

func getBaseURL(customBaseURL string) string {
	if customBaseURL != "" {
		return customBaseURL
	}
	return defaultAPIURL
}

func GetDefaultAPIPageSize() int {
	return defaultAPIPageSize
}
