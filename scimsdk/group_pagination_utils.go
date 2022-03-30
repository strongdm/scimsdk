package scimsdk

type listGroupsOperationFunc func(opts *PaginationOptions) ([]*Group, bool, error)

type groupsIteratorImpl struct {
	buffer       []*Group
	index        int
	haveNextPage bool
	fetchFn      listGroupsOperationFunc
	err          error
	opts         *PaginationOptions
}

func newGroupsIterator(fetchFn listGroupsOperationFunc, opts *PaginationOptions) *groupsIteratorImpl {
	if opts == nil {
		opts = &PaginationOptions{
			Offset: 1,
		}
	}
	return &groupsIteratorImpl{
		fetchFn:      fetchFn,
		haveNextPage: true,
		opts:         opts,
	}
}

func (it *groupsIteratorImpl) Next() bool {
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

func (it *groupsIteratorImpl) Value() *Group {
	if it.index > len(it.buffer)-1 {
		return nil
	}
	return it.buffer[it.index]
}

func (it *groupsIteratorImpl) Err() error {
	if it.err == nil {
		return nil
	}
	return it.err
}

func (it *groupsIteratorImpl) IsEmpty() bool {
	return it.buffer == nil || len(it.buffer) == 0
}
