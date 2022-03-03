package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"scimsdk/internal/api"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

const (
	mockUsersPageSize = 2
	mockUserID        = "xxx"
	defaultUserSchema = "urn:ietf:params:scim:schemas:core:2.0:User"
)

func TestUsersServiceCreate(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should create a user when passing valid data", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("token")
		opts := CreateOptions{Body: nil}
		user, err := service.Create(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when creating an user without token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("")
		opts := CreateOptions{Body: nil}
		user, err := service.Create(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should return an user when creating using context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := CreateOptions{Body: nil}
		user, err := service.Create(ctx, &opts)

		assert.NotNil(t, user)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := CreateOptions{Body: nil}
		user, err := service.Create(ctx, &opts)

		assert.Nil(t, user)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of users when there's no pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		service := NewUserService("token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.NotNil(t, users)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("should return a list of users when the page size is equal or lesser than the users count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		service := NewUserService("token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockUsersPageSize})

		assert.NotNil(t, users)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		service := NewUserService("")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.NotNil(t, err)
	})

	t.Run("should return a list of users when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, users)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, users, 2)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the users count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		service := NewUserService("token")
		_, haveNextPage, _ := service.List(context.Background(), &ListOptions{PageSize: 3})

		assert.False(t, haveNextPage)
	})

	t.Run("should return zero users when the offset is greater than the page size and the users count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserPageResponse)
		service := NewUserService("token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockUsersPageSize, Offset: 3})

		assert.Zero(t, len(users))
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
	})
}

func TestUsersServiceFind(t *testing.T) {
	t.Run("should return an user when passing a valid user id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("token")
		opts := FindOptions{ID: mockUserID}
		user, err := service.Find(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("")
		opts := FindOptions{ID: mockUserID}
		user, err := service.Find(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid user id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := NewUserService("token")
		opts := FindOptions{ID: "yyy"}
		user, err := service.Find(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := FindOptions{ID: mockUserID}
		user, err := service.Find(context.Background(), &opts)

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := FindOptions{ID: mockUserID}
		user, err := service.Find(context.Background(), &opts)

		assert.NotNil(t, ctx.Err())
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceReplace(t *testing.T) {
	t.Run("should replace an user when passing a valid user id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("token")
		opts := ReplaceOptions{mockUserID, nil, ""}
		user, err := service.Replace(context.Background(), &opts)

		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid user id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := NewUserService("token")
		opts := ReplaceOptions{mockUserID, nil, ""}
		user, err := service.Replace(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("")
		opts := ReplaceOptions{mockUserID, nil, ""}
		user, err := service.Replace(context.Background(), &opts)

		assert.Nil(t, user)
		assert.NotNil(t, err)
	})

	t.Run("should replace an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := ReplaceOptions{ID: mockUserID, Body: nil}
		user, err := service.Replace(ctx, &opts)

		assert.NotNil(t, user)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := ReplaceOptions{ID: mockUserID, Body: nil}
		user, err := service.Replace(ctx, &opts)

		assert.Nil(t, user)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceUpdate(t *testing.T) {
	t.Run("should update an user when passing a valid id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("token")
		opts := UpdateOptions{ID: mockUserID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		service := NewUserService("")
		opts := UpdateOptions{ID: mockUserID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an empty user-id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := NewUserService("token")
		opts := UpdateOptions{ID: "", Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should update an user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := UpdateOptions{ID: mockUserID, Body: nil}
		ok, err := service.Update(ctx, &opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := UpdateOptions{ID: mockUserID, Body: nil}
		ok, err := service.Update(ctx, &opts)

		assert.False(t, ok)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestUsersServiceDelete(t *testing.T) {
	t.Run("should delete the user when passing a valid user-id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedUser)
		service := NewUserService("token")
		opts := DeleteOptions{ID: mockUserID}
		ok, err := service.Delete(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedUser)
		service := NewUserService("")
		opts := DeleteOptions{ID: mockUserID}
		ok, err := service.Delete(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid user-id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithUserNotFound)
		service := NewUserService("token")
		opts := DeleteOptions{ID: "yyy"}
		ok, err := service.Delete(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedUser)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := DeleteOptions{ID: mockUserID}
		ok, err := service.Delete(ctx, &opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService("token")
		opts := DeleteOptions{ID: mockUserID}
		ok, err := service.Delete(ctx, &opts)

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
