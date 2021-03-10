package tbrta


type newAdres struct {
	Usergrowth_dhh_delivery_ask_response _newres
}

type _newres struct {
	Result bool
}

type adres struct {
	Request_id string
	Error_response string
	Code string
	Msg string
	Usergrowth_delivery_ask_response _response

}

type _response struct {
	Result bool
	Type string
	ErrMsg string
	Data []_data
}

type _data struct {

	H5_url string
	PicUrl string
	Coupon string
	Price string
	Deeplink_url string
	Title string
	TraceId string
	Smart_bid string

}