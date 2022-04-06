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

type IAPI interface {
	Create(ctx context.Context, pathname string, token string, opts *CreateOptions) (*http.Response, error)
	List(ctx context.Context, pathname string, token string, opts *ListOptions) (*http.Response, error)
	Find(ctx context.Context, pathname string, token string, opts *FindOptions) (*http.Response, error)
	Replace(ctx context.Context, pathname string, token string, opts *ReplaceOptions) (*http.Response, error)
	Update(ctx context.Context, pathname string, token string, opts *UpdateOptions) (*http.Response, error)
	Delete(ctx context.Context, pathname string, token string, opts *DeleteOptions) (*http.Response, error)
	ExecuteHTTPRequest(request *http.Request) (*http.Response, error)
}

type API struct {
	internalExecuteHTTPRequest func(*http.Request) (*http.Response, error)
}

func NewAPI() IAPI {
	return &API{internalExecuteHTTPRequest}
}

func (api *API) SetInternalExecuteHTTPRequest(fn func(*http.Request) (*http.Response, error)) {
	api.internalExecuteHTTPRequest = fn
}

const (
	defaultAPIURL        = "https://app.strongdm.com/provisioning/generic/v2"
	defaultAPIPageSize   = 5
	defaultAPIPageOffset = 1
)

func (api *API) Create(ctx context.Context, pathname string, token string, opts *CreateOptions) (*http.Response, error) {
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
	return ExecuteSafeHTTPRequest(api, request, token)
}

func (api *API) List(ctx context.Context, pathname string, token string, opts *ListOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = prepareRequestQueryParams(opts)
	return ExecuteSafeHTTPRequest(api, request, token)
}

func (api *API) Find(ctx context.Context, pathname string, token string, opts *FindOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	request, err := createHTTPRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return ExecuteSafeHTTPRequest(api, request, token)
}

func (api *API) Replace(ctx context.Context, pathname string, token string, opts *ReplaceOptions) (*http.Response, error) {
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
	return ExecuteSafeHTTPRequest(api, request, token)
}

func (api *API) Update(ctx context.Context, pathname string, token string, opts *UpdateOptions) (*http.Response, error) {
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
	return ExecuteSafeHTTPRequest(api, request, token)
}

func (api *API) Delete(ctx context.Context, pathname string, token string, opts *DeleteOptions) (*http.Response, error) {
	url := fmt.Sprint(getBaseURL(opts.BaseAPIURL), "/", pathname, "/", opts.ID)
	request, err := createHTTPRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return ExecuteSafeHTTPRequest(api, request, token)
}

func createHTTPRequest(ctx context.Context, method string, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	requestWithCTX := request.WithContext(ctx)
	return requestWithCTX, err
}

func prepareRequestQueryParams(opts *ListOptions) string {
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
