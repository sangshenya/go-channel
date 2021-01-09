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

func DownloadBase(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
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
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(resultData.Uri) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}


	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return resultData, nil
}

func GDTDownloadBase(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	reqFunc(getReq)
	if getReq.Os == "2" {
		channelError := util.NewChannelRequestNoErrorWithText("该渠道只支持安卓端")
		return util.ResMsg{}, channelError
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
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(resultData.Uri) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	resultData.Json = true

	resultData.StartDownload = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=5&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Downloaded = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=6&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Installed = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y", "http://t.gdt.qq.com/conv/alliance/api/conv?client=6&action_id=7&click_id=__CLICK_ID__&product_id=100812722"}
	resultData.Open = []string{"https://www.baidu.com?IT_CLK_PNT_DOWN_XaaaIT_CLK_PNT_DOWN_YaaaIT_CLK_PNT_UP_XaaaIT_CLK_PNT_UP_Y"}

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}
	return resultData, nil
}