package sdmscim

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestUsersServiceList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of users when passing a valid token", func(t *testing.T) {
		userService := newUserService("token", context.Background())
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		users, haveNextPage, err := userService.list(&ModuleListOptions{PageSize: 1})

		assert.NotNil(t, users)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
	})

	t.Run("should return a bad request error when send an empty token", func(t *testing.T) {
		userService := newUserService("", context.Background())
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		users, haveNextPage, err := userService.list(&ModuleListOptions{})

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, "Bad request", err.Error())
	})

	t.Run("should return a context error when the context timeout exceed", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		userService := newUserService("", ctx)
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		_, _, err := userService.list(&ModuleListOptions{})

		assert.NotNil(t, ctx.Err())
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the users count", func(t *testing.T) {
		userService := newUserService("token", context.Background())
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		_, haveNextPage, _ := userService.list(&ModuleListOptions{PageSize: 3})

		assert.False(t, haveNextPage)
	})

	t.Run("should return zero users when the offset is greater than the pageSize and the users count", func(t *testing.T) {
		userService := newUserService("token", context.Background())
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		users, haveNextPage, err := userService.list(&ModuleListOptions{PageSize: 3, Offset: 4})

		assert.Zero(t, len(users))
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
	})
}

func mockedApiExecuteWithUserResponse(request *http.Request, token string) (*http.Response, error) {
	if token == "" {
		return nil, errors.New("Bad request")
	}
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUsersPageResponseJSON())))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
	response := &http.Response{Body: reader}
	if request.URL.Query().Get("startIndex") > request.URL.Query().Get("count") {
		response.Body = emptyReader
	}
	return response, nil
}

func mockedApiExecuteWithExpiredTimeout(request *http.Request, token string) (*http.Response, error) {
	time.Sleep(101 * time.Millisecond)
	return nil, errors.New("context deadline exceeded")
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
		GetUserResponseDTOJson(),
		GetUserResponseDTOJson(),
	)
}

func GetUserResponseDTOJson() string {
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
