package httpmiddleware

import (
	"errors"
	"fmt"

	"net/http"
	"runtime/debug"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"
	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/httperror"
)

func (m *GorillaMuxMiddleware) RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Handle the panic here, log the error, or send an appropriate response to the client.
				m.logger.Error("panic recovered", "err", err, "ctx", appcontext.LogContext(r.Context()), "trace", debug.Stack())
				apperror := httperror.New(http.StatusInternalServerError, "Internal Server Error", 1001, errors.New("internal server error"))
				switch v := err.(type) {
				case error:
					apperror = apperror.Join(v)
				case string:
					apperror = apperror.Join(errors.New(v))
				}

				fmt.Println(string(debug.Stack()))
				httperror.ConvertError(apperror).HttpError(w)

			}
		}()
		next.ServeHTTP(w, r)
	})
}
