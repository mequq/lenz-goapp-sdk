package husdk

import (
	"net/http"
)

type SitcomListRequest struct {
	Count  int    `json:"count"`
	Offset int    `json:"offset"`
	VodID  string `json:"vodid"`
}

type SitcomListResponse struct {
	CountTotal string `json:"counttotal"`
	VODList    []VOD
}

type SitcomListOptions func(*SitcomListRequest) error

func NewSitcomListRequest(content string, opt ...SitcomListOptions) HURequestInterface {
	huRequest := &SitcomListRequest{
		Count:  -1,
		Offset: 1,
		VodID:  content,
	}

	for _, v := range opt {
		if err := v(huRequest); err != nil {
			return nil
		}
	}
	req := NewGeneralHuRequest("/EPG/JSON/SitcomList", http.MethodPost, huRequest)
	return req
}

func WithOffsetAndCount(count int, offset int) SitcomListOptions {
	return func(s *SitcomListRequest) error {
		s.Count = count
		s.Offset = offset
		return nil
	}
}
