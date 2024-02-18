package husdk

import (
	"errors"
	"fmt"
	"math"

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
	for _, vod := range v.FatherVODList {
		if vod.VODID != "15610" {
			return vod.VODID, nil
		}
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

	// calc rating old fashion
	rateCount, err1 := strconv.Atoi(v.StaticTimes)
	rateSum, err2 := strconv.Atoi(v.ScoreSum)
	if err1 == nil && err2 == nil && rateCount > 0 {
		return math.Round(float64(rateSum)/float64(rateCount)*10.0) / 10.0

	}
	// else calc rating new fashion
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

// Favorited
func (v VOD) Favorite() bool {
	switch v.IsFavorited {
	case "0":
		return false
	case "1":
		return true
	default:
		return false
	}
}

// Get Downloadable
func (v VOD) Downloadable() bool {
	if v.Mediafiles == nil || len(v.Mediafiles) < 1 {
		return false
	}
	return true
}

// get share url
func (v VOD) GetShareURL(baseUrl string) string {
	return fmt.Sprintf("%s را ببینید\n%s/video/%s", v.Name, baseUrl, v.ID)
}

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

// Get ID
func (v VOD) GETINTID() (id int, err error) {
	return strconv.Atoi(v.ID)
}

// get Content Rated
func (v VOD) GetRated() string {
	switch v.Ratingid {

	case "4":
		return "7"
	case "6":
		return "12"
	case "8":
		return "16"
	case "9":
		return "18"
	default:
		return ""
	}
}

// get Pictures
func (v VOD) GetPictures() *Picture {
	if v.Picture == nil {
		return nil
	}
	if v.Picture.BackGround == nil {
		if v.Picture.Deflate != nil {
			v.Picture.BackGround = v.Picture.Deflate
		} else if v.Picture.Icon != nil {
			v.Picture.BackGround = v.Picture.Icon
		} else if v.Picture.Still != nil {
			v.Picture.BackGround = v.Picture.Still
		} else if v.Picture.Title != nil {
			v.Picture.BackGround = v.Picture.Title
		} else if v.Picture.AD != nil {
			v.Picture.BackGround = v.Picture.AD
		} else if v.Picture.Draft != nil {
			v.Picture.BackGround = v.Picture.Draft
		} else if v.Picture.Channelpic != nil {
			v.Picture.BackGround = v.Picture.Channelpic
		} else if v.Picture.BlackWhite != nil {
			v.Picture.BackGround = v.Picture.BlackWhite
		} else if v.Picture.ChanName != nil {
			v.Picture.BackGround = v.Picture.ChanName
		} else if v.Picture.Other != nil {
			v.Picture.BackGround = v.Picture.Other
		}
	}
	if v.Picture.Poster == nil {
		if v.Picture.Icon != nil {
			v.Picture.Poster = v.Picture.Icon
		} else if v.Picture.Still != nil {
			v.Picture.Poster = v.Picture.Still
		} else if v.Picture.Title != nil {
			v.Picture.Poster = v.Picture.Title
		} else if v.Picture.AD != nil {
			v.Picture.Poster = v.Picture.AD
		} else if v.Picture.Draft != nil {
			v.Picture.Poster = v.Picture.Draft
		} else if v.Picture.Channelpic != nil {
			v.Picture.Poster = v.Picture.Channelpic
		} else if v.Picture.BlackWhite != nil {
			v.Picture.Poster = v.Picture.BlackWhite
		} else if v.Picture.ChanName != nil {
			v.Picture.Poster = v.Picture.ChanName
		} else if v.Picture.Other != nil {
			v.Picture.Poster = v.Picture.Other
		}
	}

	if v.Picture.Icon == nil {
		if v.Picture.Poster != nil {
			v.Picture.Icon = v.Picture.Poster
		} else if v.Picture.BackGround != nil {
			v.Picture.Icon = v.Picture.BackGround
		} else if v.Picture.Deflate != nil {
			v.Picture.Icon = v.Picture.Deflate
		} else if v.Picture.Draft != nil {
			v.Picture.Icon = v.Picture.Draft
		} else if v.Picture.Channelpic != nil {
			v.Picture.Icon = v.Picture.Channelpic
		} else if v.Picture.BlackWhite != nil {
			v.Picture.Icon = v.Picture.BlackWhite
		} else if v.Picture.ChanName != nil {
			v.Picture.Icon = v.Picture.ChanName
		} else if v.Picture.Other != nil {
			v.Picture.Icon = v.Picture.Other
		}
	}
	return v.Picture
}
