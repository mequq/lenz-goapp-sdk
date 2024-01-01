package httperror

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var DebugMode = false

type AppErr struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
	Err       error  `json:"err"`
}

func New(code int, message string, errorCode int, err error) *AppErr {
	return &AppErr{
		Code:      code,
		Message:   message,
		ErrorCode: errorCode,
		Err:       err,
	}
}

// MarshalJSON
func (e *AppErr) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		ErrorCode int    `json:"error_code"`
		Errors    string `json:"errors,omitempty"`
	}{
		Code:      e.Code,
		Message:   e.Message,
		ErrorCode: e.ErrorCode,
		Errors:    e.Err.Error(),
	})
}

// error
func (e *AppErr) Error() string {
	return fmt.Sprintf("code: %d, message: %s, error_code: %d, errors: %v", e.Code, e.Message, e.ErrorCode, errors.Unwrap(e.Err))
}

// join error
func (e *AppErr) Join(err error) *AppErr {
	switch err := err.(type) {
	case *AppErr:
		e.Err = errors.Join(e.Err, err.Err)
	default:
		e.Err = errors.Join(e.Err, err)
	}

	return e
}

func (e *AppErr) LogValue() slog.Value {

	return slog.GroupValue(
		slog.Int("errorCode", e.ErrorCode),
		slog.Int("statusCode", e.Code),
		slog.String("message", e.Message),
		slog.Any("errors", e.Err),
	)
}

// http error
func (e *AppErr) httpErrorDebug(w http.ResponseWriter) {

	w.WriteHeader(e.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(e)
}

func (e *AppErr) httpError(w http.ResponseWriter) {

	w.WriteHeader(e.Code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(&struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		ErrorCode int    `json:"error_code"`
	}{
		Code:      e.Code,
		Message:   e.Message,
		ErrorCode: e.ErrorCode,
	})
}

func (e *AppErr) HttpError(w http.ResponseWriter) {
	if DebugMode {
		e.httpErrorDebug(w)
		return
	}
	e.httpError(w)
}

// is error
func (e *AppErr) Is(err error) bool {
	return errors.Is(e.Err, err)
}

// convert error to http error
func ConvertError(err error) *AppErr {
	switch err := err.(type) {
	case *AppErr:
		return err
	default:
		return New(500, "internal server error", 1000, err)
	}
}
