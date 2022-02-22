package sdmscim

type ListUsersOperation func(opts *ModuleListOptions) (users []*User, haveNextPage bool, err error)

type UsersIterator struct {
	buffer       []*User
	index        int
	haveNextPage bool
	fetchFn      ListUsersOperation
	err          error
	opts         *ModuleListOptions
}

func newUsersIterator(fetchFn ListUsersOperation, opts *ModuleListOptions) *UsersIterator {
	if opts.Offset == 0 {
		opts.Offset = 1
	}
	if opts.PageSize == 0 {
		opts.PageSize = defaultAPIPageSize
	}
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
