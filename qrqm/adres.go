package qrqm

type tbadres struct {
	Tbk_thor_creative_launch_response _tbk_thor_creative_launch_response
	Error_response _error_response
}

type _error_response struct {
	Code int
	Msg string
	Sub_code string
	Request_id string

}

type _tbk_thor_creative_launch_response struct {
	Result _tbk_thor_creative_launch_response_result
}

type _tbk_thor_creative_launch_response_result struct {
	Click_through_url string
	Image_url4 string
	Impression_tracking_url string
	Deeplink_url string
	Pid string
	Title string
	Creative_id string
	Creative_type int
	Image_url2 string
	Image_url string
	Image_url3 string
	Id string
	Desc string
	Status int
}