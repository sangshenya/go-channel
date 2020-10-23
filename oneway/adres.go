package oneway

type adres struct {
	Success bool
	Message string
	ErrorCode int
	Data    ad
}

type ad struct {
	AppName      	string
	AppStoreId   	string
	Title        	string
	ClickUrl     	string
	Deeplink	 	string
	InteractionType int
	LandingPageUrl  string
	TrackingEvents  _trackingEvent
	Images 			[]_image
}

type _trackingEvent struct {
	Show 				[]string
	Click 				[]string
	Dp 					[]string
	ApkDownloadStart 	[]string
	ApkDownloadFinish 	[]string
	PackageAdded 		[]string
	ActiveOpen			[]string

}

type _image struct {
	Url string
}
