package dongqiudi

type adres struct {
	Id string
	BidId string
	Seats []ResSeats

}

type ResSeats struct {
	Bids []ResBid
	Group int
	Seat string
}

type ResBid struct {
	Id string
	ImpId string
	DealId string
	CrId string
	Cid string
	AdId string
	Domains []string
	Cat []string
	Price float64
	NUrl string
	BUrl string
	LUrl string
	Target string
	DeepLink string
	ActionType string
	Demand string
	Bundle string
	Adm string
	App ResApp
	Banner ResBanner
	Native ResNative
	Events ResEvent
}

type ResApp struct {
	Size int
	Md5 string
	Icon string
	Pack string
	Vers string
	Name string

}

type ResBanner struct {
	W int
	H int
	Url string
	Mimes string
	Skip int
	Duration int
	VideoUrl string
	SkipMinTime int
}

type ResNative struct {
	Version string
	Assets []ResAssets
}

type ResAssets struct {
	Id int
	Title ResTitle
	Data ResData
	Img ResImg
}

type ResTitle struct {
	Len int
	Text string

}

type ResImg struct {
	W int
	H int
	Url string
	Mimes string
}

type ResData struct {
	Type int
	Len int
	Value string
}

type ResEvent struct {
	Els []string
	Cls []string
	Sdls []string
	Edls []string
	Sils []string
	Eils []string
	Ials []string
	Dcls []string

}