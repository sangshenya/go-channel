package ymtb

type adres struct {
	Status int
	Creative creative
}

type creative struct {
	Type int
	Imp_url string
	Clk_url string
	Creative_id string
	Deeplink_url string
	Title string
	Img string
	Imgs []string

}