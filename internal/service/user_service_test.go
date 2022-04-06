package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/strongdm/scimsdk/internal/api"

	"github.com/stretchr/testify/assert"
)

const (
	mockUsersPageSize = 2
	mockUserID        = "xxx"
	defaultUserSchema = "urn:ietf:params:scim:schemas:core:2.0:User"
)

func TestUsersServiceCreate(t *testing.T) {
	t.Run("should create a user when passing valid data", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "token")
		user, err := service.Create(context.Background(), &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.NotNil(user)
		assertT.Nil(err)
	})

	t.Run("should return an error when creating an user without token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "")
		user, err := service.Create(context.Background(), &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(err)
	})

	t.Run("should return an user when creating using context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Create(ctx, &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.NotNil(user)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Create(ctx, &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestUsersServiceList(t *testing.T) {
	t.Run("should return a list of users when there's no pagination options", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		service := NewUserService(mock, "token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.NotNil(users)
		assertT.False(haveNextPage)
		assertT.Nil(err)
		assertT.Len(users, 2)
	})

	t.Run("should return a list of users when the page size is equal or lesser than the users count", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		service := NewUserService(mock, "token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockUsersPageSize})
		assertT := assert.New(t)

		assertT.NotNil(users)
		assertT.True(haveNextPage)
		assertT.Nil(err)
		assertT.Len(users, 2)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		service := NewUserService(mock, "")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.Nil(users)
		assertT.False(haveNextPage)
		assertT.NotNil(err)
	})

	t.Run("should return a list of users when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.Nil(ctx.Err())
		assertT.NotNil(users)
		assertT.False(haveNextPage)
		assertT.Nil(err)
		assertT.Len(users, 2)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.Nil(users)
		assertT.False(haveNextPage)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the users count", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		service := NewUserService(mock, "token")
		_, haveNextPage, _ := service.List(context.Background(), &ListOptions{PageSize: 3})
		assertT := assert.New(t)

		assertT.False(haveNextPage)
	})

	t.Run("should return zero users when the offset is greater than the page size and the users count", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserPageResponse)
		service := NewUserService(mock, "token")
		users, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockUsersPageSize, Offset: 3})
		assertT := assert.New(t)

		assertT.Zero(len(users))
		assertT.False(haveNextPage)
		assertT.Nil(err)
	})
}

func TestUsersServiceFind(t *testing.T) {
	t.Run("should return an user when passing a valid user id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "token")
		user, err := service.Find(context.Background(), &FindOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.NotNil(user)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "")
		user, err := service.Find(context.Background(), &FindOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an invalid user id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserNotFound)
		service := NewUserService(mock, "token")
		user, err := service.Find(context.Background(), &FindOptions{ID: "yyy"})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should return an user when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Find(context.Background(), &FindOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.Nil(ctx.Err())
		assertT.NotNil(user)
		assertT.Nil(err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Find(context.Background(), &FindOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.NotNil(ctx.Err())
		assertT.Nil(user)
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestUsersServiceReplace(t *testing.T) {
	t.Run("should replace an user when passing a valid user id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "token")
		user, err := service.Replace(context.Background(), &ReplaceOptions{mockUserID, nil, ""})
		assertT := assert.New(t)

		assertT.NotNil(user)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid user id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserNotFound)
		service := NewUserService(mock, "token")
		user, err := service.Replace(context.Background(), &ReplaceOptions{mockUserID, nil, ""})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "")
		user, err := service.Replace(context.Background(), &ReplaceOptions{mockUserID, nil, ""})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(err)
	})

	t.Run("should replace an user when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Replace(ctx, &ReplaceOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.NotNil(user)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		user, err := service.Replace(ctx, &ReplaceOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.Nil(user)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestUsersServiceUpdate(t *testing.T) {
	t.Run("should update an user when passing a valid id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		service := NewUserService(mock, "")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an empty user-id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserNotFound)
		service := NewUserService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: "", Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should update an user when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		ok, err := service.Update(ctx, &UpdateOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		ok, err := service.Update(ctx, &UpdateOptions{ID: mockUserID, Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestUsersServiceDelete(t *testing.T) {
	t.Run("should delete the user when passing a valid user-id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteDeletedUser)
		service := NewUserService(mock, "token")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteDeletedUser)
		service := NewUserService(mock, "")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an invalid user-id", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithUserNotFound)
		service := NewUserService(mock, "token")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: "yyy"})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteDeletedUser)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		ok, err := service.Delete(ctx, &DeleteOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		mock := api.NewAPI()
		mock.(*api.API).SetInternalExecuteHTTPRequest(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewUserService(mock, "token")
		ok, err := service.Delete(ctx, &DeleteOptions{ID: mockUserID})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
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
	if startIndex > request.URL.Query().Get("count") {
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

func getNotFoundResponseJSON() string {
	return `
		{
			"schemas": ["X.0:Error"],
			"detail": "Resource yyy not found.",
			"status": "404"
		}
	`
}
