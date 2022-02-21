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

func Execute(request *http.Request, token string) (*http.Response, error) {
	httpClient := http.Client{}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return httpClient.Do(request)
}

func getPathnameWithPagination(pathname string, offset int, pageLimit int) string {
	return fmt.Sprintf("%s?startIndex=%d&count=%d", pathname, offset+1, pageLimit)
}
