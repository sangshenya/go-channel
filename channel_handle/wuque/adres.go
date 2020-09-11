package wuque

type adres struct {
	Error_code int
	Request_id string
	Wxad       Ad
}

type Ad struct {
	Title                	string
	Ad_title             	string
	Brand_name           	string
	Description          	string
	Image_src            	string
	Icon_src             	string
	Creative_type        	int
	Interaction_type     	int
	App_package          	string
	Win_notice_url       	[]string
	Click_url           	[]string
	Download_track_urls  	[]string
	Downloaded_track_urls 	[]string
	Installed_track_urls    []string
	Open_track_urls  		[]string
	Action_track_urls 		[]string
	Dp_success_track_urls 	[]string
	Landing_page_url        string
	Deep_link         		string
}