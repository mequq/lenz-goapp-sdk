package husdk

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
)

type HURequestInterface interface {
	GetPath() (path string)
	SetPath(path string)
	AddQueryParameter(key string, val string)
	GetQueryParameter() url.Values
	GetMethod() string
	GetReuestData() (io.Reader, error)
	GetRequestBody() any
}

type HUResponseInterface interface {
	GetRetCode()
}

type HURequest struct {
	path           string
	queryParameter url.Values
	method         string
	data           any
}

func NewGeneralHuRequest(path string, method string, data any) HURequestInterface {
	return &HURequest{
		path:   path,
		method: method,
		data:   data,
	}
}

func (h HURequest) GetPath() string {
	return h.path
}

func (h *HURequest) SetPath(path string) {
	h.path = path
}

func (h *HURequest) AddQueryParameter(key string, val string) {
	h.queryParameter.Add(key, val)
}

func (h HURequest) GetQueryParameter() url.Values {
	return h.queryParameter
}

func (h HURequest) GetMethod() string {
	return h.method
}

func (h HURequest) GetReuestData() (io.Reader, error) {
	b, err := json.Marshal(h.data)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

func (h HURequest) GetRequestBody() any {
	return h.data
}
