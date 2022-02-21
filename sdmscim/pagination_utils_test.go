package sdmscim

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	MOCK_USERS_PAGE_LIMIT = 1
)

type mockUsersApplication struct {
	mock.Mock
}

func TestIterators(t *testing.T) {
	t.Run("should have one user in users iterator", func(t *testing.T) {
		mUsersApplication := mockUsersApplication{}
		mUsersApplication.On("fetchUsers", 0).Return(getUsers(1), true, nil)
		mUsersApplication.On("fetchUsers", 1).Return(getUsers(2), false, nil)
		usersIterator := newUsersIterator(mUsersApplication.list)

		assert.True(t, usersIterator.Next())
		assert.Equal(t, getUsers(1)[0], usersIterator.Value())
		assert.Equal(t, "", usersIterator.Err())
		assert.False(t, usersIterator.Next())
		assert.Equal(t, (*User)(nil), usersIterator.Value())
		assert.True(t, mUsersApplication.AssertExpectations(t))
	})

	t.Run("should return an offset error", func(t *testing.T) {
		mUsersApplication := mockUsersApplication{}
		mUsersApplication.On("fetchUsers", -1).Return(([]*User)(nil), false, errors.New("offset error"))
		usersIterator := newUsersIterator(mUsersApplication.listWithError)

		assert.False(t, usersIterator.Next())
		assert.Nil(t, usersIterator.Value())
		assert.Equal(t, "offset error", usersIterator.Err())
		assert.True(t, mUsersApplication.AssertExpectations(t))
	})

	t.Run("should have only one page with one user and stop the pagination", func(t *testing.T) {
		MOCK_USERS_PAGE_LIMIT = 5
		mFetchUsers := mockUsersApplication{}
		mFetchUsers.On("fetchUsers", 0).Return(getUsers(1), false, nil)
		usersIterator := newUsersIterator(mFetchUsers.list)
		firstUser := getUsers(1)[0]

		assert.True(t, usersIterator.Next())
		assert.Equal(t, firstUser, usersIterator.Value())
		assert.Equal(t, "", usersIterator.Err())
		assert.False(t, usersIterator.Next())
		// The iterator in stopped because doesn't have more items and the index wasn't increased
		assert.Equal(t, firstUser, usersIterator.Value())
		assert.True(t, mFetchUsers.AssertExpectations(t))
	})
}

func (m mockUsersApplication) list(offset int) (users []*User, haveNextPage bool, err error) {
	m.Called(offset)
	if offset == -1 {
		return nil, false, errors.New("offset error")
	}
	users = getUsers(offset + 1)
	haveNextPage = len(users) >= MOCK_USERS_PAGE_LIMIT
	return users, haveNextPage, nil
}

func (m mockUsersApplication) listWithError(offset int) (users []*User, haveNextPage bool, err error) {
	m.Called(offset)
	return nil, false, errors.New("offset error")
}

func getUsers(offset int) []*User {
	if offset <= MOCK_USERS_PAGE_LIMIT {
		return []*User{
			{},
		}
	}
	return []*User{}
}
