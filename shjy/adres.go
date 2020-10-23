package shjy

type adres struct {
	Version int
	Status int
	Adm string
	Native _native
	Ext _ext
	Gdt bool
}

type _native struct {
	Assets []_assets
	Link _link
	Imptrackers []string

}

type _ext struct {
	Iurl string
	Materialtype int
	Fallback string
	Clickurl string
	Imptrackers []string
	Clicktrackers []string
	Fallbacktrackers []string
	Action int
	Eventtrackers _Eventtrackers
	Dfn string
	Bundle string
	Title string
	Desc string

}

type _assets struct {
	Id int
	Title _title
	Img _img
	Data _data
}

type _title struct {
	Text string

}

type _img struct {
	Type int
	Url []string
}

type _data struct {
	Label string
	Value string
}

type _link struct {
	Url string
	Clicktrackers []string
	Fallback string
	Fallbacktrackers []string
	Action int
	Eventtrackers _Eventtrackers
	Dfn string
	Bundle string

}

type _Eventtrackers struct {
	Iconurl string
	Startdownload []string
	Completedownload []string
	Startinstall []string
	Completeinstall []string
}