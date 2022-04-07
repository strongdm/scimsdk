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

const mockGroupsPageSize = 2
const mockGroupID = "xxx"

func TestGroupServiceCreate(t *testing.T) {
	t.Run("should create a group when passing valid data", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewUserService(mock, "token")
		group, err := service.Create(context.Background(), &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.NotNil(group)
		assertT.Nil(err)
	})

	t.Run("should return an error when creating a group without token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "")
		group, err := service.Create(context.Background(), &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
	})

	t.Run("should return a group when creating using context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Create(ctx, &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.NotNil(group)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Create(ctx, &CreateOptions{Body: nil})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestGroupServiceList(t *testing.T) {
	t.Run("should return a list of groups when there's no pagination options", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.NotNil(groups)
		assertT.False(haveNextPage)
		assertT.Nil(err)
		assertT.Len(groups, 2)
	})

	t.Run("should return a list of users when the page size is equal or lesser than the groups count", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockGroupsPageSize})
		assertT := assert.New(t)

		assertT.NotNil(groups)
		assertT.True(haveNextPage)
		assertT.Nil(err)
		assertT.Len(groups, 2)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{})
		assertT := assert.New(t)

		assertT.Nil(groups)
		assertT.False(haveNextPage)
		assertT.NotNil(err)
	})

	t.Run("should return a list of groups when usign a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.List(ctx, &ListOptions{})
		assertT := assert.New(t)

		assertT.NotNil(groups)
		assertT.False(haveNextPage)
		assertT.Nil(err)
		assertT.Nil(ctx.Err())
		assertT.Len(groups, 2)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		service := NewGroupService(mock, "token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.List(ctx, &ListOptions{})
		assertT := assert.New(t)

		assertT.Nil(groups)
		assertT.False(haveNextPage)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the groups count", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "token")
		_, haveNextPage, _ := service.List(context.Background(), &ListOptions{PageSize: 3})
		assertT := assert.New(t)

		assertT.False(haveNextPage)
	})

	t.Run("should return zero groups when the offset is greater than the page size and the groups count", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService(mock, "token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockGroupsPageSize, Offset: 3})
		assertT := assert.New(t)

		assertT.Zero(len(groups))
		assertT.False(haveNextPage)
		assertT.Nil(err)
	})
}

func TestGroupServiceFind(t *testing.T) {
	t.Run("should return a group when passing a valid group id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		group, err := service.Find(context.Background(), &FindOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.NotNil(group)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "")
		group, err := service.Find(context.Background(), &FindOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an invalid group id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupNotFound)
		service := NewGroupService(mock, "token")
		group, err := service.Find(context.Background(), &FindOptions{ID: "yyy"})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should return a group when using a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Find(context.Background(), &FindOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.Nil(ctx.Err())
		assertT.NotNil(group)
		assertT.Nil(err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Find(context.Background(), &FindOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.NotNil(ctx.Err())
		assertT.Nil(group)
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestGroupServiceReplace(t *testing.T) {
	t.Run("should replace a group when passing a valid group id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		group, err := service.Replace(context.Background(), &ReplaceOptions{mockGroupID, nil, ""})
		assertT := assert.New(t)

		assertT.NotNil(group)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupNotFound)
		service := NewGroupService(mock, "token")
		group, err := service.Replace(context.Background(), &ReplaceOptions{"", nil, ""})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupNotFound)
		service := NewGroupService(mock, "")
		group, err := service.Replace(context.Background(), &ReplaceOptions{mockGroupID, nil, ""})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Replace(ctx, &ReplaceOptions{"yyy", nil, ""})
		assertT := assert.New(t)

		assertT.NotNil(group)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		group, err := service.Replace(ctx, &ReplaceOptions{"yyy", nil, ""})
		assertT := assert.New(t)

		assertT.Nil(group)
		assertT.NotNil(err)
		assertT.NotNil(ctx.Err())
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestGroupsServiceUpdate(t *testing.T) {
	t.Run("should update a group when passing a valid id and replace name body", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should update a group when passing a valid id and add members body", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should update a group when passing a valid id and replace members body", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should update a group when passing a valid id and remove members body", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		service := NewGroupService(mock, "")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an empty group-id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupNotFound)
		service := NewGroupService(mock, "token")
		ok, err := service.Update(context.Background(), &UpdateOptions{ID: "", Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should update an group when using a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		ok, err := service.Update(ctx, &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService(mock, "token")
		ok, err := service.Update(ctx, &UpdateOptions{ID: mockGroupID, Body: nil})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
	})
}

func TestGroupServiceDelete(t *testing.T) {
	t.Run("should delete the group when passing a valid token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteDeletedGroup)
		service := NewGroupService(mock, "token")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(err)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteDeletedGroup)
		service := NewGroupService(mock, "")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithGroupNotFound)
		service := NewGroupService(mock, "token")
		ok, err := service.Delete(context.Background(), &DeleteOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(err)
		assertT.Contains(err.Error(), "not found")
	})

	t.Run("should delete the group when using a context with timeout", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteDeletedGroup)
		service := NewGroupService(mock, "token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		ok, err := service.Delete(ctx, &DeleteOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.True(ok)
		assertT.Nil(ctx.Err())
		assertT.Nil(err)
	})

	t.Run("should delete the group when the context timeout exceed", func(t *testing.T) {
		mock := api.NewMockAPI(mockedApiExecuteWithExpiredTimeout)
		service := NewGroupService(mock, "token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		ok, err := service.Delete(ctx, &DeleteOptions{ID: mockGroupID})
		assertT := assert.New(t)

		assertT.False(ok)
		assertT.NotNil(ctx.Err())
		assertT.NotNil(err)
		assertT.Equal("context deadline exceeded", ctx.Err().Error())
		assertT.Equal("context deadline exceeded", err.Error())
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
	if startIndex > pageCount {
		response.Body = emptyReader
	}
	return response, nil
}

func mockedApiExecuteWithGroupResponse(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getGroupResponseJSON())))
	response := &http.Response{Body: reader}
	return response, nil
}

func mockedApiExecuteWithGroupNotFound(request *http.Request) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getNotFoundResponseJSON())))
	response := &http.Response{Body: reader, StatusCode: 404}
	return response, nil
}

func mockedApiExecuteDeletedGroup(request *http.Request) (*http.Response, error) {
	token := extractAuthorizationToken(request.Header.Get("Authorization"))
	if token == "" {
		return nil, errors.New("Bad request")
	}
	response := &http.Response{StatusCode: 204}
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
