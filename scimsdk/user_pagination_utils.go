package scimsdk

type listUsersOperationFunc func(opts *PaginationOptions) (users []*User, haveNextPage bool, err error)

type usersIteratorImpl struct {
	buffer       []*User
	index        int
	haveNextPage bool
	fetchFn      listUsersOperationFunc
	err          error
	opts         *PaginationOptions
}

func newUsersIterator(fetchFn listUsersOperationFunc, opts *PaginationOptions) *usersIteratorImpl {
	if opts == nil {
		opts = &PaginationOptions{
			Offset: 1,
		}
	}
	return &usersIteratorImpl{
		haveNextPage: true,
		fetchFn:      fetchFn,
		opts:         opts,
	}
}

func (it *usersIteratorImpl) Next() bool {
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

func (it *usersIteratorImpl) Value() *User {
	if it.index > len(it.buffer)-1 {
		return nil
	}
	return it.buffer[it.index]
}

func (it *usersIteratorImpl) Err() error {
	if it.err == nil {
		return nil
	}
	return it.err
}

func (it *usersIteratorImpl) IsEmpty() bool {
	return it.buffer == nil || len(it.buffer) == 0
}
