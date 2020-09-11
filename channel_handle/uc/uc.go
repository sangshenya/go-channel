package uc

import (
	"bytes"
	"encoding/json"
	"go-channel/channel_handle/util"
	"io/ioutil"
	"net/http"
	"strings"
)

const(
	URL = "http://huichuan.sm.cn/nativead"
	URL_TEST = "https://test.huichuan.sm.cn/nativead"
)

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	udid := getReq.Idfa
	adt := "ios"
	if getReq.Os != "2" {
		adt = "android"
		if len(getReq.Imei) == 0 {
			udid = getReq.Androidid
		} else {
			udid = getReq.Imei
		}
	}

	carrier := "Unknown"
	if getReq.Imsi == "1" {
		carrier = "ChinaMobile"
	} else if getReq.Imsi == "2" {
		carrier = "ChinaUnicom"
	} else if getReq.Imsi == "3" {
		carrier = "ChinaTelecom"
	}

	fr := "android"
	if getReq.Os == "2" {
		fr = "iphone"
	}

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)

	adid := getReq.ChannelReq.Adid
	wid := getReq.ChannelReq.Appid

	pkg := getReq.ChannelReq.Pkg
	appname := getReq.ChannelReq.Appname
	w, _ := paramsMap["w"]
	h, _ := paramsMap["h"]

	if len(adid) == 0 || len(wid) == 0 {
		failFunc(getReq)
		return util.ResMsg{}
	}

	postData := adreq{
		Ad_device_info: _ad_device_info{
			Android_id: getReq.Androidid,
			Devid:      udid,
			Imei:       getReq.Imei,
			Oaid:       getReq.Oaid,
			Udid:       "",
			Open_udid:  getReq.Openudid,
			Idfa:       getReq.Idfa,
			Device:     getReq.Model,
			Os:         adt,
			Osv:        getReq.Osversion,
			Mac:        getReq.Mac,
			Sw:         getReq.Screenwidth,
			Sh:         getReq.Screenheight,
			Is_jb:      "2",
			Access:     strings.ToUpper(getReq.Network),
			Carrier:    carrier,
			Cp:         "",
			Aid:        "",
			Client_ip:  getReq.Ip,
		},
		Ad_app_info:    _ad_app_info{
			Fr:       fr,
			Dn:       "",
			Sn:       "",
			Utdid:    "",
			Is_ssl:   "1",
			Pkg_name: pkg,
			Pkg_ver:  getReq.Appversion,
			App_name: appname,
			Ua:       getReq.Ua,
		},
		Ad_gps_info:    _ad_gps_info{
			Lng: getReq.Lng,
			Lat: getReq.Lat,
		},
		Ad_pos_info:    []_ad_pos_info{
			{
				Slot_type: "0",
				Slot_id:   adid,
				Ad_style:  []string{},
				Req_cnt:   "1",
				Wid:       wid,
				Aw:        w,
				Ah:        h,
			},
		},
	}

	ma, err := json.Marshal(&postData)
	if err != nil {
		failFunc(getReq)
		return util.ResMsg{}
	}

	byteData := []byte{
		0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
	}
	//fmt.Println("hh",string(ma))

	req, err := http.NewRequest("POST", URL, bytes.NewReader(append(byteData, ma...)))
	if err != nil {
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

	resData := adres{}
	json.Unmarshal(data[16:], &resData)

	if len(resData.Slot_ad) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	slot_ad := resData.Slot_ad[0]

	if len(slot_ad.Ad) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	ad := slot_ad.Ad[0]

	if len(ad.Ad_content.Img_1) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	if len(ad.Turl) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	resultData := util.ResMsg{
		Id:       "0",
		Weight:   0,
		State:    0,
		Title:    ad.Ad_content.Title,
		Content:  ad.Ad_content.Description,
		ImageUrl: ad.Ad_content.Img_1,
		Uri:      replace(ad.Turl[0], false),
	}

	if len(ad.Turl) >= 2 {
		resultData.Uri = replace(ad.Turl[1], false)
	}

	for _, item := range ad.Vurl {
		resultData.Displayreport = append(resultData.Displayreport, replace(item, true))
	}

	for _, item := range ad.Curl {
		resultData.Clickreport = append(resultData.Clickreport, replace(item, false))
	}

	for _, item := range ad.Turl {
		resultData.Clickreport = append(resultData.Clickreport, replace(item, false))
	}

	resultData.StartDownload = append(resultData.StartDownload, ad.Eurl + "&client_event=download_begin")
	resultData.Downloaded = append(resultData.Downloaded, ad.Eurl + "&client_event=download_done")
	resultData.Installed = append(resultData.Installed, ad.Eurl + "&client_event=install_begin")

	return resultData
}

func replace(urlStr string, imp bool) string {
	timeArray := util.GetTime()

	if imp {
		urlStr = strings.Replace(urlStr, "{TS}", timeArray[1], -1)
	} else {
		urlStr = strings.Replace(urlStr, "{TS}", timeArray[2], -1)
	}

	return urlStr
}