package httpmiddleware

import (
	"net/http"

	"github.com/mequq/lenz-goapp-sdk/pkg/utils/appcontext"
)

// NewGorillaMuxServer creates a new HTTP server and set up all routes.
func (m *GorillaMuxMiddleware) ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = appcontext.SetContextFromHttpReq(ctx, r)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
