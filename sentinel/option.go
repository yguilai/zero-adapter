package sentinel

import "net/http"

type (
	Option  func(*options)
	options struct {
		resourceExtract func(*http.Request) string
		blockFallback   func(http.ResponseWriter, *http.Request)
	}
)

// WithResourceExtractor sets the resource extractor of the web requests.
func WithResourceExtractor(fn func(*http.Request) string) Option {
	return func(opts *options) {
		opts.resourceExtract = fn
	}
}

// WithBlockFallback sets the fallback handler when requests are blocked.
func WithBlockFallback(fn func(http.ResponseWriter, *http.Request)) Option {
	return func(opts *options) {
		opts.blockFallback = fn
	}
}
