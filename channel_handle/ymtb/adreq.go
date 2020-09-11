package ymtb

type adreq struct {
	Version string `json:"version"`
	Id string `json:"id"`
	Pid string `json:"pid"`
	Channel_id string `json:"channel_id"`
	User _user `json:"user"`
	Device _device `json:"device"`
}

type _user struct {
	Uid string `json:"uid"`
	Uid_type string `json:"uid_type"`
	Uid_encryption string `json:"uid_encryption"`
}

type _device struct {
	Ipv4 string `json:"ipv4"`
	Device_type string `json:"device_type"`
	Device_make string `json:"device_make"`
	Device_os string `json:"device_os"`
	Network string `json:"network"`
}