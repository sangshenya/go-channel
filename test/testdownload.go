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
	reqFunc(getReq)
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


	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		getReq.ChannelReq.Errorinfo = "数据不完整"
		noFunc(getReq)
		return util.ResMsg{}
	}

	return resultData
}

func GDTDownloadBase(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	reqFunc(getReq)
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
		Uri:      "http://101.37.79.7:8088/jump",
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

	resultData.Json = true

	resultData.StartDownload = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=5&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Downloaded = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=6&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Installed = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=7&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Open = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y"}

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		getReq.ChannelReq.Errorinfo = "数据不完整"
		noFunc(getReq)
		return util.ResMsg{}
	}
	return resultData
}