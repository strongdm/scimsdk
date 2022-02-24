package sdmscim

import (
	"errors"
	"fmt"
	"net/http"
)

// executeSafeHTTPRequest controls the executeHTTPRequest response passing an
// authenticated http request and treating the http response.
func executeSafeHTTPRequest(request *http.Request, token string) (*http.Response, error) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := executeHTTPRequest(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(getResponseDetails(response.Body))
	}
	return response, nil
}

func executeHTTPRequest(request *http.Request) (*http.Response, error) {
	httpClient := http.Client{}
	return httpClient.Do(request)
}
