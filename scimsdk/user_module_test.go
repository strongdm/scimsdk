package scimsdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"scimsdk/internal/api"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

const mockUsersPageSize = 2

func TestUsersServiceListIterator(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return an users iteartor when there's no pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		client := NewClient("token", nil)
		iterator := client.Users().List(context.Background(), nil)

		assert.NotNil(t, iterator)
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
	})

	t.Run("should return an users iterator when there's pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: mockUsersPageSize, Offset: 1}
		iterator := client.Users().List(context.Background(), opts)

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

	t.Run("should return an empty users iterator when the offset is greater than page size and users count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: mockUsersPageSize, Offset: 3}
		iterator := client.Users().List(context.Background(), opts)

		assert.NotNil(t, iterator)
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.IsEmpty())
	})

	t.Run("should return a list of users when the offset is greater than the page size", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		client := NewClient("token", nil)
		opts := &PaginationOptions{PageSize: 1, Offset: 2}
		iterator := client.Users().List(context.Background(), opts)

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

func mockedApiExecuteWithUserPageResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUsersPageResponseJSON())))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
	response := &http.Response{Body: reader}
	startIndex := request.URL.Query().Get("startIndex")
	if startIndex > request.URL.Query().Get("count") && startIndex > "2" {
		response.Body = emptyReader
	}
	return response, nil
}

func getUsersPageResponseJSON() string {
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
		getUserResponseJSON(),
		getUserResponseJSON(),
	)
}

func getUserResponseJSON() string {
	return `
		{
			"active": false,
			"displayName": "xxx",
			"emails": [
				{
					"primary": true,
					"value": "xxx@zzz.com"
				}
			],
			"groups": [],
			"id": "a-xxx",
			"meta": {
				"resourceType": "User",
				"location": "Users/a-xxx"
			},
			"name": {
				"familyName": "00",
				"formatted": "yyy xxx",
				"givenName": "yyy"
			},
			"schemas": [
				"X.0:yyy"
			],
			"userName": "xxx@zzz.com",
			"userType": "account"
		}
	`
}
