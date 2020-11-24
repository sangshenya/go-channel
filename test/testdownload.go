package test

import (
	"github.com/sangshenya/go-channel/util"
	"time"
)

var adinfo_download = adInfo{
	Title:    "下载测试",
	Content:  "网上厨房-好用的美食菜谱",
	ImageUrl: "https://pic.ecook.cn/ad_test_S.jpg",
}

func DownloadBase(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    adinfo_download.Title,
		Content:  adinfo_download.Content,
		ImageUrl: adinfo_download.ImageUrl,
		Uri:      "https://at.umeng.com/yyWDiy?cid=4944",
	}

	if getReq.Os == "2" {
		resultData.Uri = "https://itunes.apple.com/cn/app/id455221304"
	}

	switch getReq.ChannelReq.Adtype {
	case "flow":
		resultData.ImageUrl = "https://pic.ecook.cn/test_flow.jpg"
	case "banner":
		resultData.ImageUrl = "https://pic.ecook.cn/test_banner.jpg"
	}

	if len(resultData.ImageUrl) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	if len(resultData.Uri) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}

	return resultData
}

func GDTDownloadBase(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	if getReq.Os == "2" {
		noFunc(getReq)
		return util.ResMsg{}
	}
	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    "广点通下载测试",
		Content:  adinfo_download.Content,
		ImageUrl: adinfo_download.ImageUrl,
		Uri:      "http://c.gdt.qq.com/gdt_mclick.fcg?viewid=S4iZOkpHQgRseIqymvFIAbdFlZTY2YWMYBCSeWEIcQtKdF67MRxwyaw2b9CcxCwGqLyVjYy3MAnTmheaSY0xWdraZpqhZUqENErv81VxkFQE_bcbjE2lJTfwPEkPoUSb9bU8RID9BI4QFMopEgs!2krkZR9qo7q!0lLcQ9pKBvLPQQQ!zIMwTnnHc8!m8a_t1wv1VEZBkk4&jtype=0&i=1&os=2&asi=%7B%22mf%22%3A%22%E5%BA%B7%E4%BD%B3%22%7D&acttype=1&s=%7B%22req_width%22%3A%22300%22%2C%22req_height%22%3A%22250%22%2C%22width%22%3A%22300%22%2C%22height%22%3A%22250%22%2C%22down_x%22%3A%22__DOWN_X__%22%2C%22down_y%22%3A%22__DOWN_Y__%22%2C%22up_x%22%3A%22__UP_X__%22%2C%22up_y%22%3A%22__UP_Y__%22%7DIT_CLK_PNT_DOWN_X",
	}

	switch getReq.ChannelReq.Adtype {
	case "flow":
		resultData.ImageUrl = "https://pic.ecook.cn/test_flow.jpg"
	case "banner":
		resultData.ImageUrl = "https://pic.ecook.cn/test_banner.jpg"
	}

	if len(resultData.ImageUrl) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	if len(resultData.Uri) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}

	resultData.StartDownload = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=5&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Downloaded = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=6&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Installed = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=7&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Open = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y"}

	return resultData
}