package inmobi

type adreq struct {
	App _app `json:"app"`
	Imp _imp `json:"imp"`
	Device _device `json:"device"`
	Ext _ext `json:"ext"`

}

type _app struct {
	Id string `json:"id"`
	Bundle string `json:"bundle"`
	Ver string `json:"ver"`
	Orientation int `json:"orientation"`
	Paid int `json:"paid"`
}

type _device struct {
	Ifa string `json:"ifa,omitempty"`
	Oaid string `json:"oaid,omitempty"`
	Gpid string `json:"gpid,omitempty"`
	Md5_imei string `json:"md5_imei,omitempty"`
	Sha1_imei string `json:"sha1_imei,omitempty"`
	O1 string `json:"o1,omitempty"`
	Um5 string `json:"um5,omitempty"`
	Type string `json:"type"`
	Ua string `json:"ua"`
	Ip string `json:"ip,omitempty"`
	Os string `json:"os"`
	Osv string `json:"osv"`
	Model string `json:"model"`
	Connectiontype int `json:"connectiontype,omitempty"`
	Language string `json:"language"`
	Orientation int `json:"orientation"`
	Carrier string `json:"carrier,omitempty"`
	Geo _geo `json:"geo,omitempty"`
	Ext _deviceExt `json:"ext"`
	Idfv string `json:"idfv"`
	Make string `json:"make"`
	H int `json:"h"`
	W int `json:"w"`
	Ppi int `json:"ppi"`


}

type _deviceExt struct {
	Orientation int `json:"orientation,omitempty"`
}

type _geo struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
	Accu int `json:"accu"`
}

type _imp struct {
	Secure int `json:"secure,omitempty"`
	Trackertype string `json:"trackertype,omitempty"`
	Native *_native `json:"native,omitempty"`
	Ext _impext `json:"ext"`
	Banner *_banner `json:"banner,omitempty"`
}

type _impext struct {
	AdsCount int `json:"adsCount"`
}

type _native struct {
	Layout int `json:"layout"`
	ScreenWidth int `json:"screenWidth,omitempty"`
}

type _banner struct {
	H int `json:"h,omitempty"`
	W int `json:"w,omitempty"`

}

type _ext struct {
	Responseformat string `json:"responseformat"`
	SupportDeeplink bool `json:"SupportDeeplink"`
}
