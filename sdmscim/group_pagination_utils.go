package sdmscim

type ListGroupsOperationFunc func(opts *serviceListOptions) (groups []*Group, haveNextPage bool, err error)

type GroupsIterator struct {
	buffer       []*Group
	index        int
	haveNextPage bool
	fetchFn      ListGroupsOperationFunc
	err          error
	opts         *serviceListOptions
}

func newGroupsIterator(fetchFn ListGroupsOperationFunc, opts *serviceListOptions) *GroupsIterator {
	return &GroupsIterator{
		fetchFn:      fetchFn,
		haveNextPage: true,
		opts:         opts,
	}
}

func (it *GroupsIterator) Next() bool {
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

func (it *GroupsIterator) Value() *Group {
	if it.index > len(it.buffer)-1 {
		return nil
	}
	return it.buffer[it.index]
}

func (it *GroupsIterator) Err() string {
	if it.err == nil {
		return ""
	}
	return it.err.Error()
}

func (it *GroupsIterator) IsEmpty() bool {
	return it.buffer == nil || len(it.buffer) == 0
}
