package sentinel

import (
	sentinelApi "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

// SentinelMiddleware returns a sentinelApi middleware for go-sentinelApi
// default resource name is {method}:{uri}, such as "GET:/api/users"
// default block fallback is returning http.StatusTooManyRequests
func SentinelMiddleware(opts ...Option) rest.Middleware {
	options := &options{}
	for _, optFn := range opts {
		optFn(options)
	}
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			resourceName := r.Method + ":" + r.RequestURI

			if options.resourceExtract != nil {
				resourceName = options.resourceExtract(r)
			}

			entry, err := sentinelApi.Entry(
				resourceName,
				sentinelApi.WithResourceType(base.ResTypeWeb),
				sentinelApi.WithTrafficType(base.Inbound),
			)
			if err != nil {
				if options.blockFallback != nil {
					options.blockFallback(w, r)
				} else {
					w.WriteHeader(http.StatusTooManyRequests)
				}
				return
			}
			defer entry.Exit()
			next(w, r)
		}
	}
}
