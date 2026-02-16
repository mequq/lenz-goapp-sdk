package main

import (
	"context"
	"log/slog"

	region "github.com/mequq/lenz-goapp-sdk/pkg/utils/regiondetector"
)

// func main() {
// 	apperr := httperror.New(400, "bad request", 1001, errors.New("test"))
// 	apperr.Join(errors.New("test1"))

// 	apperr.Join(errors.New("test2"))
// 	hdl := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	})

// 	logger := slog.New(hdl)

// 	// testErr := errors.New("test")
// 	// set debug mode
// 	httperror.DebugMode = true

// 	mid := httpmiddleware.NewGorilaMuxMiddleware(httpmiddleware.WithJwtSecret("test"), httpmiddleware.WithLogger(logger))

// 	// add middleware logger and recover
// 	http.Handle("/",
// 		mid.ContextMiddleware(
// 			mid.RecoverMiddleware(
// 				mid.AuthMiddleware(
// 					mid.LoggerMiddleware(
// 						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 							ctx := appcontext.SetContextFromHttpReq(r.Context(), r)
// 							logger.Info("err", "apperr", apperr, "ctx", appcontext.LogContext(ctx))
// 							httperror.ConvertError(apperr).HttpError(w)
// 						}),
// 					),
// 				),
// 			),
// 		),
// 	)

// 	http.ListenAndServe(":8080", nil)

// }

// func main() {
// 	ctx := context.Background()
// 	reg, err := region.NewRegionDetector(region.WithEndpoint("http://localhost:8088"))
// 	if err != nil {
// 		slog.Error("err", "err", err)
// 	}
// 	ip := "10.10.10.10"
// 	region, err := reg.GetRegion(ctx, ip)
// 	if err != nil {
// 		slog.Error("err", "err", err)
// 	}
// 	slog.Info("region", "region", region)

// 	slog.Info("region", "isiran", reg.IsIranIP(ctx, ip))
// }

func main() {
	// channelList()
	checkIP("192.168.18.10")
}

// func channelList() {
// 	req := husdk.NewChannelListRequest()
// 	resp := &husdk.ChannelListResponse{}

// 	err := husdk.NewEpg()
// }

func checkIP(ip string) {
	ctx := context.Background()
	reg, err := region.NewRegionDetector(region.WithEndpoint("http://localhost:8088"))
	if err != nil {
		slog.Error("err", "err", err)
		panic(err)
	}

	region, err := reg.GetRegion(ctx, ip)
	if err != nil {
		slog.Error("err", "err", err)
		panic(err)
	}
	slog.Info("region", "region", region)

	slog.Info("region", "isiran", reg.IsIranIP(ctx, ip))
}
