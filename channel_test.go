package main

import (
	"encoding/json"
	"fmt"
	"github.com/sangshenya/go-channel/channel_handle"
	"github.com/sangshenya/go-channel/util"
	"testing"
)

func TestChannel(t *testing.T)  {
	requestData := []byte(`{"appid":"2593686","package":"me.zhouzhuo810.zznote","ts":"1608602307856","sign":"33f0c32150e33bf10af4c1ae6130bfdf","adtype":"vod","channel":"","ts_hour":"2020122211","screenwidth":"","screenheight":"","os":"1","osversion":"8.1.0","imsi":"460070346411979","ip":"10.1.1.195","mac":"04:B1:67:AA:DD:0E","network":"WIFI","device":"","androidid":"a1fb2bcf2af0f0bf","idfa":"","idfv":"","openudid":"","appversion":"3.2.9","terminal":"","ua":"PostmanRuntime/7.6.0","version":"3.2.9","sd":"480","machine":"4b2549ebe2325061cbe575b9349b689a","vendor":"Xiaomi","imei":"866655036743262","model":"MI 5X","id":"","width":"1080","height":"1920","lat":"30.252501","lng":"120.165024","sdkid":"","orientation":"0","imsi_long":"","AndroidApiLevel":"","carrier":"","admobile":"","oaid":"cfa94a0b78f5759b","adwidth":"","adheight":"","machinedmp":"530940603958034432","vodsubtype":"","sdkVersion":"4.8.2","posid":"1a4015ce9a9035c0f3","uniqueid":"2160","wifiname":"","wifimac":"","romversion":"V11","comptime":"1570642365000"}`)
	getReq := util.ReqMsg{}
	error := json.Unmarshal(requestData, &getReq)
	if error != nil {
		fmt.Println("error1:", error)
		return
	}

	getReq.Adtype = "flow"
	requestChannelName := "jdlm"
	getReq.ChannelReq.Adid = "3003263355"
	getReq.ChannelReq.Appid = "4100249275"
	getReq.ChannelReq.Token = "b3350050d63d4c438cb1e0728221cd2d"
	getReq.ChannelReq.Params = "appkey=48dd082fdb5841c95409e52c1b9db083"
	getReq.ChannelReq.Adtype = "flow"
	getReq.ChannelReq.Appname = ""
	getReq.ChannelReq.Appurl = ""
	getReq.ChannelReq.Pkg = ""
	getReq.ChannelReq.Loginfo = "--test--"

	resultData := channel_handle.RequestChannel(requestChannelName, &getReq, channelErrorSubFunc, failSubFunc, reqSubFunc, requestnoSubFunc, timeoutSubFunc, noimgSubFunc, nourlSubFunc)
	ma, error := json.Marshal(resultData)
	fmt.Println(string(ma), error)
}

func channelErrorSubFunc(reqMsg *util.ReqMsg)  {
	fmt.Println("channelErrorFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func failSubFunc(reqMsg *util.ReqMsg)  {
	fmt.Println("failFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func reqSubFunc(reqMsg *util.ReqMsg) {
	fmt.Println("reqSubFunc:", reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func requestnoSubFunc(reqMsg *util.ReqMsg) {
	fmt.Println("requestnoFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func timeoutSubFunc(reqMsg *util.ReqMsg)  {
	fmt.Println("timeoutFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func noimgSubFunc(reqMsg *util.ReqMsg)  {
	fmt.Println("noimgFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}

func nourlSubFunc(reqMsg *util.ReqMsg)  {
	fmt.Println("nourlFunc:",reqMsg.ChannelReq.Errorinfo, reqMsg.ChannelReq.Loginfo)
}