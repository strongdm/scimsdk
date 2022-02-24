package sdmscim

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
	defaultAPIURL      = "https://app.strongdm.com/provisioning/generic/v2"
	defaultAPIPageSize = 5
)

func apiCreate(ctx context.Context, pathname string, token string, opts *serviceCreateOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname)
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

func apiList(ctx context.Context, pathname string, token string, opts *serviceListOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = prepareRequestQueryParams(request, opts)
	return executeSafeHTTPRequest(request, token)
}

func apiFind(ctx context.Context, pathname string, token string, opts *serviceFindOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname, "/", opts.ID)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return executeSafeHTTPRequest(request, token)
}

func apiReplace(ctx context.Context, pathname string, token string, opts *serviceReplaceOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname, "/", opts.ID)
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

func apiUpdate(ctx context.Context, pathname string, token string, opts *serviceUpdateOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname, "/", opts.ID)
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

func apiDelete(ctx context.Context, pathname string, token string, opts *serviceDeleteOptions) (*http.Response, error) {
	url := fmt.Sprint(opts.BaseAPIURL, "/", pathname, "/", opts.ID)
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

func prepareRequestQueryParams(request *http.Request, opts *serviceListOptions) string {
	query := url.Values{}
	query.Set("startIndex", fmt.Sprint(opts.Offset))
	query.Set("count", fmt.Sprint(opts.PageSize))
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
