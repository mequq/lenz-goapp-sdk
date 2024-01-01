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

// NewGorillaMuxServer auth middleware
// func (m *GorillaMuxMiddleware) GuestLoginMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		r.Header.Set("Is-Open-Api", "True")
// 		w.Header().Set("Is-Open-Api", "True")

// 		ctx := r.Context()
// 		//  check if header Authorization is set
// 		deviceType := biz.GetDeviceType(r.Header.Get("User-Agent"))
// 		if deviceType == biz.UnknownDevice {
// 			deviceType = biz.AndroidDevice
// 		}

// 		cientIPString := r.Header.Get("X-Forwarded-For")
// 		ip := net.ParseIP(cientIPString)
// 		if ip == nil {
// 			m.logger.Error("invalid client ip", "ip", cientIPString)
// 			httperror.ConvertError(ErrorBadAuthHeader).HttpError(w)
// 			return
// 		}

// 		var authToken string
// 		authToken = r.Header.Get("Authorization")
// 		m.logger.Debug("authToken", "authToken", authToken, "clientIPString", cientIPString, "ip", ip)
// 		if authToken == "" {
// 			if m.loginUri == nil {
// 				panic("login uri is not set")
// 			}

// 			t, err := guestLogin(deviceType, ip, m.loginUri)
// 			if err != nil {
// 				m.logger.Error("guest login failed", "err", err)
// 				httperror.ConvertError(err).HttpError(w)
// 				return
// 			}

// 			r.Header.Set("Authorization", t)
// 			w.Header().Set("Authorization", t)

// 			authToken = t
// 		}
// 		token, err := parseAuthHeader(authToken, m.jwtSecret)
// 		m.logger.Debug("token", "token", token, "err", err)
// 		if err != nil {
// 			m.logger.Error("invalid auth token", "token", authToken, "err", err)
// 			httperror.ConvertError(err).HttpError(w)
// 			return
// 		}
// 		isGuest := false
// 		if guest, ok := token["is_guest"].(bool); ok {
// 			isGuest = guest
// 		}

// 		if isGuest {
// 			r.Header.Set("Is-Guest", "True")
// 			w.Header().Set("Is-Guest", "True")
// 		}

// 		jwtIP := token["ip"].(string)
// 		m.logger.Debug("jwtIP", "jwtIP", jwtIP, "cientIPString", cientIPString)
// 		if jwtIP != cientIPString {
// 			switch isGuest {
// 			case true:
// 				t, err := guestLogin(deviceType, ip, m.loginUri)
// 				if err != nil {
// 					m.logger.Error("guest login failed", "err", err)
// 					httperror.ConvertError(err).HttpError(w)
// 					return
// 				}
// 				r.Header.Set("Authorization", t)
// 				w.Header().Set("Authorization", t)
// 				token, err = parseAuthHeader(t, m.jwtSecret)
// 				if err != nil {
// 					m.logger.Error("invalid auth token", "token", authToken, "err", err)
// 					httperror.ConvertError(err).HttpError(w)
// 				}
// 			case false:
// 				m.logger.Warn("invalid auth token", "token", authToken)
// 				httperror.ConvertError(ErrorReloginNeeded).Join(errors.New("ip is mismatch")).HttpError(w)
// 				return

// 			}
// 		}

// 		m.logger.Debug("token", "token", token)
// 		if err := checkNessaryClame(token); err != nil {
// 			m.logger.Error("invalid auth token", "token", authToken, "err", err)
// 			httperror.ConvertError(err).HttpError(w)
// 			return
// 		}
// 		ctx = setContextFromClame(ctx, token)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func checkNessaryClame(clame map[string]interface{}) error {
// 	if clame == nil {
// 		return ErrotBadClameToken.Join(errors.New("clame is nil"))
// 	}
// 	// check if clame have ip
// 	if _, ok := clame["ip"]; !ok {
// 		return ErrotBadClameToken.Join(errors.New("ip is not set"))
// 	}
// 	// check if clame have msisdn
// 	if _, ok := clame["user_id"]; !ok {
// 		return ErrotBadClameToken.Join(errors.New("user_id is not set"))
// 	}

// 	//  chek if clame have backend
// 	if _, ok := clame["backend"]; !ok {
// 		return ErrotBadClameToken.Join(errors.New("backend is not set"))
// 	}

// 	//  chek if clame have jsessionid
// 	if _, ok := clame["jsessionid"]; !ok {
// 		return ErrotBadClameToken.Join(errors.New("jsessionid is not set"))
// 	}

// 	return nil
// }

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
