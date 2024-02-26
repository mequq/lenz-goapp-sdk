package region

import (
	"context"
	"encoding/json"
	"fmt"
	region "git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils/regiondetector"
	"io"
	"net/http"
	"os"
)

type Country string

const (
	CopyrightedError string = ""
	OutOfIranCatId   string = "8010"
)

type ChannelState struct {
	ChannelVOID  string `json:"channel_id" validate:"required,number,max=25"`
	Name         string `json:"channel_name"`
	Desc         string `json:"channel_desc"`
	DefaultAvail bool   `json:"default_avail"`
	Available    bool   `json:"available"`
}

func IsCopyrightedVod(catIds []string, ip string) bool {
	isCopyrighted := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			fmt.Println("caught panic:", r)
		}
	}()
	if catIds != nil && len(catIds) > 0 {
		isCopyrighted = IsOutOfIran(ip) && HasCopyrightFlag(catIds)
	}
	return isCopyrighted
}

func GetVodCopyrightDetail(catIds []string, ip string) (bool, bool) {
	isCopyrighted := true
	isOutOfIran := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			isOutOfIran = true
			fmt.Println("caught panic:", r)
		}
	}()
	if catIds != nil && len(catIds) > 0 {
		isOutOfIran = IsOutOfIran(ip)
		isCopyrighted = isOutOfIran && HasCopyrightFlag(catIds)
	}
	return isCopyrighted, isOutOfIran
}

func IsOutOfIran(ip string) bool {
	isOutOfIran := true
	defer func() {
		if r := recover(); r != nil {
			isOutOfIran = true
			fmt.Println("caught panic:", r)
		}
	}()
	detector, err := region.NewRegionDetector(region.WithEndpoint(os.Getenv("REGION_DOMAIN")))
	if err != nil {
		return true
	}
	isOutOfIran = !detector.IsIranIP(context.Background(), ip)
	return isOutOfIran
}

func HasCopyrightFlag(catIds []string) bool {
	hasFlag := false
	for _, cat := range catIds {
		if cat == OutOfIranCatId {
			hasFlag = true
			break
		}
	}
	return !hasFlag
}

func IsCopyrighted(isLive bool, ids []string, ip string) bool {
	isCopyrighted := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			fmt.Println("caught panic:", r)
		}
	}()
	if ids != nil && len(ids) > 0 {
		if isLive {
			isCopyrighted = IsCopyrightedLive(ids[0], ip)
		} else { //vod
			isCopyrighted = IsCopyrightedVod(ids, ip)
		}
	}
	return isCopyrighted
}

func GetContentCopyrightDetail(isLive bool, ids []string, ip string) (bool, bool) {
	isCopyrighted := true
	isOutOfIran := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			isOutOfIran = true
			fmt.Println("caught panic:", r)
		}
	}()
	if ids != nil && len(ids) > 0 {
		if isLive {
			isCopyrighted, isOutOfIran = GetLiveCopyrightDetail(ids[0], ip)
		} else { //vod
			isCopyrighted, isOutOfIran = GetVodCopyrightDetail(ids, ip)
		}
	}
	return isCopyrighted, isOutOfIran
}

func IsCopyrightedLive(channelId string, ip string) bool {
	isCopyrighted := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			fmt.Println("caught panic:", r)
		}
	}()
	isCopyrighted = IsOutOfIran(ip) && !IsChannelAvailable(channelId)
	return isCopyrighted
}

func GetLiveCopyrightDetail(channelId string, ip string) (bool, bool) {
	isCopyrighted := true
	isOutOfIran := true
	defer func() {
		if r := recover(); r != nil {
			isCopyrighted = true
			isOutOfIran = true
			fmt.Println("caught panic:", r)
		}
	}()
	isOutOfIran = IsOutOfIran(ip)
	isCopyrighted = isOutOfIran && !IsChannelAvailable(channelId)
	return isCopyrighted, isOutOfIran
}

func IsChannelAvailable(channelId string) bool {
	isAvailable := false
	defer func() {
		if r := recover(); r != nil {
			isAvailable = false
			fmt.Println("caught panic:", r)
		}
	}()
	response, err := http.Get(os.Getenv("LIVE_STATE_DOMAIN") + "/api/v1/admin/live-state/channel/" + channelId)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	var channelState *ChannelState
	err = json.Unmarshal(body, &channelState)
	if err != nil {
		return false
	}

	isAvailable = channelState.Available
	return isAvailable
}
