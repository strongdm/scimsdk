package sdmscim

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var MOCK_USERS_PAGE_SIZE = 1

type mockUsersService struct {
	mock.Mock
}

type TestIterators struct{}

func TestTestIterators(t *testing.T) {
	executeTests(t, reflect.TypeOf(TestIterators{}), nil, afterEach)
}

func afterEach() {
	MOCK_USERS_PAGE_SIZE = 1
}

func (ti *TestIterators) TestIterators(t *testing.T) {
	t.Run("should have one user in users iterator when paginating successfully", func(t *testing.T) {
		mUsersService := mockUsersService{}
		mUsersService.On("list", 0).Return(getUsers(1), true, nil)
		mUsersService.On("list", 1).Return(getUsers(2), false, nil)
		usersIterator := newUsersIterator(mUsersService.list, &ModuleListOptions{})

		assert.True(t, usersIterator.Next())
		assert.Equal(t, getUsers(1)[0], usersIterator.Value())
		assert.Equal(t, "", usersIterator.Err())
		assert.False(t, usersIterator.Next())
		assert.Equal(t, (*User)(nil), usersIterator.Value())
		assert.True(t, mUsersService.AssertExpectations(t))
	})

	t.Run("should return an offset error when passing a negative offset", func(t *testing.T) {
		mUsersService := mockUsersService{}
		mUsersService.On("listWithError", -1).Return(([]*User)(nil), false, errors.New("offset error"))
		usersIterator := newUsersIterator(mUsersService.listWithError, &ModuleListOptions{})

		assert.False(t, usersIterator.Next())
		assert.Nil(t, usersIterator.Value())
		assert.Equal(t, "offset error", usersIterator.Err())
		assert.True(t, mUsersService.AssertExpectations(t))
	})

	t.Run("should have only one page with one user and stop the pagination when the page size is greater than the users count", func(t *testing.T) {
		MOCK_USERS_PAGE_SIZE = 5
		mUsersService := mockUsersService{}
		mUsersService.On("list", 0).Return(getUsers(1), false, nil)
		usersIterator := newUsersIterator(mUsersService.list, &ModuleListOptions{})
		firstUser := getUsers(1)[0]

		assert.True(t, usersIterator.Next())
		assert.Equal(t, firstUser, usersIterator.Value())
		assert.Equal(t, "", usersIterator.Err())
		assert.False(t, usersIterator.Next())
		// The iterator is stopped because doesn't have more items and the index wasn't increased
		assert.Equal(t, firstUser, usersIterator.Value())
		assert.True(t, mUsersService.AssertExpectations(t))
	})
}

func (m mockUsersService) list(opts *ModuleListOptions) (users []*User, haveNextPage bool, err error) {
	m.Called(opts.Offset)
	offset := opts.Offset
	if offset == -1 {
		return nil, false, errors.New("offset error")
	}
	users = getUsers(offset + 1)
	haveNextPage = len(users) >= MOCK_USERS_PAGE_SIZE
	return users, haveNextPage, nil
}

func (m mockUsersService) listWithError(opts *ModuleListOptions) (users []*User, haveNextPage bool, err error) {
	m.Called(opts.Offset)
	return nil, false, errors.New("offset error")
}

func getUsers(offset int) []*User {
	if offset <= MOCK_USERS_PAGE_SIZE {
		return []*User{
			{},
		}
	}
	return []*User{}
}
