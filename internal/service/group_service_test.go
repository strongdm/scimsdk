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

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

const mockGroupsPageSize = 2
const mockGroupID = "xxx"

func TestGroupServiceCreate(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should create a group when passing valid data", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := CreateOptions{Body: nil}
		group, err := service.Create(context.Background(), &opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when creating a group without token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("")
		opts := CreateOptions{Body: nil}
		group, err := service.Create(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should return a group when creating using context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := CreateOptions{Body: nil}
		group, err := service.Create(ctx, &opts)

		assert.NotNil(t, group)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := CreateOptions{Body: nil}
		group, err := service.Create(ctx, &opts)

		assert.Nil(t, group)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of groups when there's no pagination options", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.NotNil(t, groups)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, groups, 2)
	})

	t.Run("should return a list of users when the page size is equal or lesser than the groups count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockGroupsPageSize})

		assert.NotNil(t, groups)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, groups, 2)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{})

		assert.Nil(t, groups)
		assert.False(t, haveNextPage)
		assert.NotNil(t, err)
	})

	t.Run("should return a list of groups when usign a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.List(ctx, &ListOptions{})

		assert.NotNil(t, groups)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Nil(t, ctx.Err())
		assert.Len(t, groups, 2)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		service := NewGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.List(ctx, &ListOptions{})

		assert.Nil(t, groups)
		assert.False(t, haveNextPage)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})

	t.Run("should return false in haveNextPage when the page size is greater than the groups count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("token")
		_, haveNextPage, _ := service.List(context.Background(), &ListOptions{PageSize: 3})

		assert.False(t, haveNextPage)
	})

	t.Run("should return zero groups when the offset is greater than the page size and the groups count", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := NewGroupService("token")
		groups, haveNextPage, err := service.List(context.Background(), &ListOptions{PageSize: mockGroupsPageSize, Offset: 3})

		assert.Zero(t, len(groups))
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
	})
}

func TestGroupServiceFind(t *testing.T) {
	t.Run("should return a group when passing a valid group id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := FindOptions{ID: mockGroupID}
		group, err := service.Find(context.Background(), &opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("")
		opts := FindOptions{ID: mockGroupID}
		group, err := service.Find(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid group id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupNotFound)
		service := NewGroupService("token")
		opts := FindOptions{ID: "yyy"}
		group, err := service.Find(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := FindOptions{ID: mockGroupID}
		group, err := service.Find(context.Background(), &opts)

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := FindOptions{ID: mockGroupID}
		group, err := service.Find(context.Background(), &opts)

		assert.NotNil(t, ctx.Err())
		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceReplace(t *testing.T) {
	t.Run("should replace a group when passing a valid group id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := &ReplaceOptions{mockGroupID, nil, ""}
		group, err := service.Replace(context.Background(), opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupNotFound)
		opts := &ReplaceOptions{"", nil, ""}
		service := NewGroupService("token")
		group, err := service.Replace(context.Background(), opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupNotFound)
		opts := &ReplaceOptions{mockGroupID, nil, ""}
		service := NewGroupService("")
		group, err := service.Replace(context.Background(), opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := &ReplaceOptions{"yyy", nil, ""}
		group, err := service.Replace(ctx, opts)

		assert.NotNil(t, group)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := &ReplaceOptions{"yyy", nil, ""}
		group, err := service.Replace(ctx, opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.NotNil(t, ctx.Err())
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupsServiceUpdate(t *testing.T) {
	t.Run("should update a group when passing a valid id and replace name body", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should update a group when passing a valid id and add members body", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should update a group when passing a valid id and replace members body", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should update a group when passing a valid id and remove members body", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := NewGroupService("")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an empty group-id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupNotFound)
		service := NewGroupService("token")
		opts := UpdateOptions{ID: "", Body: nil}
		ok, err := service.Update(context.Background(), &opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should update an group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(ctx, &opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := NewGroupService("token")
		opts := UpdateOptions{ID: mockGroupID, Body: nil}
		ok, err := service.Update(ctx, &opts)

		assert.False(t, ok)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceDelete(t *testing.T) {
	t.Run("should delete the group when passing a valid token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedGroup)
		service := NewGroupService("token")
		opts := DeleteOptions{ID: mockGroupID}
		ok, err := service.Delete(context.Background(), &opts)

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedGroup)
		service := NewGroupService("")
		opts := &DeleteOptions{ID: mockGroupID}
		ok, err := service.Delete(context.Background(), opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithGroupNotFound)
		service := NewGroupService("token")
		opts := &DeleteOptions{ID: mockGroupID}
		ok, err := service.Delete(context.Background(), opts)

		assert.False(t, ok)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should delete the group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteDeletedGroup)
		service := NewGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		opts := &DeleteOptions{ID: mockGroupID}
		ok, err := service.Delete(ctx, opts)

		assert.True(t, ok)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should delete the group when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(api.ExecuteHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		service := NewGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		opts := &DeleteOptions{ID: mockGroupID}
		ok, err := service.Delete(ctx, opts)

		assert.False(t, ok)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
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
