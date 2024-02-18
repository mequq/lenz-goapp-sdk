package husdk

import (
	"net/http"
)

type ChannelListRequest struct {
	MetaDataVer      *string         `json:"metaDataVer,omitempty"`
	CategoryID       *string         `json:"categoryid,omitempty"`
	Count            *int            `json:"count,omitempty"`
	Offset           *int            `json:"offset,omitempty"`
	ContentType      *string         `json:"contenttype,omitempty"`
	Domain           *int            `json:"domain,omitempty"`
	ChannelNamespace *string         `json:"channelNamespace,omitempty"`
	FilterList       *NamedParameter `json:"filterlist,omitempty"`
	RetuenSatChannel *int            `json:"returnSatChannel,omitempty"`
}

type ChannelListResponse struct {
	CountTotal  int       `json:"countTotal,omitempty"`
	ChannelList []Channel `json:"channellist,omitempty"`
}

type ChannelListOptions func(*ChannelListRequest) error

func NewChannelListRequest(opt ...ChannelListOptions) HURequestInterface {
	cRequest := &ChannelListRequest{}

	for _, v := range opt {
		if err := v(cRequest); err != nil {
			return nil
		}
	}
	req := NewGeneralHuRequest("/EPG/JSON/ChannelList", http.MethodPost, cRequest)
	return req
}
