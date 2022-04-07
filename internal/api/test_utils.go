package api

import (
	"net/http"
)

func NewMockAPI(internalExecuteHTTPRequest func(*http.Request) (*http.Response, error)) *apiImpl {
	return &apiImpl{internalExecuteHTTPRequest}
}
