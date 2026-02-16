package httpmiddleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/mequq/lenz-goapp-sdk/pkg/utils/appcontext"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func (m *GorillaMuxMiddleware) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recored time duration
		startTime := time.Now()

		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		m.logger.Log(r.Context(), m.level, "request fulfilled",
			slog.Group(
				"request-info",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
			),
			"ctx", appcontext.LogContext(r.Context()),
			"status", recorder.Status,
			"duration", time.Since(startTime).String(),
		)
	})
}
