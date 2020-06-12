package repo

type FilterOption func(*FilterOptions)

type FilterOptions struct {
	limit  *uint32
	offset *uint32

	countryCode *string
}

func WithLimit(limit uint32) FilterOption {
	return func(options *FilterOptions) {
		options.limit = &limit
	}
}

func GetLimit(opts FilterOptions) *uint32 {
	return opts.limit
}

func WithOffset(offset uint32) FilterOption {
	return func(options *FilterOptions) {
		options.offset = &offset
	}
}

func GetOffset(opts FilterOptions) *uint32 {
	return opts.offset
}

func WithCountryCode(countryCode string) FilterOption {
	return func(options *FilterOptions) {
		options.countryCode = &countryCode
	}
}

func GetCountryCode(opts FilterOptions) *string {
	return opts.countryCode
}