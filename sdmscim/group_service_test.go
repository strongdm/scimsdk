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

var MOCK_GROUPS_PAGE_SIZE = 2

// TODO: comparar testes com os testes de usuário
func TestGroupServiceCreate(t *testing.T) {
	t.Run("should create a group when passing valid data", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := newGroupService("token")
		opts := serviceCreateOptions{Body: getValidCreateGroup()}
		group, err := service.create(context.Background(), &opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when creating a group without token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := newGroupService("")
		opts := serviceCreateOptions{Body: getValidCreateGroup()}
		group, err := service.create(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should return a group when creating using context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := serviceCreateOptions{Body: getValidCreateGroup()}
		group, err := service.create(ctx, &opts)

		assert.NotNil(t, group)
		assert.Nil(t, ctx.Err())
		assert.Nil(t, err)
	})

	t.Run("should return a context error when context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := serviceCreateOptions{Body: getValidCreateGroup()}
		group, err := service.create(ctx, &opts)

		assert.Nil(t, group)
		assert.NotNil(t, ctx.Err())
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceListIterator(t *testing.T) {
	t.Run("should return a groups iterator when passing no pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		iterator := service.listIterator(context.Background(), &serviceListOptions{})

		assert.NotNil(t, iterator)
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
	})

	t.Run("should return a groups iterator when passing pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		iterator := service.listIterator(context.Background(), &serviceListOptions{PageSize: MOCK_GROUPS_PAGE_SIZE, Offset: 1})

		assert.NotNil(t, iterator)
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.True(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.NotNil(t, iterator.Value())
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.True(t, iterator.IsEmpty())
	})

	t.Run("should return a groups iterator when passing a page size greater than the offset and the groups count", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		iterator := service.listIterator(context.Background(), &serviceListOptions{PageSize: MOCK_GROUPS_PAGE_SIZE, Offset: 3})

		assert.NotNil(t, iterator)
		assert.False(t, iterator.Next())
		assert.Empty(t, iterator.Err())
		assert.Nil(t, iterator.Value())
		assert.True(t, iterator.IsEmpty())
	})
}

func TestGroupServiceList(t *testing.T) {
	t.Run("should return a list of groups when have no pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		groups, haveNextPage, err := service.list(context.Background(), &serviceListOptions{})

		assert.NotNil(t, groups)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, groups, 2)
	})

	t.Run("should return a list of groups when have pagination options", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		groups, haveNextPage, err := service.list(context.Background(), &serviceListOptions{PageSize: MOCK_GROUPS_PAGE_SIZE})

		assert.NotNil(t, groups)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
		assert.Len(t, groups, 2)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("")
		groups, haveNextPage, err := service.list(context.Background(), &serviceListOptions{})

		assert.Nil(t, groups)
		assert.False(t, haveNextPage)
		assert.NotNil(t, err)
	})

	t.Run("should return a list of groups when passing a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupPageResponse)
		service := newGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.list(ctx, &serviceListOptions{})

		assert.NotNil(t, groups)
		assert.False(t, haveNextPage)
		assert.Nil(t, err)
		assert.Nil(t, ctx.Err())
		assert.Len(t, groups, 2)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		service := newGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		groups, haveNextPage, err := service.list(ctx, &serviceListOptions{})

		assert.Nil(t, groups)
		assert.False(t, haveNextPage)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
		assert.Len(t, groups, 0)
	})
}

func TestGroupServiceFind(t *testing.T) {
	t.Run("should return a group when passing a valid group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := newGroupService("token")
		opts := serviceFindOptions{ID: "xxx"}
		group, err := service.find(context.Background(), &opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an invalid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		service := newGroupService("")
		opts := serviceFindOptions{ID: "xxx"}
		group, err := service.find(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an invalid group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupNotFound)
		service := newGroupService("token")
		opts := serviceFindOptions{ID: "yyy"}
		group, err := service.find(context.Background(), &opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := serviceFindOptions{ID: "xxx"}
		group, err := service.find(context.Background(), &opts)

		assert.Nil(t, ctx.Err())
		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := serviceFindOptions{ID: "xxx"}
		group, err := service.find(context.Background(), &opts)

		assert.NotNil(t, ctx.Err())
		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceReplace(t *testing.T) {
	t.Run("should replace a group when passing a valid group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		opts := &serviceReplaceOptions{"yyy", getValidReplaceGroup(), ""}
		service := newGroupService("token")
		group, err := service.replace(context.Background(), opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupNotFound)
		opts := &serviceReplaceOptions{"", getValidReplaceGroup(), ""}
		service := newGroupService("token")
		group, err := service.replace(context.Background(), opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupNotFound)
		opts := &serviceReplaceOptions{"yyy", getValidReplaceGroup(), ""}
		service := newGroupService("")
		group, err := service.replace(context.Background(), opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupResponse)
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := &serviceReplaceOptions{"yyy", getValidReplaceGroup(), ""}
		group, err := service.replace(ctx, opts)

		assert.NotNil(t, group)
		assert.Nil(t, err)
		assert.Nil(t, ctx.Err())
	})

	t.Run("should replace a group when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		service := newGroupService("token")
		opts := &serviceReplaceOptions{"yyy", getValidReplaceGroup(), ""}
		group, err := service.replace(ctx, opts)

		assert.Nil(t, group)
		assert.NotNil(t, err)
		assert.NotNil(t, ctx.Err())
		assert.Equal(t, "context deadline exceeded", ctx.Err().Error())
		assert.Equal(t, "context deadline exceeded", err.Error())
	})
}

func TestGroupServiceDelete(t *testing.T) {
	t.Run("should delete the user when passing a valid token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedGroup)
		service := newGroupService("token")
		ok, err := service.delete(context.Background(), &serviceDeleteOptions{ID: "xxx"})

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should return an error when passing an empty token", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedGroup)
		service := newGroupService("")
		ok, err := service.delete(context.Background(), &serviceDeleteOptions{ID: "xxx"})

		assert.False(t, ok)
		assert.NotNil(t, err)
	})

	t.Run("should return an error when passing an empty group id", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithGroupNotFound)
		service := newGroupService("token")
		ok, err := service.delete(context.Background(), &serviceDeleteOptions{ID: "xxx"})

		assert.False(t, ok)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
	})

	t.Run("should delete the user when using a context with timeout", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteDeletedGroup)
		service := newGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		ok, err := service.delete(ctx, &serviceDeleteOptions{ID: "xxx"})

		assert.True(t, ok)
		assert.Nil(t, err)
	})

	t.Run("should delete the user when the context timeout exceed", func(t *testing.T) {
		monkey.Patch(executeHTTPRequest, mockedApiExecuteWithExpiredTimeout)
		service := newGroupService("token")
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		ok, err := service.delete(ctx, &serviceDeleteOptions{ID: "xxx"})

		assert.False(t, ok)
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
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getGroupsPageResponseJSON())))
	emptyReader := ioutil.NopCloser(bytes.NewReader([]byte("{}")))
	response := &http.Response{Body: reader}
	startIndex := request.URL.Query().Get("startIndex")
	if startIndex > request.URL.Query().Get("count") && startIndex > "2" {
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

func getValidCreateGroup() *CreateGroupBody {
	return &CreateGroupBody{
		DisplayName: "xxx",
		Members: []*GroupMember{
			{
				Display: "xxx",
				Value:   "yyy",
			},
		},
	}
}

func getValidReplaceGroup() *ReplaceGroupBody {
	return &ReplaceGroupBody{
		DisplayName: "xxx",
		Members: []*GroupMember{
			{
				Display: "zzz",
				Value:   "www",
			},
		},
	}
}
