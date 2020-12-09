package shjy

import (
	"bytes"
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const(
	URL = "http://qiling.goryun.com/adx/v1/info"
)

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {

	width, err := strconv.Atoi(getReq.Screenwidth)
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(getReq.Screenheight)
	if err != nil {
		height = 0
	}
	network := 0
	switch getReq.Network {
	case "wifi":
		network = 2
	case "2g":
		network = 4
	case "3g":
		network = 5
	case "4g":
		network = 6
	default:
		network = 0
	}

	os := "iOS"
	if getReq.Os == "1" {
		os = "Android"
	}

	carrier, error := strconv.Atoi(getReq.Carrier)
	if error != nil {
		carrier = 0
	}

	sd, err := strconv.Atoi(getReq.Sd)
	//ppi := float64(sd) / 160

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)


	appid := getReq.ChannelReq.Appid
	adid := getReq.ChannelReq.Adid
	adtype := getReq.ChannelReq.Adtype

	pkg := getReq.ChannelReq.Pkg
	appname := getReq.ChannelReq.Appname
	w, _ := paramsMap["w"]
	h, _ := paramsMap["h"]

	if len(adid) == 0 || len(appid) == 0 || len(adtype) == 0 {
		getReq.ChannelReq.Errorinfo = "请求必需参数部分参数为空"
		failFunc(getReq)
		return util.ResMsg{}
	}

	ad_type := 201
	if adtype == "startup" {
		ad_type = 202
	} else if adtype == "flow" {
		ad_type = 207
	}

	wint, _ := strconv.Atoi(w)
	hint, _ := strconv.Atoi(h)

	appidint, _ := strconv.Atoi(appid)
	adidint, _ := strconv.Atoi(adid)

	postData := adreq{
		Version:        "20191030",
		Dnt:            0,
		Appid:          appidint,
		Adid:           adidint,
		Ver:            getReq.Appversion,
		Storeurl:       "1",
		Pos:            1,
		Adtype:         ad_type,
		Width:          wint,
		Height:         hint,
		Bundle:         pkg,
		Appname:        appname,
		Ua:             getReq.Ua,
		Devicetype:     4,
		Os:             os,
		Osv:            getReq.Osversion,
		Carrier:        carrier,
		Connectiontype: network,
		Ip:             getReq.Ip,
		Density:        sd,
		Make:           getReq.Vendor,
		Model:          getReq.Model,
		Language:       "zh-CN",
		Js:             0,
		Oaid:           getReq.Oaid,
		Imei:           getReq.Imei,
		Idfa:           getReq.Idfa,
		Androidid:      getReq.Androidid,
		Mac:            getReq.Mac,
		Lat:            getReq.Lat,
		Lon:            getReq.Lng,
		Orientation:    0,
		Sw:             width,
		Sh:             height,
		Ishttps:        0,
	}


	ma, err := json.Marshal(&postData)
	if err != nil {
		getReq.ChannelReq.Errorinfo = err.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}

	req,err := http.NewRequest("POST", URL, bytes.NewReader(ma))
	if err != nil {
		getReq.ChannelReq.Errorinfo = err.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := util.Client.Do(req)
	reqFunc(getReq)
	if err != nil {
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		noFunc(getReq)
		return util.ResMsg{}
	}

	if resp.StatusCode != 200 {
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	resData := &adres{}
	json.Unmarshal(data, resData)

	if len(resData.Native.Assets) == 0 && len(resData.Ext.Iurl) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	title := ""
	content := ""
	imgurl := ""
	impArr := []string{}
	clkArr := []string{}
	ldp := ""
	deeplink := ""
	deeplinkArray := []string{}
	dwonPkg := ""

	eventObj := _Eventtrackers{}
	if len(resData.Native.Assets) != 0 {
		ad := resData.Native.Assets[0]
		link := resData.Native.Link
		impArr = resData.Native.Imptrackers
		clkArr = resData.Native.Link.Clicktrackers

		title = ad.Title.Text
		content = ad.Data.Value
		ldp = link.Url
		if len(ad.Img.Url) != 0 {
			imgurl = ad.Img.Url[0]
		}
		deeplink = link.Fallback
		deeplinkArray = link.Fallbacktrackers

		if len(link.Dfn) != 0 {
			ldp = link.Dfn
		}
		if len(link.Bundle) != 0 {
			dwonPkg = link.Bundle
		}

		eventObj = link.Eventtrackers
	} else if len(resData.Ext.Iurl) != 0 {
		imgurl = resData.Ext.Iurl
		deeplink = resData.Ext.Fallback
		ldp = resData.Ext.Clickurl
		impArr = resData.Ext.Imptrackers
		clkArr = resData.Ext.Clicktrackers
		deeplinkArray = resData.Ext.Fallbacktrackers
		eventObj = resData.Ext.Eventtrackers

		if len(resData.Ext.Dfn) != 0 {
			ldp = resData.Ext.Dfn
		}
		if len(resData.Ext.Bundle) != 0 {
			dwonPkg = resData.Ext.Bundle
		}
	}

	if len(imgurl) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	if len(ldp) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}


	resultData := util.ResMsg{
		Id:       util.Md5(string(data) + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    title,
		Content:  content,
		ImageUrl: imgurl,
		Uri:      ldp,
	}

	if len(dwonPkg) != 0 {
		resultData.Pkg = dwonPkg
	}

	if len(deeplink) != 0 {
		resultData.Scheme = deeplink
		resultData.Schemereport = deeplinkArray
	}

	resultData.Displayreport = impArr
	resultData.Clickreport = clkArr

	resultData.StartDownload = eventObj.Startdownload
	resultData.Downloaded = eventObj.Completedownload
	resultData.Installed = eventObj.Startinstall
	resultData.Installed = append(resultData.Installed, eventObj.Completeinstall...)

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		getReq.ChannelReq.Errorinfo = "数据不完整"
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	return resultData
}