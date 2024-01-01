package husdk

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"git.abanppc.com/lenz-public/lenz-goapp-sdk/pkg/utils"
)

type VodType string

const (
	VODTYPE_UNKNOWN VodType = "Unknown"
	VODTYPE_AUDIO   VodType = "Audio"
	VODTYPE_MOVIE   VodType = "Movie"
	VODTYPE_SERIES  VodType = "Series"
	VODTYPE_EPISODE VodType = "Episode"
	VODTYPE_SEASON  VodType = "Season"
)

var (
	ErrorContentHasNoParrent = errors.New("content has no parent")
)

type VODContentGenre struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (v VOD) ContentType() VodType {
	switch v.Type {
	case "AUDIO_VOD":
		return VODTYPE_AUDIO
	case "VIDEO_VOD":
		switch v.VodType {
		case "0":
			if v.SitcomNum == nil {
				return VODTYPE_UNKNOWN
			}
			switch *v.SitcomNum {
			case "0":
				return VODTYPE_MOVIE
			default:
				return VODTYPE_EPISODE
			}
		case "1":
			return VODTYPE_SERIES
		case "2":
			return VODTYPE_SEASON
		}
	default:
		return VODTYPE_UNKNOWN
	}
	return VODTYPE_UNKNOWN

}

func (v VOD) HasParent() bool {
	if v.FatherVODList == nil || len(v.FatherVODList) < 1 {
		return false
	}

	return true

}

func (v VOD) GetParent() (*Sitcom, error) {
	if !v.HasParent() {
		return nil, ErrorContentHasNoParrent
	}

	return &(v.FatherVODList)[0], nil

}

func (v VOD) ParentType() VodType {
	if !v.HasParent() {
		return VODTYPE_UNKNOWN
	}
	return v.FatherVODList[0].ContentType()
}

func (v VOD) ParentID() (string, error) {
	if !v.HasParent() {
		return "", ErrorClientIPAddressNotValid
	}
	return v.FatherVODList[0].GETID(), nil
}

// GET title by language
func (v VOD) GetTitle(lang string) string {
	for _, title := range v.Names {
		if title.Key == lang {
			return title.Value
		}
	}
	return ""
}

// return duration of content in seconds
// return 0 if duration is not defined
func (v VOD) GetDeuration() uint64 {

	if v.Mediafiles == nil || len(v.Mediafiles) < 1 {
		return 0
	}
	duration := v.Mediafiles[0].ElapseTime
	// string to int64
	i, err := strconv.ParseInt(*duration, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(i)
}

// return duration of content in format x ساعت و y دقیقه
// return nil if duration is not defined
func (v VOD) GetDurationString() (string, error) {
	second := v.GetDeuration()
	if second < 60 {
		return "", nil
	}

	minute := second / 60
	hour := minute / 60
	minute = minute % 60

	if minute == 0 && hour > 0 {
		return fmt.Sprintf("%d ساعت", hour), nil
	}

	if hour == 0 && minute > 0 {
		return fmt.Sprintf("%d دقیقه", minute), nil
	}

	return fmt.Sprintf("%d ساعت و %d دقیقه", hour, minute), nil

}

// return content Generes
func (v VOD) GetGenres() []VODContentGenre {
	var genres []VODContentGenre
	if v.Genres == nil {
		return genres
	}
	allGenres := strings.Split(*v.Genres, ",")
	for i, genre := range allGenres {
		if len(allGenres) >= i+1 {
			genres = append(genres, VODContentGenre{ID: v.GenreIds[i], Title: genre})
		}
	}
	return genres
}

// return content PreductionDate
func (v VOD) GetPreductionDate() time.Time {
	if v.ProduceDate == nil {
		return time.Now().UTC()
	}
	t, err := time.Parse("2006-01-02", *v.ProduceDate)
	if err != nil {
		return time.Now().UTC()
	}
	return t
}

// get language of content
func (v VOD) GetLanguages() []string {
	if v.Languages == nil {
		return []string{"فارسی"}
	}
	isoLangs := strings.Split(*v.Languages, ",")
	return getListofPersianLangueges(isoLangs)
}

// get subtitles langs  of content
func (v VOD) GetSubtitles() []string {
	if v.Subtitles == nil {
		return nil
	}
	isoLangs := strings.Split(*v.Subtitles, ",")
	return getListofPersianLangueges(isoLangs)
}

// GetCast List
func (v VOD) GetCast() []CastInfo {
	return v.Casts
}

// GetCastID List
func (v VOD) GetCastIDs() []string {
	var ids []string
	for _, cast := range v.Casts {
		ids = append(ids, cast.CastID)
	}
	return ids
}

// general function to get list of persian languages
func getListofPersianLangueges(list []string) []string {
	var langs []string
	for _, l := range list {
		lang, err := utils.GetLanguageByCode(l)
		if err == nil {
			langs = append(langs, lang.PersianName)
		}
	}
	if len(langs) < 1 {
		return nil
	}
	return langs
}

// get list produceZone in iso 3166-1 code format
func (v VOD) GetProduceZone() []string {
	if v.ProduceZone == nil {
		return nil
	}
	zones := strings.Split(*v.ProduceZone, ",")
	if len(zones) < 1 {
		return nil
	}

	regions := make([]string, 0)

	for _, r := range zones {
		region, err := GetRegionByCode(r)
		if err == nil {
			regions = append(regions, region.RegionName)
		}
	}
	if len(regions) < 1 {
		return nil
	}

	return regions
}

// have subtitle or not
func (v VOD) HaveSubtitle() bool {
	if v.Subtitles == nil || len(*v.Subtitles) < 1 {
		return false
	}
	return true
}

// get avarage score of content
func (v VOD) GetScore() float64 {
	score, err := strconv.ParseFloat(v.AverageScore, 64)
	if err != nil {
		return 0
	}
	return score
}

type SeriesType string

const (
	SERIESTYPE_UNKNOWN   SeriesType = "Unknown"
	SERIESTYPE_ONE_LEVEL SeriesType = "OneLevel"
	SERIESTYPE_TWO_LEVEL SeriesType = "TwoLevel"
)

// get series Type
func (v VOD) GetSeriesType() SeriesType {
	if v.ContentType() != VODTYPE_SERIES || v.SubsetType == nil {
		return SERIESTYPE_UNKNOWN
	}
	switch *v.SubsetType { //TODO this seems in reverse order check via payam
	case "0":
		return SERIESTYPE_TWO_LEVEL
	case "1":
		return SERIESTYPE_ONE_LEVEL
	default:
		return SERIESTYPE_UNKNOWN
	}

}
