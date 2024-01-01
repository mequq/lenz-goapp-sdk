package husdk

type VOD struct {
	ID                   string           `json:"id"`
	Name                 string           `json:"name"`
	Names                []NamedParameter `json:"names"`
	Type                 string           `json:"type"`
	Introduce            *string          `json:"introduce"`
	Picture              *Picture         `json:"picture"`
	Cast                 *Cast            `json:"cast"`
	Casts                []CastInfo       `json:"casts"`
	VodType              string           `json:"vodType"`
	SeriesType           *string          `json:"seriesType"`
	SitcomNum            *string          `json:"sitcomnum"`
	EpisodeTotalCount    *string          `json:"episodeTotalCount"`
	Price                *string          `json:"price"`
	Tags                 *string          `json:"tags"`
	Genres               *string          `json:"genres"`
	GenreIds             []string         `json:"genreIds"`
	Mediafiles           []VodMediaFile   `json:"mediafiles"`
	Clipfiles            []VodMediaFile   `json:"clipfiles"`
	Chapters             []Chapter        `json:"chapters"`
	Rentperiod           string           `json:"rentperiod"`
	Ratingid             string           `json:"ratingid"`
	StartTime            string           `json:"starttime"`
	EndTime              string           `json:"endtime"`
	IsFavorited          string           `json:"isfavorited"`
	IsSubscribed         string           `json:"issubscribed"`
	SubscriptionType     string           `json:"subscriptionType"`
	IsRefresh            string           `json:"isRefresh"`
	Restriction          *string          `json:"restriction"`
	ProduceDate          *string          `json:"producedate"`
	ForeignSN            string           `json:"foreignsn"`
	FatherVODList        []Sitcom         `json:"fathervodlist"`
	StaticTimes          string           `json:"statictimes"`
	AverageScore         string           `json:"averagescore"`
	ScoreSum             string           `json:"scoresum"`
	LoyaltyCount         *string          `json:"loyaltyCount"`
	Languages            *string          `json:"languages"`
	ProduceZone          *string          `json:"produceZone"`
	CreditVodSendLoyalty *string          `json:"creditVodSendLoyalty"`
	VisitTimes           *string          `json:"visittimes"`
	Country              *string          `json:"country"`
	Subtitles            *string          `json:"subtitles"`
	Awards               *string          `json:"awards"`
	DeviceGroups         []DeviceGroup    `json:"deviceGroups"`
	Keyword              *string          `json:"keyword"`
	Advisory             []string         `json:"advisory"`
	IsPlayable           *string          `json:"IsPlayable"`
	CategoryIDs          []string         `json:"categoryIds"`
	ContentRating        *string          `json:"contentRating"`
	SelfDefineLabels     []string         `json:"selfdefineLabels"`
	ReviewScores         *string          `json:"reviewScores"`
	ReviewID             *string          `json:"reviewId"`
	ProgramType          *string          `json:"programType"`
	SortName             *string          `json:"sortName"`
	ExtensionInfo        []NamedParameter `json:"extensionInfo"`
	PriceType            *string          `json:"priceType"`
	IsLoyalty            *string          `json:"isLoyalty"`
	ExternalContentCode  *string          `json:"externalContentCode"`
	Company              *string          `json:"company"`
	CompanyName          *string          `json:"companyName"`
	LocationCopyrights   []string         `json:"locationCopyrights"`
	ADFlag               *string          `json:"adFlag"`
	SubsetType           *string          `json:"subsetType"`
	LifeTimeID           *string          `json:"lifetimeId"`
	TitleSortName        *string          `json:"titleSortName"`
}

type Channel struct {
	ID string `json:"id"`
}

type Category struct {
	ID string `json:"id"`
}

type VAS struct {
	ID string `json:"id"`
}

type PlayBill struct {
	ID string `json:"id"`
}

type NamedParameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Picture struct {
	Deflate    *string `json:"deflate,omitempty"`
	Poster     *string `json:"poster,omitempty"`
	Still      *string `json:"still,omitempty"`
	Icon       *string `json:"icon,omitempty"`
	Title      *string `json:"title,omitempty"`
	AD         *string `json:"ad,omitempty"`
	Draft      *string `json:"draft,omitempty"`
	BackGround *string `json:"background,omitempty"`
	Channelpic *string `json:"channelpic,omitempty"`
	BlackWhite *string `json:"blackwhite,omitempty"`
	ChanName   *string `json:"channame,omitempty"`
	Other      *string `json:"other,omitempty"`
}

type Cast struct {
	Actor    *string `json:"actor"`
	Director *string `json:"director"`
	Producer *string `json:"producer"`
	Adaptor  *string `json:"adaptor"`
}

