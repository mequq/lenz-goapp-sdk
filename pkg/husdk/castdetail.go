package husdk

import "net/http"

type CastDetailRequest struct {
	CastIDs []string `json:"castIds"`
}

type CastDetailResp struct {
	Casts []CastDetail `json:"casts"`
}

type CastDetailOptions func(*CastDetailRequest) error

func NewCastDetailRequest(casts []string, opt ...CastDetailOptions) HURequestInterface {
	cRequest := &CastDetailRequest{}

	cRequest.CastIDs = casts

	for _, v := range opt {
		if err := v(cRequest); err != nil {
			return nil
		}
	}
	req := NewGeneralHuRequest("/EPG/JSON/GetCastDetail", http.MethodPost, cRequest)
	return req
}
