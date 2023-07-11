package option

// Options 选项结构体
type Options struct {
	SortField string
	SortBy    string
	isPage    bool
	PageSize  int64
	PageNum   int64
}

// OptionFunc 选项函数
type OptionFunc func(opts *Options)

// InitOptions 初始化选项
func InitOptions(opts ...OptionFunc) *Options {
	options := new(Options)
	for _, opt := range opts {
		opt(options)
	}
	return options
}

// WithSortOption 排序选项
func WithSortOption(field, sortBy string) OptionFunc {
	return func(opts *Options) {
		opts.SortField = field
		opts.SortBy = sortBy
	}
}

// WithPageOption 分页选项
func WithPageOption(pageSize, pageNum int64) OptionFunc {
	return func(opts *Options) {
		opts.isPage = true
		opts.PageSize = pageSize
		opts.PageNum = pageNum
	}
}
