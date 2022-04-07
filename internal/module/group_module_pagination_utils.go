package module

import "github.com/strongdm/scimsdk/models"

type listGroupsOperationFunc func(opts *models.PaginationOptions) ([]*models.Group, bool, error)

type GroupIterator interface {
	Next() bool
	Value() *models.Group
	Err() error
	IsEmpty() bool
}

type groupsIteratorImpl struct {
	buffer       []*models.Group
	index        int
	haveNextPage bool
	fetchFn      listGroupsOperationFunc
	err          error
	opts         *models.PaginationOptions
}

func newGroupsIterator(fetchFn listGroupsOperationFunc, opts *models.PaginationOptions) *groupsIteratorImpl {
	if opts == nil {
		opts = &models.PaginationOptions{
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

func (it *groupsIteratorImpl) Value() *models.Group {
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
