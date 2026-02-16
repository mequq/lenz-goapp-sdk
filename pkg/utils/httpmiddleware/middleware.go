package httpmiddleware

import (
	"log/slog"
	"net/url"

	region "github.com/mequq/lenz-goapp-sdk/pkg/utils/regiondetector"
)

type GorillaMuxMiddleware struct {
	logger         *slog.Logger
	level          slog.Level
	jwtSecret      []byte
	loginUri       *url.URL
	regionDetector region.RegionDetector
}

type MiddlewareOpt func(*GorillaMuxMiddleware) error

func NewGorilaMuxMiddleware(opt ...MiddlewareOpt) *GorillaMuxMiddleware {
	m := &GorillaMuxMiddleware{
		logger:         slog.Default(),
		level:          slog.LevelInfo,
		jwtSecret:      []byte(""),
		loginUri:       nil,
		regionDetector: nil,
	}
	for _, o := range opt {
		o(m)
	}
	return m
}

func WithLogger(logger *slog.Logger) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) error {
		m.logger = logger
		return nil
	}
}

func WithLevel(level slog.Level) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) error {
		m.level = level
		return nil
	}
}

func WithJwtSecret(jwtSecret string) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) error {
		m.jwtSecret = []byte(jwtSecret)
		return nil
	}
}

func WithRegionDetector(rd region.RegionDetector) MiddlewareOpt {
	return func(m *GorillaMuxMiddleware) error {
		m.regionDetector = rd
		return nil
	}
}
