package region

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"
)

type Region struct {
	Country  string `json:"country,omitempty"`
	Provider string `json:"provider,omitempty"`
}

// func IS IRAN IP
func (r *Region) IsIran() bool {
	return r.Country == "IR"
}

// func IS MTNI Provider
func (r *Region) IsMTNI() bool {
	return r.Provider == "MTNI"
}

// RegionDetector Interface
type RegionDetector interface {
	GetRegion(ctx context.Context, ipAddrss string) (*Region, error)
	IsIranIP(ctx context.Context, ipAddrss string) bool
	IsMTNIProvider(ctx context.Context, ipAddrss string) bool
}

type region struct {
	endpoint string
	logger   *slog.Logger
	client   *http.Client
}

type Option func(*region) error

const (
	Address = "/api/v3/region/ip/%s"
)

var (
	ErrorNotValidIP         = errors.New("ip address is not valid")
	ErrorRegionStatusFailed = errors.New("get region failed")
)

// new RegionDetector with options
func NewRegionDetector(options ...Option) (RegionDetector, error) {
	r := &region{
		endpoint: "http://localhost:8088",
		logger:   slog.Default().With("module", "region-detector"),
		client:   http.DefaultClient,
	}
	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

// WithEndpoint set endpoint for region detector
func WithEndpoint(endpoint string) Option {
	return func(r *region) error {
		r.endpoint = endpoint
		return nil
	}

}

// WithLogger set logger for region detector
func WithLogger(logger *slog.Logger) Option {
	return func(r *region) error {
		r.logger = logger
		return nil
	}
}

// WithClient Set client for region detector
func WithClient(client *http.Client) Option {
	return func(r *region) error {
		r.client = client
		return nil
	}
}

// GetRegion get region by ip
func (r *region) GetRegion(ctx context.Context, ipAddrss string) (*Region, error) {
	ip := net.ParseIP(ipAddrss)
	if ip == nil {
		return nil, ErrorNotValidIP
	}
	url := fmt.Sprintf(r.endpoint+Address, ip.String())
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		r.logger.Warn(
			"failed to get success status code",
			"status-code", resp.StatusCode,
			"request", request,
			"ctx", appcontext.LogContext(ctx),
		)
		return nil, ErrorRegionStatusFailed
	}
	ret := new(Region)
	if err := json.NewDecoder(resp.Body).Decode(ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// Is IRAN IP
func (r *region) IsIranIP(ctx context.Context, ipAddrss string) bool {
	reg, err := r.GetRegion(ctx, ipAddrss)
	if err != nil {
		return false
	}
	return reg.IsIran()
}

// Is MTNI Provider
func (r *region) IsMTNIProvider(ctx context.Context, ipAddrss string) bool {
	reg, err := r.GetRegion(ctx, ipAddrss)
	if err != nil {
		return false
	}
	return reg.IsMTNI()
}
