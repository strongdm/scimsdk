package models

type PaginationOptions struct {
	PageSize int
	Offset   int
	Filter   string
}

type Iterator[T interface{}] interface {
	Next() bool
	Value() *T
	Err() error
	IsEmpty() bool
}
