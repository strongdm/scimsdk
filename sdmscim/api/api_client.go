package api

import (
	"fmt"
	"net/http"
)

const (
	BASE_URL                  = "https://app.strongdm.com/provisioning/generic/v2"
	DEFAULT_GROUPS_PAGE_LIMIT = 5
	DEFAULT_USERS_PAGE_LIMIT  = 5
)

type APIClientRequest struct {
	*http.Request
}

func (apiReq APIClientRequest) Do(token string) (*http.Response, error) {
	apiReq.setAuthorization(token)
	client := http.Client{}
	response, err := client.Do(apiReq.Request)
	return response, err
}

func (apiReq APIClientRequest) setAuthorization(token string) {
	apiReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func getPathnameWithPagination(pathname string, offset int, pageLimit int) string {
	return fmt.Sprintf("%s?startIndex=%d&count=%d", pathname, offset+1, pageLimit)
}
