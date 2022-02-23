package sdmscim

import (
	"errors"
	"fmt"
	"net/http"
)

func executeHTTPRequest(request *http.Request, token string) (*http.Response, error) {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := doRequest(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(getResponseDetails(response.Body))
	}
	return response, nil
}

func doRequest(request *http.Request) (*http.Response, error) {
	httpClient := http.Client{}
	return httpClient.Do(request)
}
