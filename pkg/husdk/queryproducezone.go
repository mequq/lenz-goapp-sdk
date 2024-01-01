package husdk

import "net/http"

type QueryProduceZoneResponse struct {
	Result *string `json:"result"`
	IDList *string `json:"idList"`
}

func NewQueryProduceZone() HURequestInterface {

	req := NewGeneralHuRequest("/EPG/JSON/QueryProduceZone", http.MethodPost, nil)
	return req
}
