package uc

type adres struct {
	Code string
	Sid string
	Reason string
	Slot_ad []_slot_ad
}

type _slot_ad struct {
	Slot_id string
	Ad []_ad
}

type _ad struct {
	Ad_action _ad_action
	Ad_content _ad_content
	Furl string
	Vurl []string
	Curl []string
	Turl []string
	Eurl string
}

type _ad_action struct {
	Action string
}

type _ad_content struct {
	Title string
	Description string
	Img_1 string
}