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

func getDefaultPaginationOptions() *PaginationOptions {
	return &PaginationOptions{
		PageSize: defaultAPIPageSize,
		Offset:   1,
	}
}

func newServiceListOptions(opts *PaginationOptions, url string) *serviceListOptions {
	defaultPaginationOptions := getDefaultPaginationOptions()
	if opts == nil {
		opts = defaultPaginationOptions
	} else if opts.Offset < 0 {
		log.Fatal("The pagination offset must be positive")
	} else if opts.PageSize < 0 {
		log.Fatal("The pagination page size must be positive")
	} else if opts.PageSize == 0 {
		opts.PageSize = defaultPaginationOptions.PageSize
	} else if opts.Offset == 0 {
		opts.Offset = defaultPaginationOptions.Offset
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

func newServiceUpdateOptions(id string, active bool, url string) *serviceUpdateOptions {
	if id == "" {
		log.Fatal("You must pass the user id")
	}
	baseURL := defaultAPIURL
	if url != "" {
		baseURL = url
	}
	return &serviceUpdateOptions{
		ID:         id,
		Body:       convertPorcelainToUpdateUserRequest(active),
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
