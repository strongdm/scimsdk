package sdmscim

import (
	"fmt"
	"net/http"
)

func executeHTTPRequest(request *http.Request, token string) (*http.Response, error) {
	httpClient := http.Client{}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return httpClient.Do(request)
}
