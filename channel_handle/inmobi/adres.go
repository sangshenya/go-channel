package inmobi

type adres struct {
	RequestId string
	Ads []_ad
}

type _ad struct {
	Bid int
	PubContent _pubContent
	TargetUrl string
	EventTracking _tracking
	IsApp bool
	LandingURL string
	PackageName string

	//LandingPage string
	OpenExternal bool
	//BeaconUrl string
}

type _tracking struct {
	ImpressionTrackers []string
	ClickTrackers []string
	VideoStart []string
	FirstQuartile []string
	Midpoint []string
	ThirdQuartile []string
	VideoComplete []string
	VideoAutoPlay []string
	VideoPause []string
	VideoResume []string
	AppInstalled []string
	DplAttempt []string
	DplSuccess []string
	FallbackTrackers []string
	DownloadStart []string
	DownloadFinish []string
	InstallFinish []string

}

type _pubContent struct {
	LandingURL string
	Title string
	Description string
	Icon _screenshots
	Screenshots _screenshots
}

type _screenshots struct {
	Width int
	Height int
	Url string
	AspectRatio float64
}