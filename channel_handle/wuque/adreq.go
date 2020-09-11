package wuque

type adreq struct {
	Sign     string `json:"sign"`
	Apptoken string `json:"apptoken"`
	Data     *Data  `json:"data"`
}

type Data struct {
	App     *App     `json:"app"`
	Adslot  *Adslot  `json:"adslot"`
	Device  *Device  `json:"device"`
	Network *Network `json:"network"`
	Gps     *Gps     `json:"gps,omitempty"`
}

type App struct {
	App_version *Version `json:"app_version"`
	PkgName string `json:"pkgName"`
	AppName string `json:"appName"`
}

type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Micro int `json:"micro"`
}

type Adslot struct {
	Adslot_id   		string `json:"adslot_id"`
	Adslot_size 		*Size  `json:"adslot_size"`
	Support_deeplink 	int	   `json:"support_deeplink"`
	Secure 				int	   `json:"secure"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Device struct {
	Device_type 	int   		`json:"device_type"`
	Os_type     	int   		`json:"os_type"`
	Os_version  	*Version 	`json:"os_version"`
	Vendor      	string   	`json:"vendor"`
	Manufacturer 	string   	`json:"manufacturer"`
	Model       	string   	`json:"model"`
	Screen_size 	*Size    	`json:"screen_size"`
	Udid        	*Udid    	`json:"udid"`
	User_agent		string 		`json:"user_agent"`
	Orientation 	int 		`json:"orientation"`
}

type Udid struct {
	Idfa       	string `json:"idfa"`
	Idfv 		string `json:"idfv"`
	Imei       	string `json:"imei"`
	Android_id 	string `json:"android_id"`
	Mac        	string `json:"mac"`
	oaid 		string `json:"oaid"`
}

type Network struct {
	Ipv4            string 	`json:"ipv4"`
	Connection_type int 	`json:"connection_type"`
	Operator_type   int 	`json:"operator_type"`
}

type Gps struct {
	Coordinate_type int 	`json:"coordinate_type"`
	Longitude       string 	`json:"longitude"`
	Latitude        string 	`json:"latitude"`
	Timestamp       string 	`json:"timestamp"`
}