package oneway

type adreq struct {
	ApiVersion      string  `json:"apiVersion"`
	PlacementId     string  `json:"placementId"`
	BundleId        string  `json:"bundleId"`
	BundleVersion   string  `json:"bundleVersion"`
	AppName			string 	`json:"appName"`
	SubAffId		string 	`json:"subAffId"`
	DeviceId 		string  `json:"deviceId"`
	Imei 			string  `json:"imei"`
	AndroidId 		string  `json:"androidId"`
	Oaid			string 	`json:"oaid"`
	ApiLevel 		int 	`json:"apiLevel"`
	Os  			int 	`json:"os"`
	OsVersion 		string  `json:"osVersion"`
	ConnectionType 	string  `json:"connectionType"`
	NetworkType 	int 	`json:"networkType"`
	NetworkOperator string  `json:"networkOperator"`
	SimOperator 	string  `json:"simOperator"`
	Imsi 			string  `json:"imsi"`
	DeviceMake 		string  `json:"deviceMake"`
	DeviceModel 	string  `json:"deviceModel"`
	DeviceType		int 	`json:"deviceType"`
	Orientation		string 	`json:"orientation"`
	Mac 			string  `json:"mac"`
	WifiBSSID		string 	`json:"wifiBSSID"`
	WifiSSID		string 	`json:"wifiSSID"`
	ScreenWidth 	int  	`json:"screenWidth"`
	ScreenHeight 	int  	`json:"screenHeight"`
	ScreenDensity 	int 	`json:"screenDensity"`
	UserAgent 		string  `json:"userAgent"`
	Ip 				string  `json:"ip"`
	Language 		string  `json:"language"`
	TimeZone 		string  `json:"timeZone"`
	Longitude		float64 `json:"longitude"`
	Latitude		float64 `json:"latitude"`

}

type trackreq struct {
	EventName      string `json:"eventName"`
	SessionId      string `json:"sessionId"`
	EventId        string `json:"eventId"`
	DeviceId       string `json:"deviceId"`
	Imei           string `json:"imei"`
	AndroidId      string `json:"androidId"`
	UserAgent      string `json:"userAgent"`
	OsVersion      string `json:"osVersion"`
	ApiLevel       int    `json:"apiLevel"`
	BundleId       string `json:"bundleId"`
	ConnectionType string `json:"connectionType"`
	DeviceMake     string `json:"deviceMake"`
	DeviceModel    string `json:"deviceModel"`
	Language       string `json:"language"`
	NetworkType    int    `json:"networkType"`
	TimeZone       string `json:"timeZone"`
	CampaignId     int    `json:"campaignId"`
	Mac            string `json:"mac"`
	WifiBSSID      string `json:"wifiBSSID"`
	WifiSSID       string `json:"wifiSSID"`
}
