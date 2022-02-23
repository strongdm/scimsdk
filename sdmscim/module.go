package sdmscim

import "log"

func newServiceCreateOptions(user *CreateUser, url string) *serviceCreateOptions {
	baseURL := defaultAPIURL
	if url != "" {
		baseURL = url
	}
	return &serviceCreateOptions{
		Body:       convertPorcelainToCreateUserRequest(user),
		BaseAPIURL: baseURL,
	}
}

// PaginationOptions refer to serviceListOptions
type PaginationOptions struct {
	// PageSize defines the count of the users by page
	PageSize int
	// Offset defines the page offset referencing to the page - relative to the PageSize
	Offset int
	// Filter defines the query filter used in strongDM
	Filter string
}

func newServiceListOptions(opts *PaginationOptions, url string) *serviceListOptions {
	if opts == nil {
		opts = &PaginationOptions{}
	}
	baseURL := defaultAPIURL
	if url == "" {
		baseURL = url
	}
	return &serviceListOptions{
		PageSize:   opts.PageSize,
		Offset:     opts.Offset,
		Filter:     opts.Filter,
		BaseAPIURL: baseURL,
	}
}

func newServiceFindOptions(id string, url string) *serviceFindOptions {
	if id == "" {
		log.Fatal("You must pass the user id")
	}
	baseURL := defaultAPIURL
	if url != "" {
		baseURL = url
	}
	return &serviceFindOptions{
		ID:         id,
		BaseAPIURL: baseURL,
	}
}

func newServiceDeleteOptions(id string, url string) *serviceDeleteOptions {
	if id == "" {
		log.Fatal("You must pass the user id")
	}
	baseURL := defaultAPIURL
	if url != "" {
		baseURL = url
	}
	return &serviceDeleteOptions{
		ID:         id,
		BaseAPIURL: baseURL,
	}
}

func newServiceReplaceOptions(id string, user *ReplaceUser, url string) *serviceReplaceOptions {
	if id == "" {
		log.Fatal("You must pass the user id")
	}
	baseURL := defaultAPIURL
	if url != "" {
		baseURL = url
	}
	return &serviceReplaceOptions{
		ID:         id,
		Body:       convertPorcelainToReplaceUserRequest(user),
		BaseAPIURL: baseURL,
	}
}
