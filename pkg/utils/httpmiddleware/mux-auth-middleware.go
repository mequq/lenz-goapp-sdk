package httpmiddleware

import (
	"context"
	"errors"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/httperror"

	"net/http"
	"strings"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/appcontext"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrorInvalidAuthHeader = httperror.New(http.StatusBadRequest, "invalid auth header", 1001, errors.New("invalid auth header"))
	ErrorBadAuthHeader     = httperror.New(http.StatusForbidden, "bad auth headers", 1002, errors.New("bad auth headers"))
	ErrorReloginNeeded     = httperror.New(http.StatusForbidden, "relogin needed", 1003, errors.New("relogin needed"))
	ErrorGuestLoginFailed  = httperror.New(http.StatusInternalServerError, "guest login failed", 1004, errors.New("guest login failed"))
	ErrotBadClameToken     = httperror.New(http.StatusForbidden, "token is valid but clame have not enough info", 1005, errors.New("token is valid but clame have not enough info"))
	ErrorAuthHeaderIsEmpty = errors.New("auth header is empty")
)

// NewGorillaMuxServer auth middleware
func (m *GorillaMuxMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authToken string
		var err error
		ctx := r.Context()
		//  check if header Authorization is set
		authToken = r.Header.Get("Authorization")
		if authToken == "" {
			m.logger.Warn("authToken", "authToken", authToken)
			ctx = context.WithValue(ctx, appcontext.REQUEST_IS_AUTH, false)
			ctx = context.WithValue(ctx, appcontext.REQUEST_AUTH_ERR, ErrorBadAuthHeader.Join(ErrorAuthHeaderIsEmpty))
			next.ServeHTTP(w, r.WithContext(ctx))
			return

		}
		token, err := parseAuthHeader(authToken, m.jwtSecret)
		m.logger.Info("token", "token", token, "err", err)
		ctx = setContextFromClame(ctx, token)
		ctx = context.WithValue(ctx, appcontext.REQUEST_AUTH_ERR, err)
		if err != nil {
			m.logger.Warn("invalid auth token", "token", authToken, "token", token, "err", err)
			ctx = context.WithValue(ctx, appcontext.REQUEST_IS_AUTH, false)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		ctx = context.WithValue(ctx, appcontext.REQUEST_IS_AUTH, true)
		//  set context from clame
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func setContextFromClame(ctx context.Context, clame map[string]interface{}) context.Context {
	if clame == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, appcontext.REQUEST_AUTH_CLAMS, clame)
	// set ip
	if ip, ok := clame["ip"].(string); ok {
		ctx = context.WithValue(ctx, appcontext.REQUEST_CLIENT_IP, ip)
	}
	return ctx

}

func parseAuthHeader(authHeader string, jwtSecret []byte) (map[string]interface{}, error) {

	authInfo := strings.Split(authHeader, "Bearer ")
	if len(authInfo) != 2 {
		return nil, ErrorInvalidAuthHeader.Join(errors.New("parse auth header failed"))
	}

	// parse token from header authorization bearer
	token, err := parseToken(authInfo[1], jwtSecret)
	if err != nil {
		return token.Claims.(jwt.MapClaims), ErrorReloginNeeded.Join(err).Join(errors.New("parse token failed"))
	}

	return token.Claims.(jwt.MapClaims), nil
}

// guest login
// func guestLogin(deviceType biz.DeviceType, clientIP net.IP, url *url.URL) (token string, err error) {
// 	client := &http.Client{
// 		Timeout: 2 * time.Second,
// 	}
// 	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
// 	if err != nil {
// 		return "", ErrorGuestLoginFailed.AddError(err)
// 	}
// 	req.Header.Set("Device-Type", string(deviceType))
// 	req.Header.Set("X-Forwarded-For", clientIP.String())
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", ErrorGuestLoginFailed.AddError(err)
// 	}
// 	defer resp.Body.Close()
// 	return resp.Header.Get("Authorization"), nil
// }

func parseToken(auth string, jwtSecret []byte) (*jwt.Token, error) {
	// parse token from header authorization bearer
	// return token string
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	return token, err

}
