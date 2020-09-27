package channel_handle

import (
	"go-channel/channel_handle/inmobi"
	"go-channel/channel_handle/shjy"
	"go-channel/channel_handle/uc"
	"go-channel/channel_handle/util"
	"go-channel/channel_handle/wuque"
	"go-channel/channel_handle/ymtb"
)

var(

	FunMap = map[string]func(Req *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg{
		"inmobi":inmobi.Base,
		"shjy":shjy.Base,
		"uc":uc.Base,
		"wuque":wuque.Base,
		"ymtb":ymtb.Base,
	}

)

func RequestChannel(channelName string, getReq *util.ReqMsg, channelErrorFunc util.ReqFailFunc, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	resultData := util.ResMsg{}
	if funName, ok := FunMap[channelName]; ok {
		resultData = funName(getReq, failFunc, reqFunc, noFunc, timeoutFunc, noimgFunc, nourlFunc)
	} else {
		channelErrorFunc(getReq)
	}
	return resultData
}