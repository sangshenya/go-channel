package oneway

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	URL          = "https://ads.oneway.mobi/getCampaign?"
)

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	deviceId := ""
	os := 1
	if getReq.Os == "2" {
		os = 2
		deviceId = getReq.Idfa
	}

	connection := "wifi"
	if getReq.Network == "4g" || getReq.Network == "3g" {
		connection = "cellular"
	}

	carrier := "46000"

	switch getReq.Imsi {
	case "1":
		carrier = "46000"
	case "2":
		carrier = "46001"
	case "3":
		carrier = "46003"
	default:
		carrier = "46000"
	}

	width, err := strconv.Atoi(getReq.Screenwidth)
	if err != nil {
		width = 0
	}
	height, err := strconv.Atoi(getReq.Screenheight)
	if err != nil {
		height = 0
	}

	sd, err := strconv.Atoi(getReq.Sd)
	//dpi := sd / 160


	apiLevel,_ := strconv.Atoi(getReq.AndroidApiLevel)

	lat, err := strconv.ParseFloat(getReq.Lat, 64)
	if err != nil {
		lat = 0
	}
	lon, err := strconv.ParseFloat(getReq.Lng, 64)
	if err != nil {
		lon = 0
	}

	if len(getReq.ChannelReq.Adid) == 0 || len(getReq.ChannelReq.Pkg) == 0 || len(getReq.ChannelReq.Appname) == 0 || len(getReq.ChannelReq.Appid) == 0 || len(getReq.ChannelReq.Token) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数部分参数为空")
		return util.ResMsg{}, channelError
	}

	postData := adreq{
		ApiVersion: 	"1.5.4",
		PlacementId: 	getReq.ChannelReq.Adid,
		BundleId:       getReq.ChannelReq.Pkg,
		BundleVersion:  getReq.Appversion,
		AppName:        getReq.ChannelReq.Appname,
		SubAffId:		"admobile",
		DeviceId: 		deviceId,
		Imei:           getReq.Imei,
		AndroidId:      getReq.Androidid,
		Oaid:			getReq.Oaid,
		ApiLevel: 		apiLevel,
		Os:				os,
		OsVersion:      getReq.Osversion,
		ConnectionType: connection,
		NetworkType:    1,
		NetworkOperator:carrier,
		SimOperator:	carrier,
		Imsi:			getReq.Imsi_long,
		DeviceMake:     getReq.Vendor,
		DeviceModel:    getReq.Model,
		DeviceType:     1,
		Orientation:    "H",
		Mac: 			getReq.Mac,
		ScreenWidth: 	width,
		ScreenHeight: 	height,
		ScreenDensity: 	sd,
		UserAgent:      getReq.Ua,
		Ip: 			getReq.Ip,
		Language:       "zh_CN",
		TimeZone:       "GMT+08:00",
		Latitude:       lat,
		Longitude: 		lon,
	}

	ma, err := json.Marshal(&postData)
	if err != nil {
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}

	requestURL := URL+"publishId="+getReq.ChannelReq.Appid+"&token="+getReq.ChannelReq.Token+"&ts="+strconv.Itoa(int(time.Now().Unix()))

	//fmt.Println(string(ma),requestURL)

	req, err := http.NewRequest("POST", requestURL, bytes.NewReader(ma))
	if err != nil {
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := util.Client.Do(req)
	reqFunc(getReq)
	if err != nil {
		channelError := util.NewChannelRequestTimeoutError(err)
		return util.ResMsg{}, channelError
	}

	data := []byte{}
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		r, err := gzip.NewReader(resp.Body)
		//resp.Body.Close()
		if err != nil {
			channelError := util.NewChannelRequestNoError(err)
			return util.ResMsg{}, channelError
		}
		defer r.Close()

		data, err = ioutil.ReadAll(r)
		if err != nil {
			channelError := util.NewChannelRequestNoError(err)
			return util.ResMsg{}, channelError
		}
		resp.Body.Close()
	} else {
		data, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}

	if err != nil {
		channelError := util.NewChannelRequestNoError(err)
		return util.ResMsg{}, channelError
	}
	//fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		code := resp.StatusCode
		channelError := util.NewChannelRequestNoErrorWithText("状态码为:"+ strconv.Itoa(int(code)))
		return util.ResMsg{}, channelError
	}
	resData := adres{}
	err = json.Unmarshal(data, &resData)
	if err != nil {
		channelError := util.NewChannelRequestNoError(err)
		return util.ResMsg{}, channelError
	}

	if !resData.Success {
		channelError := util.NewChannelRequestNoErrorWithText("Success is false")
		return util.ResMsg{}, channelError
	}

	ad := resData.Data
	if len(ad.Images) == 0 {
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	image := ad.Images[0]
	if len(image.Url) == 0 {
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(ad.ClickUrl) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(string(data) + time.Now().String()),
		Weight:   0,
		State:    1,
		Title:    ad.AppName,
		Content:  ad.Title,
		ImageUrl: image.Url,
		Uri:      replace(ad.ClickUrl),
	}

	if len(ad.Deeplink) != 0 {
		resultData.Scheme = replace(ad.Deeplink)
		if len(ad.TrackingEvents.Dp) != 0 {
			for _,item := range ad.TrackingEvents.Dp {
				resultData.Schemereport = append(resultData.Schemereport, replace(item))
			}
		}
	}

	for _,item := range ad.TrackingEvents.Show {
		resultData.Displayreport = append(resultData.Displayreport, replace(item))
	}

	for _,item := range ad.TrackingEvents.Click {
		resultData.Clickreport = append(resultData.Clickreport, replace(item))
	}

	for _,item := range ad.TrackingEvents.ApkDownloadStart {
		resultData.StartDownload = append(resultData.StartDownload, replace(item))
	}

	for _,item := range ad.TrackingEvents.ApkDownloadFinish {
		resultData.Downloaded = append(resultData.Downloaded, replace(item))
	}

	for _,item := range ad.TrackingEvents.PackageAdded {
		resultData.Installed = append(resultData.Installed, replace(item))
	}

	if ad.InteractionType == 3 {
		resultData.Json = true
	}

	if len(ad.AppStoreId) != 0 {
		resultData.Pkg = ad.AppStoreId
	}

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return resultData, nil
}

func replace(urlStr string) string {

	urlStr = strings.Replace(urlStr, "__TIMESTAMP__", util.TS, -1)
	urlStr = strings.Replace(urlStr, "__C_UP_TIME__", util.TS, -1)

	urlStr = strings.Replace(urlStr, "__C_DOWN_OFFSET_X__", util.DX, -1)
	urlStr = strings.Replace(urlStr, "__C_DOWN_OFFSET_Y__", util.DY, -1)

	urlStr = strings.Replace(urlStr, "__C_DOWN_X__", util.RDX, -1)
	urlStr = strings.Replace(urlStr, "__C_DOWN_Y__", util.RDY, -1)
	urlStr = strings.Replace(urlStr, "__C_UP_X__", util.RUX, -1)
	urlStr = strings.Replace(urlStr, "__C_UP_Y__", util.RUY, -1)

	return urlStr
}


