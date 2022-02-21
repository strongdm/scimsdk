package sdmscim

type FetchUsersOperation func(offset int) (users []*User, haveNextPage bool, err error)
type FetchGroupsOperation func(offset int) (groups []*Group, haveNextPage bool, err error)

type UsersIterator struct {
	buffer       []*User
	index        int
	haveNextPage bool
	fetchFn      FetchUsersOperation
	err          error
	offset       int
}

type GroupsIterator struct {
	buffer       []*Group
	index        int
	haveNextPage bool
	fetchFn      FetchGroupsOperation
	err          error
	offset       int
}

func newUsersIterator(fetchFn FetchUsersOperation) *UsersIterator {
	return &UsersIterator{
		haveNextPage: true,
		fetchFn:      fetchFn,
	}
}

func newGroupsIterator(fetchFn FetchGroupsOperation) *GroupsIterator {
	return &GroupsIterator{
		haveNextPage: true,
		fetchFn:      fetchFn,
	}
}

// ----------------------------------
// UsersIterator
// ----------------------------------
func (it *UsersIterator) Next() bool {
	if it.index < len(it.buffer)-1 {
		it.index++
		return true
	}
	if !it.haveNextPage {
		return false
	}
	it.offset = len(it.buffer) + it.offset
	it.index = 0
	it.buffer, it.haveNextPage, it.err = it.fetchFn(it.offset)
	return len(it.buffer) > 0
}

func (it *UsersIterator) Value() *User {
	if it.index > len(it.buffer)-1 {
		return nil
	}
	return it.buffer[it.index]
}

func (it *UsersIterator) Err() string {
	if it.err == nil {
		return ""
	}
	return it.err.Error()
}

// ----------------------------------

// ----------------------------------
// GroupsIterator
// ----------------------------------
func (it *GroupsIterator) Next() bool {
	if it.index < len(it.buffer)-1 {
		it.index++
		return true
	}
	if !it.haveNextPage {
		return false
	}
	it.offset = len(it.buffer) + it.offset
	it.index = 0
	it.buffer, it.haveNextPage, it.err = it.fetchFn(it.offset)
	return len(it.buffer) > 0
}

func (it *GroupsIterator) Value() Group {
	return *it.buffer[it.index]
}

func (it *GroupsIterator) Err() string {
	if it.err == nil {
		return ""
	}
	return it.err.Error()
}
