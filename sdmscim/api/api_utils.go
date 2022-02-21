package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func List(token string, pathname string, offset int, ctx context.Context) (*http.Response, error) {
	request, err := scimRequest("GET", getPathnameWithPagination(pathname, offset, DEFAULT_USERS_PAGE_LIMIT), nil, ctx)
	if err != nil {
		return nil, err
	}
	response, err := Execute(request, token)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func scimRequest(method string, pathname string, body io.Reader, ctx context.Context) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", BASE_URL, pathname)
	request, err := http.NewRequest(method, url, body)
	requestWithCTX := request.WithContext(ctx)
	return requestWithCTX, err
}
