package sdmscim

// ModuleListOptions refer to APIListOptions
type ModuleListOptions struct {
	// PageSize defines the count of the users by page
	PageSize int
	// Offset defines the page offset referencing to the page - relative to the PageSize
	Offset int
	// Filter defines the query filter used in strongDM
	Filter string

	baseAPIURL string
}

func moduleListOptionsToAPIOptions(opts *ModuleListOptions) *apiListOptions {
	return &apiListOptions{
		PageSize:   opts.PageSize,
		Offset:     opts.Offset,
		Filter:     opts.Filter,
		BaseAPIURL: opts.baseAPIURL,
	}
}

func (opts *ModuleListOptions) SetBaseAPIURL(url string) {
	opts.baseAPIURL = url
}
