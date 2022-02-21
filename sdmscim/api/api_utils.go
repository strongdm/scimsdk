package api

import (
	"fmt"
	"io"
	"net/http"
)

func List(token string, pathname string, offset int) (*http.Response, error) {
	request, err := scimRequest("GET", getPathnameWithPagination(pathname, offset, DEFAULT_USERS_PAGE_LIMIT), nil)
	if err != nil {
		return nil, err
	}
	response, err := request.Do(token)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func scimRequest(method string, pathname string, body io.Reader) (*APIClientRequest, error) {
	url := fmt.Sprintf("%s/%s", BASE_URL, pathname)
	request, err := http.NewRequest(method, url, body)
	return &APIClientRequest{request}, err
}
