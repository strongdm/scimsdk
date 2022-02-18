package api

import (
	"net/http"
)

func BaseList(token string, pathname string, offset int) (*http.Response, error) {
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
