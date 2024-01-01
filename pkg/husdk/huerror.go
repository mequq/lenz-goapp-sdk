package husdk

import "fmt"

type GeneralError struct {
	RetCode     string `json:"retcode"`
	Description string `json:"desc"`
	ErrorCode   string `json:"errorCode"`
}

func (hue GeneralError) Error() string {
	return fmt.Sprintf("filed hu with retcode: %s and description %s", hue.RetCode, hue.Description)
}
