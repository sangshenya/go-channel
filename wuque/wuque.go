package wuque

import (
	"bytes"
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const(
	URL = "https://cn.bj.adx.adwangmai.com/v1.api"
)

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {
	mac := getReq.Mac
	if len(getReq.Mac) == 0 {
		mac = "00:00:00:00:00"
	}

	sw, err := strconv.Atoi(getReq.Screenwidth)
	if err != nil {
		sw = 0
	}
	sh, err := strconv.Atoi(getReq.Screenheight)
	if err != nil {
		sh = 0
	}

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)
	key := getReq.ChannelReq.Appid
	apptoken := getReq.ChannelReq.Token
	adid := getReq.ChannelReq.Adid
	wStr, _ := paramsMap["w"]
	hStr, _ := paramsMap["h"]

	if len(key) == 0 || len(apptoken) == 0 || len(adid) == 0 {
		getReq.ChannelReq.Errorinfo = "缺少必要的请求参数"
		failFunc(getReq)
		return util.ResMsg{}
	}

	w, _ := strconv.Atoi(wStr)
	h, _ := strconv.Atoi(hStr)

	postData := adreq{
		Sign:     util.Md5(key + apptoken),
		Apptoken: apptoken,
		Data: &Data{
			App: &App{
				App_version: &Version{
					Major: 1,
					Minor: 1,
					Micro: 1,
				},
			},
			Adslot: &Adslot{
				Adslot_id: adid,
				Adslot_size: &Size{
					Width:  w,
					Height: h,
				},
			},
			Device: &Device{
				Device_type: 1,
				Os_type:     1,
				Os_version: &Version{
					Major: 1,
					Minor: 1,
					Micro: 1,
				},
				Vendor: getReq.Vendor,
				Manufacturer:  getReq.Vendor,
				Model:  strings.Replace(getReq.Model, " ", "_", -1),
				Screen_size: &Size{
					Width:  sw,
					Height: sh,
				},
				Udid: &Udid{
					Idfa:       getReq.Idfa,
					Imei:       getReq.Imei,
					Android_id: getReq.Androidid,
					Mac:        mac,
				},
				User_agent:getReq.Ua,
			},
			Network: &Network{
				Ipv4:            getReq.Ip,
				Connection_type: 4,
				Operator_type:   3,
			},
			Gps: &Gps{
				Coordinate_type: 1,
				Latitude:        getReq.Lat,
				Longitude:       getReq.Lng,
				Timestamp:       strconv.Itoa(int(time.Now().Unix())),
			},
		},
	}
	if getReq.Os == "2" {
		postData.Data.Device.Os_type = 2
	}
	var m = [3]int{1, 1, 1}
	appversion := strings.Split(getReq.Appversion, ".")
	for index, value := range appversion {
		if index < 3 {
			m[index],_ = strconv.Atoi(value)
		}
	}
	postData.Data.App.App_version = &Version{
		Major: m[0],
		Minor: m[1],
		Micro: m[2],
	}

	var n = [3]int{1, 1, 1}
	osversion := strings.Split(getReq.Osversion, ".")
	for index, value := range osversion {
		if index < 3 {
			n[index],_ = strconv.Atoi(value)
		}
	}
	postData.Data.Device.Os_version = &Version{
		Major: n[0],
		Minor: n[1],
		Micro: n[2],
	}

	network := postData.Data.Network
	switch getReq.Network {
	case "以太网":
		network.Connection_type = 101
	case "wifi":
		network.Connection_type = 100
	case "2g":
		network.Connection_type = 2
	case "3g":
		network.Connection_type = 3
	case "4g":
		network.Connection_type = 4
	default:
		network.Connection_type = 999
	}
	// 运营商
	switch getReq.Imsi {
	case "1":
		network.Operator_type = 1
	case "2":
		network.Operator_type = 3
	case "3":
		network.Operator_type = 2
	default:
		network.Operator_type = 0
	}

	ma, err := json.Marshal(&postData)
	if err != nil {
		getReq.ChannelReq.Errorinfo = err.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}

	//fmt.Println("hhhhh",string(ma))

	req, err := http.NewRequest("POST", URL, bytes.NewReader(ma))
	if err != nil {
		getReq.ChannelReq.Errorinfo = err.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	resp, err := util.Client.Do(req)
	reqFunc(getReq)

	if err != nil {
		noFunc(getReq)
		return util.ResMsg{}
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		noFunc(getReq)
		return util.ResMsg{}
	}
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	resData := &adres{}
	json.Unmarshal(data, resData)

	if resData.Error_code != 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}
	ad := resData.Wxad
	if ad.Image_src == "" {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	imgArr := strings.Split(ad.Image_src, ";")
	if len(imgArr) == 0 || len(imgArr[0]) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	if len(ad.Landing_page_url) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}

	resultData := util.ResMsg{
		Id:       util.Md5(string(data) + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    ad.Ad_title,
		Content:  ad.Description,
		ImageUrl: imgArr[0],
		Uri:      replaceLdp(ad.Landing_page_url, wStr, hStr),
	}
	if resultData.Title == "" {
		resultData.Title = "ad"
	}

	if ad.Deep_link != "" {
		resultData.Scheme = replaceLdp(ad.Deep_link, wStr, hStr)
		for _, item := range ad.Dp_success_track_urls {
			resultData.Schemereport = append(resultData.Schemereport, replaceLdp(item, wStr, hStr))
		}
	}

	resultData.Displayreport = ad.Win_notice_url

	for _, v := range ad.Click_url {
		resultData.Clickreport = append(resultData.Clickreport, replaceLdp(v, wStr, hStr))
	}

	if ad.Interaction_type == 5 || ad.Interaction_type == 3 {
		if ad.Interaction_type == 3 {
			resultData.Json = true
		}

		for _, v := range ad.Download_track_urls {
			resultData.StartDownload = append(resultData.StartDownload, replaceLdp(v, wStr, hStr))
		}
		for _, v := range ad.Downloaded_track_urls {
			resultData.Downloaded = append(resultData.Downloaded, replaceLdp(v, wStr, hStr))
		}
		for _, v := range ad.Installed_track_urls {
			resultData.Installed = append(resultData.Installed, replaceLdp(v, wStr, hStr))
		}
		for _, v := range ad.Open_track_urls {
			resultData.Installed = append(resultData.Installed, replaceLdp(v, wStr, hStr))
		}
	}

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		getReq.ChannelReq.Errorinfo = "数据不完整"
		noimgFunc(getReq)
		return util.ResMsg{}
	}
	return resultData
}

func replaceLdp(urlStr, width, height string) string {
	urlStr = strings.Replace(urlStr, "__REQ_WIDTH__", width, -1)
	urlStr = strings.Replace(urlStr, "__REQ_HEIGHT__", height, -1)
	urlStr = strings.Replace(urlStr, "__CLICKID__", util.CLKID, -1)
	return urlStr
}