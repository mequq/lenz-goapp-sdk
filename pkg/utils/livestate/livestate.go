package livestate

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type LiveState struct {
	endpoint *url.URL
	logger   *slog.Logger
	client   *http.Client
}

type ChannelState struct {
	ChannelVOID  string `json:"channel_id" validate:"required,number,max=25"`
	Name         string `json:"channel_name"`
	Desc         string `json:"channel_desc"`
	DefaultAvail bool   `json:"default_avail"`
	Available    bool   `json:"available"`
}

// is channel copyrited
func (r *ChannelState) IsChannelCopyrited() bool {
	return !r.Available
}

var (
	ErrorNotValidEndpoint = errors.New("endpoint is not valid")
)

type LiveStateOption func(*LiveState) error

func NewLiveState(options ...LiveStateOption) (*LiveState, error) {

	endpoint, err := url.Parse("http://localhost:8088")
	if err != nil {
		return nil, ErrorNotValidEndpoint
	}

	r := &LiveState{
		endpoint: endpoint,
		logger:   slog.Default().With("module", "live-state"),
		client:   http.DefaultClient,
	}
	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func WithEndpoint(endpoint string) LiveStateOption {
	return func(r *LiveState) error {
		endpoint, err := url.Parse(endpoint)
		if err != nil {
			return ErrorNotValidEndpoint
		}
		r.endpoint = endpoint
		return nil
	}

}

func WithLogger(logger *slog.Logger) LiveStateOption {
	return func(r *LiveState) error {
		r.logger = logger
		return nil
	}
}

func WithClient(client *http.Client) LiveStateOption {
	return func(r *LiveState) error {
		r.client = client
		return nil
	}
}

func (r *LiveState) GetChannelStates(channelId string) (*ChannelState, error) {
	response, err := r.client.Get(r.endpoint.String() + "/api/v1/admin/live-state/channel/" + channelId)
	if err != nil {
		r.logger.Debug("err", "err", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		r.logger.Debug("err", "err", err)
		return nil, err
	}
	r.logger.Debug("body", "body", string(body))

	var channelState *ChannelState
	err = json.Unmarshal(body, &channelState)
	if err != nil {
		r.logger.Debug("err", "err", err)
		return nil, err
	}

	r.logger.Debug("channel", "state", channelState)
	return channelState, nil
}
