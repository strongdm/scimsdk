package scimsdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockGroupsPageSize = 2

func TestGroupServiceListIterator(t *testing.T) {
	t.Run("should return a groups iterator when there's no pagination options", func(t *testing.T) {
		client := NewMockClient(mockedApiExecuteWithGroupPageResponse, "token")
		iterator := client.Groups().List(context.Background(), nil)
		assertT := assert.New(t)

		assertT.NotNil(iterator)
		assertT.Nil(iterator.Err())
		assertT.True(iterator.Next())
		assertT.NotNil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.True(iterator.Next())
		assertT.Nil(iterator.Err())
		assertT.NotNil(iterator.Value())
		assertT.False(iterator.Next())
		assertT.Nil(iterator.Err())
	})

	t.Run("should return a groups iterator when there's pagination options", func(t *testing.T) {
		client := NewMockClient(mockedApiExecuteWithGroupPageResponse, "token")
		opts := &PaginationOptions{PageSize: mockGroupsPageSize, Offset: 1}
		iterator := client.Groups().List(context.Background(), opts)
		assertT := assert.New(t)

		assertT.NotNil(iterator)
		assertT.True(iterator.Next())
		assertT.Nil(iterator.Err())
		assertT.NotNil(iterator.Value())
		assertT.True(iterator.Next())
		assertT.Nil(iterator.Err())
		assertT.NotNil(iterator.Value())
		assertT.False(iterator.Next())
		assertT.Nil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.True(iterator.IsEmpty())
	})

	t.Run("should return an empty groups iterator iterator when the offset is greater than page size and the groups count", func(t *testing.T) {
		client := NewMockClient(mockedApiExecuteWithGroupPageResponse, "token")
		opts := &PaginationOptions{PageSize: mockGroupsPageSize, Offset: 3}
		iterator := client.Groups().List(context.Background(), opts)
		assertT := assert.New(t)

		assertT.NotNil(iterator)
		assertT.False(iterator.Next())
		assertT.Nil(iterator.Err())
		assertT.Nil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.True(iterator.IsEmpty())
	})
}

func mockedApiExecuteWithGroupPageResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	pageCount := request.URL.Query().Get("count")
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getGroupsPageResponseJSON(pageCount))))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte(getEmptyGroupsPageResponseJSON(pageCount))))
	response := &http.Response{Body: reader}
	startIndex := request.URL.Query().Get("startIndex")
	if startIndex > pageCount && pageCount >= "2" {
		response.Body = emptyReader
	}
	return response, nil
}

func getGroupsPageResponseJSON(pageCount string) string {
	return fmt.Sprintf(`
		{
			"Resources": [
				%s, %s
			],
			"itemsPerPage": %s,
			"schemas": [
				"X.0:Response"
			],
			"startIndex": 0,
			"totalResults": 2
		}
		`,
		getGroupResponseJSON(),
		getGroupResponseJSON(),
		pageCount,
	)
}

func getEmptyGroupsPageResponseJSON(pageCount string) string {
	return fmt.Sprintf(`
		{
			"Resources": [],
			"itemsPerPage": %s,
			"schemas": [
				"X.0:Response"
			],
			"startIndex": 0,
			"totalResults": 0
		}
		`,
		pageCount,
	)
}

func getGroupResponseJSON() string {
	return `
		{
			"schemas": ["X.0:Group"],
      "displayName": "xxx",
      "id": "yyy",
      "members": [],
      "meta": {
        "resourceType": "Group",
        "location": "Groups/xxx"
      }
		}
	`
}

func extractAuthorizationToken(authHeaderValue string) string {
	token := strings.Split(authHeaderValue, "Bearer")[1]
	return strings.TrimSpace(token)
}
