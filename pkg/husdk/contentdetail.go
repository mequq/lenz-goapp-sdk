package husdk

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrorVodNotFound = errors.New("vod not found")
)

type ContentDetailRequest struct {
	UserContentFilter *string `json:"userContentFilter,omitempty"`
	VOD               *string `json:"vod,omitempty"`
	MetaDataVer       *string `json:"metaDataVer,omitempty"`
	Channel           *string `json:"channel,omitempty"`
	Category          *string `json:"category,omitempty"`
	VAS               *string `json:"vas,omitempty"`
	Playbill          *string `json:"playbill,omitempty"`
	IDType            *string `json:"idType,omitempty"`
	FilterType        *int    `json:"filterType,omitempty"`
	Properties        []Prop  `json:"properties,omitempty"`
}

type ContentDetailResponse struct {
	VODList      []VOD      `json:"vodlist"`
	ChannelList  []Channel  `json:"channellist"`
	CategoryList []Category `json:"categorylist"`
	VASlist      []VAS      `json:"vaslist"`
	PlayBillList []PlayBill `json:"playbilllist"`
}

type ContentDetailOptions func(*ContentDetailRequest) error

func (c ContentDetailResponse) GetFirstVODs() (vod *VOD, err error) {
	if c.VODList == nil {
		return nil, ErrorVodNotFound
	}
	if len(c.VODList) < 1 {
		return nil, ErrorVodNotFound
	}

	return &c.VODList[0], nil
}

func (c *ContentDetailRequest) AddContent(id string) {
	newVod := fmt.Sprintf("%s,%s", *c.VOD, id)
	c.VOD = &newVod
}

func (c *ContentDetailRequest) SetVodID(id string) {
	c.VOD = &id
}

func NewContentDetailRequest(content string, opt ...ContentDetailOptions) HURequestInterface {
	cRequest := &ContentDetailRequest{}

	cRequest.VOD = &content

	for _, v := range opt {
		if err := v(cRequest); err != nil {
			return nil
		}
	}
	req := NewGeneralHuRequest("/EPG/JSON/ContentDetail", http.MethodPost, cRequest)
	return req
}

// with properties
func WithContentDetailProperties(properties []Prop) ContentDetailOptions {
	return func(c *ContentDetailRequest) error {
		c.Properties = properties
		return nil
	}
}
