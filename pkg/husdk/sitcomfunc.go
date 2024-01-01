package husdk

func (v Sitcom) ContentType() VodType {

	switch v.VodType {
	case "0":
		switch v.SitcomNum {
		case "0":
			return VODTYPE_MOVIE
		default:
			return VODTYPE_EPISODE
		}
	case "1":
		return VODTYPE_SERIES
	case "2":
		return VODTYPE_SEASON
	default:
		return VODTYPE_UNKNOWN
	}
}

func (v Sitcom) GETID() string {
	return v.VODID
}

// vod
