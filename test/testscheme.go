package test

import (
	"github.com/sangshenya/go-channel/util"
	"time"
)

func SchemeBase(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	reqFunc(getReq)
	adinfo := adInfo{
		Title:    "scheme测试",
		Content:  "潮流新品 造型新主张",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_2.png",
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    adinfo.Title,
		Content:  adinfo.Content,
		ImageUrl: adinfo.ImageUrl,
		Uri:      "https://s.click.taobao.com/GEst4hv",
		Scheme:   "taobao://s.click.taobao.com/GEst4hv",
	}

	switch getReq.ChannelReq.Adtype {
	case "flow":
		resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/tbcjhb_f2.jpg"
	case "banner":
		resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/71911590995175_.pic_hd.jpg"
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