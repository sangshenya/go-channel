package dongqiudi

type adreq struct {
	Id string `json:"id"`
	Version string `json:"version"`
	Imps []_imp `json:"imps"`
	App _app `json:"app"`
	Device _device `json:"device"`
	At int `json:"at"`
	Test int `json:"test"`
	TMax int `json:"tMax"`
	Ext _ext `json:"ext"`
	Language string `json:"language"`
}

type _imp struct {
	Id string `json:"id"`
	Aw int `json:"aw"`
	Ah int `json:"ah"`
	TagId string `json:"tagId"`
	BidFloor float64 `json:"bidFloor"`
	Banner *_banner `json:"banner,omitempty"`
	Native *_native `json:"native,omitempty"`
	Mts []string `json:"mts"`

}

type _banner struct {
	W int `json:"w"`
	H int `json:"h"`
	Pos int `json:"pos"`
	Type int `json:"type"`
	Mimes []string `json:"mimes"`

}

type _native struct {
	Version string `json:"version"`
	Assets []_assets `json:"assets"`

}

type _assets struct {
	Id int `json:"id"`
	Title *_title `json:"title,omitempty"`
	Data *_data `json:"data,omitempty"`
	Img *_img `json:"img,omitempty"`
	Required int `json:"required"`

}

type _title struct {
	Len int `json:"len"`

}

type _data struct {
	Type int `json:"type"`
	Len int `json:"len"`
}

type _img struct {
	W int `json:"w"`
	H int `json:"h"`
	Pos int `json:"pos"`
	Type int `json:"type"`
	Mimes []string `json:"mimes"`

}

type _app struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Bundle string `json:"bundle"`
	Ver string `json:"ver"`
	Paid int `json:"paid"`
	Keywords string `json:"keywords"`
	Cat []string `json:"cat"`

}

type _device struct {
	Ua string `json:"ua"`
	Geo _geo `json:"geo"`
	Ip string `json:"ip"`
	Ipv6 string `json:"ipv6"`
	DeviceType int `json:"deviceType"`
	Make string `json:"make"`
	Model string `json:"model"`
	Os string `json:"os"`
	Osv string `json:"osv"`
	Rvs string `json:"rvs"`
	Sct string `json:"sct"`
	Anal string `json:"anal"`
	Hwv string `json:"hwv"`
	H int `json:"h"`
	W int `json:"w"`
	Sw int `json:"sw"`
	Sh int `json:"sh"`
	Ppi int `json:"ppi"`
	Density float64 `json:"density"`
	Dpi int `json:"dpi"`
	Ifa string `json:"ifa"`
	Ifv string `json:"ifv"`
	Did string `json:"did"`
	Dpid string `json:"dpid"`
	Oaid string `json:"oaid"`
	Mac string `json:"mac"`
	Carrier string `json:"carrier"`
	ConnectionType int `json:"connectionType"`
	Ibis string `json:"ibis"`
	Orientation int `json:"orientation"`
	Language string `json:"language"`

}

type _geo struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type _ext struct {
	Rdt int `json:"rdt"`
	Https int `json:"https"`
	DeepLink int `json:"deepLink"`
	Download int `json:"download"`
	Admt int `json:"admt"`
	Vech int `json:"vech"`
	Vecv int `json:"vecv"`
}