type CastInfo struct {
	CastID   string  `json:"castId"`
	RoleType string  `json:"roleType"`
	CastName string  `json:"castName"`
	CastCode *string `json:"castCode"`
}

type VodMediaFile struct {
	ID                 string            `json:"id"`
	ElapseTime         *string           `json:"elapsetime"`
	Bitrate            *string           `json:"bitrate"`
	IsDownload         *string           `json:"isdownload"`
	Definition         *string           `json:"definition"`
	NDefinition        *string           `json:"nDefinition"`
	HDCPEnable         *string           `json:"HDCPEnable"`
	MacroVision        *string           `json:"macrovision"`
	Dimension          *string           `json:"dimension"`
	FormatOf3D         *string           `json:"formatOf3D"`
	SupportTerminal    *string           `json:"supportTerminal"`
	FileFormat         string            `json:"fileFormat"`
	Encrypt            string            `json:"encrypt"`
	CGMSA              string            `json:"CGMSA"`
	AnalogOutputEnable string            `json:"analogOutputEnable"`
	VideoCodec         *string           `json:"videoCodec"`
	SPID4VIP           *string           `json:"spId4VIP"`
	ExtensionInfo      []NamedParameter  `json:"extensionInfo"`
	AudioType          *string           `json:"audioType"`
	ExternalMediaCode  *string           `json:"externalMediaCode"`
	MultiBitRate       *string           `json:"multiBitRate"`
	FPS                *string           `json:"fps"`
	MaxBitRate         *string           `json:"maxBitRate"`
	Picture            *Picture          `json:"picture"`
	VODBR              *BusinessRight    `json:"vodBR"`
	Preview            *string           `json:"preview"`
	PreviewStart       *string           `json:"previewStart"`
	Time               *string           `json:"Time"`
	PreviewEndTime     *string           `json:"previewEndTime"`
	MediaProfiles      []ProfileMetadata `json:"mediaProfiles"`
}

type Chapter struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	OffTime int      `json:"offtime"`
	Picture *Picture `json:"picture"`
}

type Sitcom struct {
	VODID     string `json:"vodid"`
	SitcomNum string `json:"sitcomnum"`
	VodType   string `json:"vodtype"`
	Name      string `json:"name"`
}

type DeviceGroup struct {
	GroupID   string `json:"groupId"`
	GroupName string `json:"groupName"`
	GroupType string `json:"groupType"`
}

type ProfileMetadata struct {
	MediaProfiled         *string `json:"mediaProfiled"`
	MediaProfileType      *string `json:"mediaProfileType"`
	AVEBitrate            *string `json:"aveBitrate"`
	MaxBitrate            *string `json:"maxBitrate"`
	MediaProfileTotalSize *string `json:"mediaProfileTotalSize"`
}

type BusinessRight struct {
	IS string      `json:"is"` // is subscribed
	VA *string     `json:"va"` //is valid
	R  []Condition `json:"r"`  //restrictionList
	PT *string     `json:"pt"` //price type
	IL *string     `json:"il"` //is loyalty
	ST *string     `json:"st"` //subscribe type
}

type Condition struct {
	N string   `json:"n"`
	V []string `json:"v"`
	T string   `json:"t"`
}

type CastDetail struct {
	CastID     string        `json:"castId"`
	Name       string        `json:"name"`
	FirstName  *string       `json:"firstName"`
	MiddleName *string       `json:"middleName"`
	LastName   *string       `json:"lastName"`
	Favorite   *string       `json:"favorite"`
	Sex        *string       `json:"sex"`
	Birthday   *string       `json:"birthday"`
	HomeTown   *string       `json:"hometown"`
	Education  *string       `json:"education"`
	Height     *string       `json:"height"`
	Weight     *string       `json:"weight"`
	BloodGroup *string       `json:"bloodGroup"`
	Marriage   *string       `json:"marriage"`
	WebPage    *string       `json:"webpage"`
	Picture    *Picture      `json:"picture"`
	Pictures   []PictureInfo `json:"pictures"`
	Title      *string       `json:"title"`
	CastCode   *string       `json:"castCode"`
	Introduce  *string       `json:"introduce"`
}

type PictureInfo struct {
	REL             *string  `json:"rel"`
	HREF            string   `json:"href"`
	Description     *string  `json:"description"`
	ImageType       *string  `json:"imageType"`
	CopyRightNotice *string  `json:"copyrightNotice"`
	MIMType         *string  `json:"mimType"`
	Resolution      []string `json:"resolution"`
}

type Prop struct {
	Name     *string `json:"name"`
	Included *string `json:"included"`
	Excluded *string `json:"excluded"`
}
