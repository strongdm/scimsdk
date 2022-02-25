package sdmscim

type ListUsersOperationFunc func(opts *serviceListOptions) (users []*User, haveNextPage bool, err error)

type UsersIterator struct {
	buffer       []*User
	index        int
	haveNextPage bool
	fetchFn      ListUsersOperationFunc
	err          error
	opts         *serviceListOptions
}

func newUsersIterator(fetchFn ListUsersOperationFunc, opts *serviceListOptions) *UsersIterator {
	return &UsersIterator{
		haveNextPage: true,
		fetchFn:      fetchFn,
		opts:         opts,
	}
}

func (it *UsersIterator) Next() bool {
	if it.index < len(it.buffer)-1 {
		it.index++
		return true
	}
	if !it.haveNextPage {
		return false
	}
	it.opts.Offset = len(it.buffer) + it.opts.Offset
	it.index = 0
	it.buffer, it.haveNextPage, it.err = it.fetchFn(it.opts)
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

func (it *UsersIterator) IsEmpty() bool {
	return it.buffer == nil || len(it.buffer) == 0
}
