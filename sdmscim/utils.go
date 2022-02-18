package sdmscim

type FetchUsersOperation func(offset int) (users []*User, haveNextPage bool, err error)
type FetchGroupsOperation func(offset int) (groups []*Group, haveNextPage bool, err error)

type usersIterator struct {
	buffer       []*User
	index        int
	haveNextPage bool
	fetchFn      FetchUsersOperation
	err          error
	offset       int
}

type groupsIterator struct {
	buffer       []*Group
	index        int
	haveNextPage bool
	fetchFn      FetchGroupsOperation
	err          error
	offset       int
}

func newUsersIterator(fetchFn FetchUsersOperation) *usersIterator {
	return &usersIterator{
		haveNextPage: true,
		fetchFn:      fetchFn,
	}
}

func newGroupsIterator(fetchFn FetchGroupsOperation) *groupsIterator {
	return &groupsIterator{
		haveNextPage: true,
		fetchFn:      fetchFn,
	}
}

// ----------------------------------
// UsersIterator
// ----------------------------------
func (it *usersIterator) Next() bool {
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

func (it *usersIterator) Value() User {
	// TODO: Need to pay attention to indexOutBounds
	return *it.buffer[it.index]
}

func (it *usersIterator) Err() string {
	if it.err == nil {
		return ""
	}
	return it.err.Error()
}

// ----------------------------------

// ----------------------------------
// GroupsIterator
// ----------------------------------
func (it *groupsIterator) Next() bool {
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

func (it *groupsIterator) Value() Group {
	return *it.buffer[it.index]
}

func (it *groupsIterator) Err() string {
	if it.err == nil {
		return ""
	}
	return it.err.Error()
}
