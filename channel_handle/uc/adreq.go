package uc

type adreq struct {
	Ad_device_info _ad_device_info `json:"ad_device_info"`
	Ad_app_info _ad_app_info `json:"ad_app_info"`
	Ad_gps_info _ad_gps_info `json:"ad_gps_info"`
	Ad_pos_info []_ad_pos_info `json:"ad_pos_info"`
}

type _ad_device_info struct {
	Android_id string `json:"android_id"`
	Devid string `json:"devid"`
	Imei string `json:"imei"`
	Oaid string `json:"oaid"`
	Udid string `json:"udid"`
	Open_udid string `json:"open_udid"`
	Idfa string `json:"idfa"`
	Device string `json:"device"`
	Os string `json:"os"`
	Osv string `json:"osv"`
	Mac string `json:"mac"`
	Sw string `json:"sw"`
	Sh string `json:"sh"`
	Is_jb string `json:"is_jb"`
	Access string `json:"access"`
	Carrier string `json:"carrier"`
	Cp string `json:"cp"`
	Aid string `json:"aid"`
	Client_ip string `json:"client_ip"`
}

type _ad_app_info struct {
	Fr string `json:"fr"`
	Dn string `json:"dn"`
	Sn string `json:"sn"`
	Utdid string `json:"utdid"`
	Is_ssl string `json:"is_ssl"`
	Pkg_name string `json:"pkg_name"`
	Pkg_ver string `json:"pkg_ver"`
	App_name string `json:"app_name"`
	Ua string `json:"ua"`

}

type _ad_gps_info struct {
	Lng string `json:"lng"`
	Lat string `json:"lat"`

}

type _ad_pos_info struct {
	Slot_type string `json:"slot_type"`
	Slot_id string `json:"slot_id"`
	Ad_style []string `json:"ad_style"`
	Req_cnt string `json:"req_cnt"`
	Wid string `json:"wid"`
	Aw string `json:"aw"`
	Ah string `json:"ah"`

}

type _page_info struct {

}

type _res_info struct {

}

type _ext_info struct {

}