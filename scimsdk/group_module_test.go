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

	"github.com/strongdm/scimsdk/internal/api"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

const mockGroupsPageSize = 2

func TestGroupServiceListIterator(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a groups iterator when there's no pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		client := NewClient("token", nil)
		iterator := client.Groups().List(context.Background(), nil)

		assert.NotNil(t, iterator)
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
	})

	t.Run("should return a groups iterator when there's pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: mockGroupsPageSize, Offset: 1}
		iterator := client.Groups().List(context.Background(), opts)

		assert.NotNil(t, iterator)
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.IsEmpty())
	})

	t.Run("should return an empty groups iterator iterator when the offset is greater than page size and the groups count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: mockGroupsPageSize, Offset: 3}
		iterator := client.Groups().List(context.Background(), opts)

		assert.NotNil(t, iterator)
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.IsEmpty())
	})

	t.Run("should return a groups iterator when the offset is greater than the page size", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: 1, Offset: 2}
		iterator := client.Groups().List(context.Background(), opts)

		assert.NotNil(t, iterator)
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.IsEmpty())
	})
}

func mockedApiExecuteWithGroupPageResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getGroupsPageResponseJSON())))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
	response := &http.Response{Body: reader}
	startIndex := request.URL.Query().Get("startIndex")
	if startIndex > request.URL.Query().Get("count") && startIndex > "2" {
		response.Body = emptyReader
	}
	return response, nil
}

func getGroupsPageResponseJSON() string {
	return fmt.Sprintf(`
		{
			"Resources": [
				%s, %s
			],
			"itemsPerPage": 2,
			"schemas": [
				"X.0:Response"
			],
			"startIndex": 0,
			"totalResults": 2
		}
		`,
		getGroupResponseJSON(),
		getGroupResponseJSON(),
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
