package module

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/strongdm/scimsdk/internal/service"
	"github.com/strongdm/scimsdk/models"
)

const mockUsersPageSize = 2

func TestUsersServiceListIterator(t *testing.T) {
	t.Run("should return an users iteartor when there's no pagination options", func(t *testing.T) {
		mockApi := getMockedAPI(mockedApiExecuteWithUserPageResponse)
		serviceApi := service.NewUserService(mockApi, "token")
		module := NewMockUserModule(serviceApi)
		iterator := module.List(context.Background(), nil)
		assertT := assert.New(t)

		assertT.NotNil(iterator)
		assertT.Nil(iterator.Err())
		assertT.True(iterator.Next())
		assertT.NotNil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.True(iterator.Next())
		assertT.NotNil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.False(iterator.Next())
		assertT.Nil(iterator.Err())
	})

	t.Run("should return an users iterator when there's pagination options", func(t *testing.T) {
		mockApi := getMockedAPI(mockedApiExecuteWithUserPageResponse)
		serviceApi := service.NewUserService(mockApi, "token")
		module := NewMockUserModule(serviceApi)
		opts := &models.PaginationOptions{PageSize: mockUsersPageSize, Offset: 1}
		iterator := module.List(context.Background(), opts)
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

	t.Run("should return an empty users iterator when the offset is greater than page size and users count", func(t *testing.T) {
		mockApi := getMockedAPI(mockedApiExecuteWithUserPageResponse)
		serviceApi := service.NewUserService(mockApi, "token")
		module := NewMockUserModule(serviceApi)
		opts := &models.PaginationOptions{PageSize: mockUsersPageSize, Offset: 3}
		iterator := module.List(context.Background(), opts)
		assertT := assert.New(t)

		assertT.NotNil(iterator)
		assertT.False(iterator.Next())
		assertT.Nil(iterator.Err())
		assertT.Nil(iterator.Value())
		assertT.Nil(iterator.Err())
		assertT.True(iterator.IsEmpty())
	})
}

func mockedApiExecuteWithUserPageResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	pageCount := request.URL.Query().Get("count")
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUsersPageResponseJSON(pageCount))))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte(getEmptyUsersPageResponseJSON(pageCount))))
	response := &http.Response{Body: reader}
	startIndex := request.URL.Query().Get("startIndex")
	if startIndex > pageCount && pageCount >= "2" {
		response.Body = emptyReader
	}
	return response, nil
}

func getUsersPageResponseJSON(pageCount string) string {
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
		getUserResponseJSON(),
		getUserResponseJSON(),
		pageCount,
	)
}

func getEmptyUsersPageResponseJSON(pageCount string) string {
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
