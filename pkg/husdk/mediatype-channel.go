package husdk

type Channel struct {
	ChannelID           string   `json:"id"`
	ForeignSN           string   `json:"foreignsn"`
	Name                string   `json:"name"`
	Type                string   `json:"type"`
	Introduce           *string  `json:"introduce"`
	PreviewEnable       string   `json:"previewenable"`
	PreviewLength       string   `json:"previewlength"`
	PreviewCount        string   `json:"previewcount"`
	MulticastSourceIP   *string  `json:"multicastsourceip"`
	HasPIP              string   `json:"haspip"`
	PIPMulicastIP       *string  `json:"pipmulticastip"`
	PIPMulicastPort     *string  `json:"pipmulticastport"`
	PIPMulicastSourceIP *string  `json:"pipmulticastsourceip"`
	PIPFCCEnable        *string  `json:"pipfccenable"`
	Status              *string  `json:"status"`
	ChanNO              *string  `json:"channo"`
	FCCEnable           *string  `json:"fccenable"`
	Encrypt             string   `json:"encrypt"`
	Bitrate             string   `json:"bitrate"`
	PlayURL             string   `json:"playurl"`
	Definition          string   `json:"definition"`
	Picture             *Picture `json:"picture"`
	ISCPVR              string   `json:"iscpvr"`
	ISPLTV              string   `json:"ispltv"`
	PLTTVLength         *string  `json:"pltvlength"`
	ISTVOD              string   `json:"istvod"`
	ISLocalTimeShift    string   `json:"islocaltimeshift"`
	Logo                *string  `json:"logo"`
	ISFavrited          string   `json:"isfavrited"`
	RatingID            string   `json:"ratingid"`
	ISSubscribed        string   `json:"issubscribed"`
	ISPPV               string   `json:"isppv"`
	CPVRSubscribed      string   `json:"cpvrsubscribed"`
	RecordLength        *string  `json:"recordlength"`
	SLSType             string   `json:"slstype"`
}
