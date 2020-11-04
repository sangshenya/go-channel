package channel_handle

import (
	"github.com/sangshenya/go-channel/inmobi"
	"github.com/sangshenya/go-channel/oneway"
	"github.com/sangshenya/go-channel/qrqm"
	"github.com/sangshenya/go-channel/shjy"
	"github.com/sangshenya/go-channel/test"
	"github.com/sangshenya/go-channel/uc"
	"github.com/sangshenya/go-channel/util"
	"github.com/sangshenya/go-channel/wuque"
	"github.com/sangshenya/go-channel/ymtb"
)

/*
渠道的宏替换
	__TS__:当前时间,单位:毫秒
	__TS_S__:当前时间，单位秒

	__DOWN_X__:相对于广告位的按下x坐标
	__DOWN_Y__:相对于广告位的按下y坐标
	__UP_X__:相对于广告位的抬起x坐标
	__UP_Y__:相对于广告位的抬起y坐标

	__RE_DOWN_X__:相对于屏幕的按下x坐标
	__RE_DOWN_Y__:相对于屏幕的按下y坐标
	__RE_UP_X__:相对于屏幕的抬起x坐标
	__RE_UP_Y__:相对于屏幕的抬起y坐标

	__WIDTH__:在手机上真实展示的宽度，与手机屏幕宽度相关
	__HEIGHT__:在手机上真实展示的高度，与手机屏幕宽度、广告类型相关

	__CLICK_ID__:广点通下载id

	其中请求宽高应该当在请求是进行替换

*/

var(

	FunMap = map[string]func(Req *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg{
		"inmobi":inmobi.Base,
		"ymtb":ymtb.Base,
		"wuque":wuque.Base,
		"uc":uc.Base,
		"shjy":shjy.Base,
		"oneway":oneway.Base,
		"跃盟":ymtb.Base,
		"test":test.Base,
		"qrqm":qrqm.Base,
	}

)

func RequestChannel(channelName string, getReq *util.ReqMsg, channelErrorFunc util.ReqFailFunc, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	resultData := util.ResMsg{}
	if funName, ok := FunMap[channelName]; ok {
		resultData = funName(getReq, failFunc, reqFunc, noFunc, timeoutFunc, noimgFunc, nourlFunc)
	} else {
		getReq.ChannelReq.Errorinfo = "渠道号未匹配"
		channelErrorFunc(getReq)
	}
	return resultData
}
