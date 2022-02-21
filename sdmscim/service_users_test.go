package sdmscim

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sdmscim/sdmscim/api"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestUsersServiceList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of users", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		userService := newUserService("", ctx)
		monkey.Patch(api.Execute, mockedApiExecuteWithUserResponse)
		users, haveNextPage, err := userService.list(0)

		assert.NotNil(t, users)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
	})

	t.Run("should return a bad request error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		userService := newUserService("", ctx)
		monkey.Patch(api.Execute, mockedApiExecuteWithBadRequestError)
		users, haveNextPage, err := userService.list(0)

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, "Bad request", err.Error())
	})

	t.Run("should return a permission denied message", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		userService := newUserService("", ctx)
		monkey.Patch(api.Execute, mockedApiExecuteWithPermissionDenied)
		users, haveNextPage, err := userService.list(0)

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, "invalid character 'p' looking for beginning of value", err.Error())
	})

	t.Run("should return a context error", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		userService := newUserService("", ctx)
		monkey.Patch(api.Execute, mockedApiExecuteWithExpiredTimeout)
		_, _, err := userService.list(0)

		assert.NotNil(t, ctx.Err())
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func mockedApiExecuteWithUserResponse(request *http.Request, token string) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUsersSCIMResponseJSON())))
	return &http.Response{
		Body: reader,
	}, nil
}

func mockedApiExecuteWithBadRequestError(request *http.Request, token string) (*http.Response, error) {
	return nil, errors.New("Bad request")
}

func mockedApiExecuteWithPermissionDenied(request *http.Request, token string) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte("permission denied: access denied")))
	return &http.Response{
		Body: reader,
	}, nil
}

func mockedApiExecuteWithExpiredTimeout(request *http.Request, token string) (*http.Response, error) {
	time.Sleep(101 * time.Millisecond)
	return nil, errors.New("context deadline exceeded")
}

func getUsersSCIMResponseJSON() string {
	return fmt.Sprintf(`
		{
			"Resources": [
				%s, %s, %s, %s, %s
			],
			"itemsPerPage": 1,
			"schemas": [
				"X.0:Response"
			],
			"startIndex": 0,
			"totalResults": 1
		}
		`,
		GetUserResponseDTOJson(),
		GetUserResponseDTOJson(),
		GetUserResponseDTOJson(),
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
