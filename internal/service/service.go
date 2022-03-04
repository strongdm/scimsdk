package service

import "github.com/strongdm/scimsdk/internal/api"

type CreateOptions struct {
	Body       interface{}
	BaseAPIURL string
}

type ListOptions struct {
	PageSize   int
	Offset     int
	Filter     string
	BaseAPIURL string
}

type FindOptions struct {
	ID         string
	BaseAPIURL string
}

type ReplaceOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type UpdateOptions struct {
	ID         string
	Body       interface{}
	BaseAPIURL string
}

type DeleteOptions struct {
	ID         string
	BaseAPIURL string
}

func newAPICreateOptions(opts *CreateOptions) *api.CreateOptions {
	return api.NewCreateOptions(opts.Body, opts.BaseAPIURL)
}

func newAPIListOptions(opts *ListOptions) *api.ListOptions {
	return api.NewListOptions(opts.PageSize, opts.Offset, opts.Filter, opts.BaseAPIURL)
}

func newAPIFindOptions(opts *FindOptions) *api.FindOptions {
	return api.NewFindOptions(opts.ID, opts.BaseAPIURL)
}

func newAPIReplaceOptions(opts *ReplaceOptions) *api.ReplaceOptions {
	return api.NewReplaceOptions(opts.ID, opts.Body, opts.BaseAPIURL)
}

func newAPIUpdateOptions(opts *UpdateOptions) *api.UpdateOptions {
	return api.NewUpdateOptions(opts.ID, opts.Body, opts.BaseAPIURL)
}

func newAPIDeleteOptions(opts *DeleteOptions) *api.DeleteOptions {
	return api.NewDeleteOptions(opts.ID, opts.BaseAPIURL)
}
