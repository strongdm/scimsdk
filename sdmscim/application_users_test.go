package sdmscim

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sdmscim/sdmscim/api"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestUsersApplicationList(t *testing.T) {
	defer monkey.UnpatchAll()

	t.Run("should return a list of users", func(t *testing.T) {
		userApplication := newUserApplication("")
		monkey.Patch(api.BaseList, mockedBaseListWithUserResponse)
		users, haveNextPage, err := userApplication.list(0)
		assert.NotNil(t, users)
		assert.True(t, haveNextPage)
		assert.Nil(t, err)
	})

	t.Run("should return a bad request error", func(t *testing.T) {
		userApplication := newUserApplication("")
		monkey.Patch(api.BaseList, mockedBaseListWithBadRequestError)
		users, haveNextPage, err := userApplication.list(0)
		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, "Bad request", err.Error())
	})

	t.Run("should return a permission denied message", func(t *testing.T) {
		userApplication := newUserApplication("")
		monkey.Patch(api.BaseList, mockedBaseListWithPermissionDenied)
		users, haveNextPage, err := userApplication.list(0)
		assert.Nil(t, users)
		assert.False(t, haveNextPage)
		assert.Equal(t, err.Error(), "invalid character 'p' looking for beginning of value")
	})
}

func mockedBaseListWithUserResponse(token string, pathname string, offset int) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte(getUsersSCIMResponseJSON())))
	return &http.Response{
		Body: reader,
	}, nil
}

func mockedBaseListWithBadRequestError(token string, pathname string, offset int) (*http.Response, error) {
	return nil, errors.New("Bad request")
}

func mockedBaseListWithPermissionDenied(token string, pathname string, offset int) (*http.Response, error) {
	reader := ioutil.NopCloser(bytes.NewReader([]byte("permission denied: access denied")))
	return &http.Response{
		Body: reader,
	}, nil
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
