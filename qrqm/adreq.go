package qrqm

type tbreq struct {
	Id string `json:"id"`
	Version int `json:"version"`
	Device tbdevice `json:"device"`
	Adzone_id string `json:"adzone_id"`
}

type tbdevice struct {
	Osv string `json:"osv"`
	Os string `json:"os"`
	Ip string `json:"ip"`
	Idfa string `json:"idfa"`
	Imei string `json:"imei"`
	Oaid string `json:"oaid"`
	Mac string `json:"mac"`
	Network int `json:"network"`
	Imei_md5 string `json:"imei_md5"`
	Device_type int `json:"device_type"`
}