package husdk

import "errors"

var (
	// Region is not found
	ErrorRegionNotFound = errors.New("region not found")
)

type Region struct {
	RegionID   string `json:"regionid"`
	RegionName string `json:"regionname"`
}

var Regions = map[string]Region{
	"0":  {RegionID: "0", RegionName: "نامشخص"},
	"1":  {RegionID: "1", RegionName: "آمریکا"},
	"2":  {RegionID: "2", RegionName: "ایران"},
	"3":  {RegionID: "3", RegionName: "آلمان"},
	"4":  {RegionID: "4", RegionName: "انگلیس"},
	"5":  {RegionID: "5", RegionName: "کانادا"},
	"6":  {RegionID: "6", RegionName: "ژاپن"},
	"7":  {RegionID: "7", RegionName: "کره جنوبی"},
	"8":  {RegionID: "8", RegionName: "هند"},
	"9":  {RegionID: "9", RegionName: "چین"},
	"10": {RegionID: "10", RegionName: "ایتالیا"},
	"11": {RegionID: "11", RegionName: "سوئد"},
	"12": {RegionID: "12", RegionName: "هنگ کنگ"},
	"13": {RegionID: "13", RegionName: "ایسلند"},
	"14": {RegionID: "14", RegionName: "آرژانتین"},
	"15": {RegionID: "15", RegionName: "فرانسه"},
	"16": {RegionID: "16", RegionName: "فلیپین"},
	"17": {RegionID: "17", RegionName: "اسپانیا"},
	"18": {RegionID: "18", RegionName: "اوکراین"},
	"19": {RegionID: "19", RegionName: "پرو"},
	"20": {RegionID: "20", RegionName: "برزیل"},
	"21": {RegionID: "21", RegionName: "ترکیه"},
	"22": {RegionID: "22", RegionName: "رومانی"},
	"23": {RegionID: "23", RegionName: "رومانی"},
	"24": {RegionID: "24", RegionName: "روسیه"},
	"25": {RegionID: "25", RegionName: "سوئد"},
	"26": {RegionID: "26", RegionName: "فنلاند"},
	"27": {RegionID: "27", RegionName: "مالزی"},
	"28": {RegionID: "28", RegionName: "مجارستان"},
	"29": {RegionID: "29", RegionName: "هلند"},
	"30": {RegionID: "30", RegionName: "مکزیک"},
	"31": {RegionID: "31", RegionName: "استرالیا"},
	"32": {RegionID: "32", RegionName: "آفریقا جنوبی"},
	"33": {RegionID: "33", RegionName: "نروژ"},
}

// func get region by id
func GetRegionByCode(id string) (*Region, error) {
	if r, ok := Regions[id]; ok {
		return &r, nil
	}
	return nil, ErrorRegionNotFound
}
