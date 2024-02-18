package appcontext

import (
	"context"
	"net/http"
	"strings"

	"log/slog"
)

type appContextKey int

const (
	REQUEST_ID appContextKey = iota
	REQUEST_CLIENT_IP
	REQUEST_AUTH_CLAMS
	REQUEST_MSISDN
	REQUEST_EPG_BACKEND
	REQUEST_EPG_SESSION
	REQUEST_IS_GUEST
	REQUEST_IS_AUTH
	REQUEST_AUTH_ERR
	REQUEST_IS_IRAN
	REQUEST_IS_MTNI
)

type AppContext struct {
	// Context
	ctx context.Context
	// Logger
}

// Log Context  is a function that logs the context in a slog.Value
// usage example: logger.Info("message","ctx", utils.LogContext(ctx))
func LogContext(ctx context.Context) *AppContext {
	return &AppContext{
		ctx: ctx,
	}
}

// logValue returns a slog.Value with all the context values
// this methoth is used by the slog library to log the context
// note: this method is not used directly use LogContext instead
func (a *AppContext) LogValue() slog.Value {

	var attrs []slog.Attr

	if a.ctx.Value(REQUEST_AUTH_CLAMS) != nil {
		attrs = append(attrs, slog.Any("authClams", a.ctx.Value(REQUEST_AUTH_CLAMS)))
	}

	if a.ctx.Value(REQUEST_CLIENT_IP) != nil {
		attrs = append(attrs, slog.String("clientIP", a.ctx.Value(REQUEST_CLIENT_IP).(string)))
	}

	if a.ctx.Value(REQUEST_ID) != nil {
		attrs = append(attrs, slog.String("requestID", a.ctx.Value(REQUEST_ID).(string)))
	}
	// msisdn
	if a.ctx.Value(REQUEST_MSISDN) != nil {
		attrs = append(attrs, slog.String("msisdn", a.ctx.Value(REQUEST_MSISDN).(string)))
	}

	// backend
	if a.ctx.Value(REQUEST_EPG_BACKEND) != nil {
		attrs = append(attrs, slog.String("backend", a.ctx.Value(REQUEST_EPG_BACKEND).(string)))
	}

	// isGuest
	if a.ctx.Value(REQUEST_IS_GUEST) != nil {
		attrs = append(attrs, slog.Bool("isGuest", a.ctx.Value(REQUEST_IS_GUEST).(bool)))
	}

	// isAuth
	if a.ctx.Value(REQUEST_IS_AUTH) != nil {
		attrs = append(attrs, slog.Bool("isAuth", a.ctx.Value(REQUEST_IS_AUTH).(bool)))
	}

	// authErr
	if a.ctx.Value(REQUEST_AUTH_ERR) != nil {
		attrs = append(attrs, slog.Any("authErr", a.ctx.Value(REQUEST_AUTH_ERR)))
	}

	// isIran
	if a.ctx.Value(REQUEST_IS_IRAN) != nil {
		attrs = append(attrs, slog.Bool("isIran", a.ctx.Value(REQUEST_IS_IRAN).(bool)))
	}

	// isMTNI
	if a.ctx.Value(REQUEST_IS_MTNI) != nil {
		attrs = append(attrs, slog.Bool("isMTNI", a.ctx.Value(REQUEST_IS_MTNI).(bool)))
	}

	attrs = append(attrs, slog.Any("err", a.ctx.Err()))
	return slog.GroupValue(attrs...)
}

// SetContextFromHttpReq set context from http request
// usage example: ctx = utils.SetContextFromHttpReq(ctx, r)
// set requestID from x-request-id header and clientIP from x-forwarded-for header or remoteAddr
func SetContextFromHttpReq(ctx context.Context, r *http.Request) context.Context {
	nCtx := context.WithValue(ctx, REQUEST_ID, r.Header.Get("x-request-id"))
	var requestIP string
	if r.Header.Get("x-forwarded-for") != "" {
		requestIP = r.Header.Get("x-forwarded-for")
	} else {
		requestIP = strings.Split(r.RemoteAddr, ":")[0]
	}
	nCtx = context.WithValue(nCtx, REQUEST_CLIENT_IP, requestIP)
	return nCtx
}
