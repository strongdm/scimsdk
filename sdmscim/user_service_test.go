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

var MOCK_USERS_PAGE_SIZE = 2

func TestUsersServiceCreate(t *testing.T) {
	t.Run("should create a user when passing valid data", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("token")
		opts := serviceCreateOptions{Body: getValidCreateUser()}
		user, err := service.create(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when creating an user without token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("")
		opts := serviceCreateOptions{Body: getValidCreateUser()}
		user, err := service.create(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should return an user when creating using context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceCreateOptions{Body: getValidCreateUser()}
		user, err := service.create(ctx, &opts)

		assert.NotNil(t, user)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceCreateOptions{Body: getValidCreateUser()}
		user, err := service.create(ctx, &opts)

		assert.Nil(t, user)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceListIterator(t *testing.T) {
	t.Run("should return a list of users when passing the default pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		opts := newServiceListOptions(getDefaultPaginationOptions(), "")
		iterator := userService.listIterator(context.Background(), opts)

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

	t.Run("should return a list of users when passing pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		opts := serviceListOptions{PageSize: MOCK_USERS_PAGE_SIZE, Offset: 1}
		iterator := userService.listIterator(context.Background(), &opts)

		assert.NotNil(t, iterator)
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
	})

	t.Run("should return an empty list of users when the offset is greater than page size and users count", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		opts := serviceListOptions{PageSize: MOCK_USERS_PAGE_SIZE, Offset: 3}
		iterator := userService.listIterator(context.Background(), &opts)

		assert.NotNil(t, iterator)
		assert.Empty(t, iterator.Err())
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
	})

	t.Run("should return a list of users when the offset is greater than the page size", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		opts := serviceListOptions{PageSize: 1, Offset: 2}
		iterator := userService.listIterator(context.Background(), &opts)

		assert.NotNil(t, iterator)
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.Next())
		assert.NotNil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
		assert.Empty(t, iterator.Err())
	})
}

func TestUsersServiceList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of users when passing a valid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		users, haveNextPage, err := userService.list(context.Background(), &serviceListOptions{PageSize: MOCK_USERS_PAGE_SIZE})

		assert.NotNil(t, users)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("should return a bad request error when send an empty token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("")
		users, haveNextPage, err := userService.list(context.Background(), &serviceListOptions{})

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, "Bad request", err.Error())
	})

	t.Run("should return an users list when using context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		userService := newUserService("token")
		users, haveNextPage, err := userService.list(context.Background(), &serviceListOptions{})

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, users)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
	})

	t.Run("should return a context error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		userService := newUserService("token")
		users, haveNextPage, err := userService.list(context.Background(), &serviceListOptions{})

		assert.NotNil(t, ctx.Err())
		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the users count", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		_, haveNextPage, _ := userService.list(context.Background(), &serviceListOptions{PageSize: 3})

		assert.False(t, haveNextPage)
	})

	t.Run("should return zero users when the offset is greater than the pageSize and the users count", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserPageResponse)
		userService := newUserService("token")
		users, haveNextPage, err := userService.list(context.Background(), &serviceListOptions{PageSize: MOCK_USERS_PAGE_SIZE, Offset: 3})

		assert.Zero(t, len(users))
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
	})
}

