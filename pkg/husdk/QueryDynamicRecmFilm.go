package husdk

import "net/http"

// QueryDynamicRecmFilmRequest is the request struct for api QueryDynamicRecmFilm
type QueryDynamicRecmFilmRequest struct {
	VODID             *string  `json:"vodid"`
	Count             int      `json:"count"`
	Offset            int      `json:"offset"`
	Type              *int     `json:"type,omitempty"`
	Scenario          *string  `json:"scenario,omitempty"`
	Network           *int     `json:"network,omitempty"`
	RecmdRegionID     *string  `json:"recmdRegionId,omitempty"`
	PreferContentIds  []string `json:"preferContentIds,omitempty"`
	ContentType       *string  `json:"contentType,omitempty"`
	BusinessType      *int     `json:"businessType,omitempty"`
	UnknownContentIds []string `json:"unknownContentIds,omitempty"`
	DislikeContentIds []string `json:"dislikeContentIds,omitempty"`
	MetaDataVer       *string  `json:"metaDataVer,omitempty"`
	Properties        []Prop   `json:"properties,omitempty"`
	Action            *int     `json:"action,omitempty"`
	OrderType         *int     `json:"orderType,omitempty"`
}

type QueryDynamicRecmFilmOptions func(*QueryDynamicRecmFilmRequest) error

// response
type QueryDynamicRecmFilmVODResponse struct {
	CountTotal  string `json:"counttotal"`
	ContentList []VOD  `json:"contentlist"`
}

// NewQueryDynamicRecmFilmRequest instantiates a new QueryDynamicRecmFilmRequest object
func NewQueryDynamicRecmFilmVODRequest(vodID string, count int, offset int, opt ...QueryDynamicRecmFilmOptions) HURequestInterface {

	qdctype := "VIDEO_VOD"
	qdtype := 1

	cRequest := &QueryDynamicRecmFilmRequest{
		VODID:       &vodID,
		Count:       count,
		Offset:      offset,
		ContentType: &qdctype,
		Type:        &qdtype,
	}

	for _, o := range opt {
		o(cRequest)
	}

	req := NewGeneralHuRequest("/EPG/JSON/QueryDynamicRecmFilm", http.MethodPost, cRequest)
	return req
}

// with properties
func WithQueryDynamicRecmFilmProperties(properties []Prop) QueryDynamicRecmFilmOptions {
	return func(c *QueryDynamicRecmFilmRequest) error {
		c.Properties = properties
		return nil
	}
}

// with count
func WithQueryDynamicRecmFilmCount(count int) QueryDynamicRecmFilmOptions {
	return func(c *QueryDynamicRecmFilmRequest) error {
		c.Count = count
		return nil
	}
}

// with offset
func WithQueryDynamicRecmFilmOffset(offset int) QueryDynamicRecmFilmOptions {
	return func(c *QueryDynamicRecmFilmRequest) error {
		c.Offset = offset
		return nil
	}
}
