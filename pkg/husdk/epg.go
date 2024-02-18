package husdk

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"

	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	ErrorClientIPAddressNotValid = errors.New("client ip is not valid")
	ErrorEPGNil                  = errors.New("epg is not defined")
	ErrorBadRequestBody          = errors.New("bad request body")
	ErrorFaildToGetResponse      = errors.New("failed to get response")
	ErrorFailedToReadResponse    = errors.New("failed to read response")
	ErrorHURequestFailed         = errors.New("hu request failed")
	ErrorUserNeedToRelogin       = errors.New("user need to login")
)

type EPG struct {
	logger          *slog.Logger
	logLevel        slog.Level
	epgAddress      *url.URL
	Backend         string
	ClientIPAddress net.IP
	SessionID       string
	MSISDN          string
	httpclient      http.Client
}

type EPGOptons func(*EPG) error

func NewEpg(backend string, clientIP string, session string, msisdn string, opt ...EPGOptons) (epg *EPG, err error) {

	epgaddress, _ := url.Parse("http://10.230.37.171:33201")
	cip := net.ParseIP(clientIP)
	if cip == nil {
		return nil, ErrorClientIPAddressNotValid
	}
	epg = &EPG{
		logger:          slog.Default(),
		logLevel:        slog.LevelDebug,
		epgAddress:      epgaddress,
		Backend:         backend,
		ClientIPAddress: cip,
		SessionID:       session,
		MSISDN:          msisdn,
		httpclient:      http.Client{},
	}

	for _, v := range opt {
		if err := v(epg); err != nil {
			return nil, err
		}
	}
	return epg, nil

}

func (epg *EPG) Execute(ctx context.Context, req HURequestInterface, resp any) error {

	if epg == nil {
		return ErrorEPGNil
	}

	logger := epg.logger.With("method", "Execute", "ctx", appcontext.LogContext(ctx))
	logger.Debug("execute request", "url", req.GetPath(), "body", req.GetRequestBody())
	if epg.epgAddress == nil {
		logger.Warn("epg address is not defined",
			ctx, appcontext.LogContext(ctx),
		)
		return errors.New("epg address is not defined")
	}

	requrl := *epg.epgAddress
	requrl.Path = req.GetPath()
	q := req.GetQueryParameter()
	requrl.RawQuery = q.Encode()

	body, err := req.GetReuestData()
	if err != nil {
		logger.Warn("failed to get request data",
			"err", err,
			"ctx", appcontext.LogContext(ctx),
		)
		return err
	}

	request, err := http.NewRequest(req.GetMethod(), requrl.String(), body)
	if err != nil {
		epg.logger.Warn("create request error", "err", err)
		return errors.Join(err, ErrorBadRequestBody)
	}

	// epg.logger.Warn("execute request", "url", requrl.String(), "ctx", utils.LogContext(ctx), "body", req.GetRequestBody())

	request.Header.Set("Cookie", "JSESSIONID="+epg.SessionID)
	request.Header.Set("User-Agent", "Tiara-Middleware")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Backend", epg.Backend)
	request.Header.Set("X-Real-Ip", epg.ClientIPAddress.String())
	request.Header.Set("MSISDN", epg.MSISDN)

	// add header
	client := epg.httpclient

	startTime := time.Now()
	huresp, err := client.Do(request)
	if err != nil {
		epg.logger.Warn("failed to execute request to hu",
			"err", err,
			"ctx", appcontext.LogContext(ctx),
			"request", request.URL,
			"body", req.GetRequestBody(),
			"header", request.Header,
		)
		return errors.Join(err, ErrorFaildToGetResponse)
	}
	defer huresp.Body.Close()

	b, err := io.ReadAll(huresp.Body)
	if err != nil {
		epg.logger.Warn("failed to read response from hu",
			"err", err,
			"code", huresp.StatusCode,
			"body", string(b),
			"ctx", appcontext.LogContext(ctx),
		)
		return errors.Join(err, ErrorFailedToReadResponse)
	}

	huErr := &GeneralError{}
	err = json.Unmarshal(b, huErr)
	if err != nil {
		epg.logger.Warn("failed to unmarshal response from hu",
			"err", err,
			"code", huresp.StatusCode,
			"body", string(b),
			"request", request.URL,
			"request_body", req.GetRequestBody(),
		)
		return errors.Join(err, ErrorFailedToReadResponse)
	}

	if huErr.RetCode == "" || huErr.RetCode == "0" {
		err = json.Unmarshal(b, resp)
		if err != nil {
			epg.logger.Warn("failed to unmarshal response from hu",
				"err", err,
				"code", huresp.StatusCode,
				"body", string(b),
			)
			return errors.Join(err, ErrorFailedToReadResponse)
		}
		logger.Log(ctx, epg.logLevel, "execute request",
			"url", requrl.String(),
			"body", req.GetRequestBody(),
			// "response", b,
			"code", huresp.StatusCode,
			"resp", resp,
			"duration", time.Since(startTime).String(),
		)

		return nil
	}

	if huErr.RetCode == "-2" {
		return ErrorUserNeedToRelogin
	}

	logger.Warn("hu request failed",
		"err", huErr,
		"status", huresp.Status,
		"retCode", huErr.RetCode,
		"description", huErr.Description,
		"errorCode", huErr.ErrorCode,
		"body", string(b),
		"request", request.URL,
		"request_body", req.GetRequestBody(),
		"duration", time.Since(startTime).String(),
	)

	return huErr

}

func WithEPGADDRESS(epgAddress string) EPGOptons {
	return func(e *EPG) error {
		u, err := url.Parse(epgAddress)
		if err != nil {
			return err
		}
		e.epgAddress = u
		return nil
	}

}

func WithLogger(logger *slog.Logger) EPGOptons {
	return func(e *EPG) error {
		e.logger = logger
		return nil
	}
}

// with log level
func WithLogLevel(level slog.Level) EPGOptons {
	return func(e *EPG) error {
		e.logLevel = level
		return nil
	}
}

// with client
func WithClient(client http.Client) EPGOptons {
	return func(e *EPG) error {
		e.httpclient = client
		return nil
	}
}
