package api

import (
	"errors"
	"fmt"
	"net/http"
)

// ExecuteSafeHTTPRequest controls the executeHTTPRequest response passing an
// authenticated http request and treating the http response.
func ExecuteSafeHTTPRequest(api *API, request *http.Request, token string) (*http.Response, error) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := api.ExecuteHTTPRequest(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(getResponseDetails(response.Body))
	}
	return response, nil
}

func (api API) ExecuteHTTPRequest(request *http.Request) (*http.Response, error) {
	return api.internalExecuteHTTPRequest(request)
}

func internalExecuteHTTPRequest(req *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(req)
}
