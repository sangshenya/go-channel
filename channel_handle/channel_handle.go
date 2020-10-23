package channel_handle

import (
	"github.com/sangshenya/go-channel/inmobi"
	"github.com/sangshenya/go-channel/shjy"
	"github.com/sangshenya/go-channel/uc"
	"github.com/sangshenya/go-channel/util"
	"github.com/sangshenya/go-channel/wuque"
	"github.com/sangshenya/go-channel/ymtb"
)

var(

	FunMap = map[string]func(Req *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg{
		"inmobi":inmobi.Base,
		"ymtb":ymtb.Base,
		"wuque":wuque.Base,
		"uc":uc.Base,
		"shjy":shjy.Base,
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