func TestUsersServiceFind(t *testing.T) {
	t.Run("should return an user when passing a valid user id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("token")
		opts := serviceFindOptions{ID: "xxx"}
		user, err := service.find(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("")
		opts := serviceFindOptions{ID: "xxx"}
		user, err := service.find(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid user id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := newUserService("token")
		opts := serviceFindOptions{ID: "yyy"}
		user, err := service.find(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceFindOptions{ID: "xxx"}
		user, err := service.find(context.Background(), &opts)

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceFindOptions{ID: "xxx"}
		user, err := service.find(context.Background(), &opts)

		assert.NotNil(t, ctx.Err())
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceReplace(t *testing.T) {
	t.Run("should replace an user when passing a valid body", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("token")
		opts := serviceReplaceOptions{ID: MOCK_USER_ID, Body: convertPorcelainToReplaceUserRequest(MOCK_USER_ID, getValidReplaceUser())}
		user, err := service.replace(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("")
		opts := serviceReplaceOptions{ID: MOCK_USER_ID, Body: convertPorcelainToReplaceUserRequest(MOCK_USER_ID, getValidReplaceUser())}
		user, err := service.replace(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid user-id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := newUserService("token")
		opts := serviceReplaceOptions{ID: MOCK_USER_ID, Body: convertPorcelainToReplaceUserRequest(MOCK_USER_ID, getValidReplaceUser())}
		user, err := service.replace(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should replace an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceReplaceOptions{ID: MOCK_USER_ID, Body: convertPorcelainToReplaceUserRequest(MOCK_USER_ID, getValidReplaceUser())}
		user, err := service.replace(ctx, &opts)

		assert.NotNil(t, user)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceReplaceOptions{ID: MOCK_USER_ID, Body: convertPorcelainToReplaceUserRequest(MOCK_USER_ID, getValidReplaceUser())}
		user, err := service.replace(ctx, &opts)

		assert.Nil(t, user)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceUpdate(t *testing.T) {
	t.Run("should update an user when passing a valid id and an active state", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("token")
		opts := serviceUpdateOptions{ID: MOCK_USER_ID, Body: convertPorcelainToUpdateUserRequest(true)}
		ok, err := service.update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		service := newUserService("")
		opts := serviceUpdateOptions{ID: MOCK_USER_ID, Body: convertPorcelainToUpdateUserRequest(true)}
		ok, err := service.update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an empty user-id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := newUserService("token")
		opts := serviceUpdateOptions{ID: "", Body: convertPorcelainToUpdateUserRequest(true)}
		ok, err := service.update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should update an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceUpdateOptions{ID: MOCK_USER_ID, Body: convertPorcelainToUpdateUserRequest(true)}
		ok, err := service.update(ctx, &opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceUpdateOptions{ID: MOCK_USER_ID, Body: convertPorcelainToUpdateUserRequest(true)}
		ok, err := service.update(ctx, &opts)

		assert.False(t, ok)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceDelete(t *testing.T) {
	t.Run("should delete the user when passing a valid user-id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedUser)
		service := newUserService("token")
		opts := serviceDeleteOptions{ID: "xxx"}
		ok, err := service.delete(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedUser)
		service := newUserService("")
		opts := serviceDeleteOptions{ID: "xxx"}
		ok, err := service.delete(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid user-id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := newUserService("token")
		opts := serviceDeleteOptions{ID: "yyy"}
		ok, err := service.delete(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedUser)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceDeleteOptions{ID: "xxx"}
		ok, err := service.delete(ctx, &opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newUserService("token")
		opts := serviceDeleteOptions{ID: "xxx"}
		ok, err := service.delete(ctx, &opts)

		assert.False(t, ok)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
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

func mockedApiExecuteWithExpiredTimeout(request *http.Request) (*http.Response, error) {
	time.Sleep(10 * time.Millisecond)
	return nil, errors.New("context deadline exceeded")
}

func mockedApiExecuteWithUserResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUserResponseJSON())))
	response := &http.Response{Body: reader}
	return response, nil
}

func mockedApiExecuteWithUserNotFound(request *http.Request) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getNotFoundResponseJSON())))
	response := &http.Response{Body: reader, StatusCode: 404}
	return response, nil
}

func mockedApiExecuteDeletedUser(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	response := &http.Response{StatusCode: 204}
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

func getNotFoundResponseJSON() string {
	return `
		{
			"schemas": ["X.0:Error"],
			"detail": "Resource yyy not found.",
			"status": "404"
		}
	`
}

func getValidCreateUser() *CreateUserBody {
	return &CreateUserBody{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}

func getValidReplaceUser() *ReplaceUserBody {
	return &ReplaceUserBody{
		UserName:   "xxx",
		GivenName:  "yyy",
		FamilyName: "zzz",
		Active:     true,
	}
}