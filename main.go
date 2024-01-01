package main

import (
	"errors"
	"log/slog"
	"net/http"
	"os"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"
	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/httperror"
	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/httpmiddleware"
)

func main() {
	apperr := httperror.New(400, "bad request", 1001, errors.New("test"))
	apperr.Join(errors.New("test1"))

	apperr.Join(errors.New("test2"))
	hdl := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logger := slog.New(hdl)

	// testErr := errors.New("test")
	// set debug mode
	httperror.DebugMode = true

	mid := httpmiddleware.NewGorilaMuxMiddleware(httpmiddleware.WithJwtSecret("test"), httpmiddleware.WithLogger(logger))

	// add middleware logger and recover
	http.Handle("/",
		mid.ContextMiddleware(
			mid.RecoverMiddleware(
				mid.AuthMiddleware(
					mid.LoggerMiddleware(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							ctx := appcontext.SetContextFromHttpReq(r.Context(), r)
							logger.Info("err", "apperr", apperr, "ctx", appcontext.LogContext(ctx))
							httperror.ConvertError(apperr).HttpError(w)
						}),
					),
				),
			),
		),
	)

	http.ListenAndServe(":8080", nil)

}
