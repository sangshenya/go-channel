package test

import (
	"github.com/sangshenya/go-channel/util"
	"time"
)

// "",""
func Baseflowvod(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	if getReq.ChannelReq.Adtype != "flow" {
		getReq.ChannelReq.Errorinfo = "只支持flow类型"
		failFunc(getReq)
		return util.ResMsg{}
	}
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
		resultData.Videourl = "https://video-ecook.oss-cn-hangzhou.aliyuncs.com/apple_ad_60s.mp4"
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
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	return resultData
}