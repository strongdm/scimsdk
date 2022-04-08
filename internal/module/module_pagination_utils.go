package module

import (
	"github.com/strongdm/scimsdk/models"
)

type iteratorFetchFunc[T interface{}] func(opts *models.PaginationOptions) (data []*T, haveNextPage bool, err error)

type iteratorImpl[T interface{}] struct {
	buffer       []*T
	index        int
	haveNextPage bool
	fetchFn      iteratorFetchFunc[T]
	err          error
	opts         *models.PaginationOptions
}

func newIterator[T interface{}](fetchFn iteratorFetchFunc[T], opts *models.PaginationOptions) *iteratorImpl[T] {
	if opts == nil {
		opts = &models.PaginationOptions{
			Offset: 1,
		}
	}
	return &iteratorImpl[T]{
		haveNextPage: true,
		fetchFn:      fetchFn,
		opts:         opts,
	}
}

func (it *iteratorImpl[T]) Next() bool {
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

func (it *iteratorImpl[T]) Value() *T {
	if it.index > len(it.buffer)-1 {
		return nil
	}
	return it.buffer[it.index]
}

func (it *iteratorImpl[T]) Err() error {
	if it.err == nil {
		return nil
	}
	return it.err
}

func (it *iteratorImpl[T]) IsEmpty() bool {
	return it.buffer == nil || len(it.buffer) == 0
}
