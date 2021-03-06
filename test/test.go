package test

import (
	"github.com/sangshenya/go-channel/util"
	"time"
)

type adInfo struct {
	Title string
	Content string
	ImageUrl string
}
// "",""
func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	reqFunc(getReq)
	adinfo := adInfo{
		Title:    "大牌好货,每满300减40立即前往",
		Content:  "每满300减40立即前往，普通测试广告",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_1.png",
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    adinfo.Title,
		Content:  adinfo.Content,
		ImageUrl: adinfo.ImageUrl,
		Uri:      "https://pages.tmall.com/wow/a/act/tmall/dailygroup/1773/wupr?wh_pid=daily-218058&activity_id=100000000145",
	}

	switch getReq.ChannelReq.Adtype {
		case "flow":
			resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/tbdhh_fz.jpg"
		case "banner":
			resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/71911590995175_.pic_hd.jpg"
		case "noticead":
			resultData.Title = "京东主页，浮窗测试广告标题"
			resultData.Content = "京东主页，浮窗测试广告描述"
			resultData.ImageUrl = "https://img.admobile.top/admobile-adRequest/jd_icon.png"
			resultData.Uri = "https://www.jd.com"
			resultData.Scheme = "openapp.jdmobile://"
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

func TestVideoFlow(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	reqFunc(getReq)
	if getReq.ChannelReq.Adtype != "flow" {
		getReq.ChannelReq.Errorinfo = "不支持flow以外的类型请求"
		channelError := util.NewChannelRequestFailErrorWithText("不支持flow以外的类型请求")
		return util.ResMsg{}, channelError
	}
	resultData, err := Base(getReq, reqFunc)
	if err != nil {
		return util.ResMsg{}, err
	}
	resultData.VideoUrl = "http://video-ad.sm.cn/38560410a55748f58356284317a86208/5146645ddc904cf18906e24384e34c78-4d7cdb48cf1619d99b743366e70333b4-ld.mp4"

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return resultData, nil
}