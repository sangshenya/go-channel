package test

import (
	"github.com/sangshenya/go-channel/util"
	"math/rand"
	"time"
)

type adInfo struct {
	Title string
	Content string
	ImageUrl string
}

var pgyAdInfo = []adInfo{
	{
		Title:    "大牌好货,每满300减40立即前往",
		Content:  "每满300减40立即前往",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_1.png",
	},
	{
		Title:    "潮流新品 造型新主张",
		Content:  "潮流新品 造型新主张",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_2.png",
	},
	{
		Title:    "双11狂欢季",
		Content:  "红包多多 美味多多！",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_3.png",
	},
}

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {

	adinfo := pgyAdInfo[rand.Intn(len(pgyAdInfo))]

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    adinfo.Title,
		Content:  adinfo.Content,
		ImageUrl: adinfo.ImageUrl,
		Uri:      "https://s.click.taobao.com/1P4Jjuu",
		Scheme:   "taobao://s.click.taobao.com/1P4Jjuu",
	}

	switch getReq.ChannelReq.Adtype {
		case "flow":
			resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/tbdhh_fz.jpg"
		case "banner":
			resultData.ImageUrl = "https://admobile.oss-cn-hangzhou.aliyuncs.com/admobile-adRequest/71911590995175_.pic_hd.jpg"
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