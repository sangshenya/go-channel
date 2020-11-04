package qrqm

import (
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

const(
	TAOBAO_URL = "http://gw.api.taobao.com/router/rest?"
)

var ImageArray = []string{
	"https://img.admobile.top/admobile-adRequest/1_taobao.jpg",
	"https://img.admobile.top/admobile-adRequest/2_taobao.jpg",
	"https://img.admobile.top/admobile-adRequest/3_taobao.jpg",
	"https://img.admobile.top/admobile-adRequest/4_taobao.jpg",
}

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {

	os := "android"
	if getReq.Os == "2" {
		os = "ios"
	}

	network := 0
	switch getReq.Network {
	case "wifi":
		network = 1
	case "2g":
		network = 2
	case "3g":
		network = 3
	case "4g":
		network = 4
	default:
		network = 0
	}

	adid := getReq.ChannelReq.Adid
	if len(adid) == 0 {
		getReq.ChannelReq.Errorinfo = "adid不能为空"
		failFunc(getReq)
		return util.ResMsg{}
	}
	appkey := getReq.ChannelReq.Appid
	if len(appkey) == 0 {
		getReq.ChannelReq.Errorinfo = "appid不能为空"
		failFunc(getReq)
		return util.ResMsg{}
	}
	appsecret := getReq.ChannelReq.Token
	if len(appsecret) == 0 {
		getReq.ChannelReq.Errorinfo = "appsecret不能为空"
		failFunc(getReq)
		return util.ResMsg{}
	}

	req := tbreq{
		Id:        util.GetRandom(),
		Version:   1,
		Device:    tbdevice{
			Osv:         getReq.Osversion,
			Os:          os,
			Ip:          getReq.Ip,
			Idfa:        getReq.Idfa,
			Imei:        getReq.Imei,
			Oaid:        getReq.Oaid,
			Mac:         getReq.Mac,
			Network:     network,
			Imei_md5:    util.Md5(getReq.Imei),
			Device_type: 0,
		},
		Adzone_id: adid,
	}

	ma, _ := json.Marshal(&req)

	v := CommonParam("taobao.tbk.thor.creative.launch", appkey)
	v.Set("req", string(ma))

	sign := CreateTaobaoSign(v, appsecret)
	v.Set("sign", sign)

	header := http.Header{}
	response, error := util.HttpGET(TAOBAO_URL + v.Encode(), &header)
	reqFunc(getReq)
	if error != nil {
		getReq.ChannelReq.Errorinfo = error.Error()
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		getReq.ChannelReq.Errorinfo = error.Error()
		noFunc(getReq)
		return util.ResMsg{}
	}

	if response.StatusCode != 200 {
		getReq.ChannelReq.Errorinfo = "状态码不为200"
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	//fmt.Println("GetTaobaoQrqmData:", string(data))

	resData := &tbadres{}
	json.Unmarshal(data, resData)
	if len(resData.Tbk_thor_creative_launch_response.Result.Click_through_url) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}

	imageUrl := ""
	if len(resData.Tbk_thor_creative_launch_response.Result.Image_url) != 0 {
		imageUrl = resData.Tbk_thor_creative_launch_response.Result.Image_url
	} else if len(resData.Tbk_thor_creative_launch_response.Result.Image_url2) != 0 {
		imageUrl = resData.Tbk_thor_creative_launch_response.Result.Image_url
	} else if len(resData.Tbk_thor_creative_launch_response.Result.Image_url3) != 0 {
		imageUrl = resData.Tbk_thor_creative_launch_response.Result.Image_url
	} else if len(resData.Tbk_thor_creative_launch_response.Result.Image_url4) != 0 {
		imageUrl = resData.Tbk_thor_creative_launch_response.Result.Image_url
	} else {
		index := rand.Int() % len(ImageArray)
		imageUrl = ImageArray[index]
	}

	title := "广告"
	if len(resData.Tbk_thor_creative_launch_response.Result.Title) != 0 {
		title = resData.Tbk_thor_creative_launch_response.Result.Title
	}

	content := "广告"
	if len(resData.Tbk_thor_creative_launch_response.Result.Desc) != 0 {
		title = resData.Tbk_thor_creative_launch_response.Result.Desc
	}

	postData := util.ResMsg{
		Title:         title,
		Content:       content,
		ImageUrl:      imageUrl,
		Uri:           resData.Tbk_thor_creative_launch_response.Result.Click_through_url,
	}

	if len(resData.Tbk_thor_creative_launch_response.Result.Impression_tracking_url) != 0 {
		postData.Displayreport = append(postData.Displayreport, resData.Tbk_thor_creative_launch_response.Result.Impression_tracking_url)
	}

	if len(resData.Tbk_thor_creative_launch_response.Result.Deeplink_url) != 0 {
		postData.Scheme = resData.Tbk_thor_creative_launch_response.Result.Deeplink_url
	}

	if len(postData.ImageUrl) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	if len(postData.Uri) == 0 {
		nourlFunc(getReq)
		return util.ResMsg{}
	}

	return postData
}

func CommonParam(method, appkey string) url.Values {
	v := url.Values{}
	v.Set("method", method)
	v.Set("app_key", appkey)
	v.Set("sign_method", "md5")
	v.Set("timestamp", getTaobaoTimeString())
	v.Set("format", "json")
	v.Set("v", "2.0")

	return v
}

func CreateTaobaoSign(v url.Values, appsecret string) string {
	valueString := urlValuesEncode(v)
	return strings.ToUpper(util.Md5(appsecret + valueString + appsecret))
}

func urlValuesEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	urlstring := ""

	for _, k := range keys {
		vs := v[k]
		v := ""
		if len(vs) != 0 {
			v = vs[0]
		}
		urlstring =  urlstring + k + v
	}

	return urlstring
}

func getTaobaoTimeString() string {
	return strconv.Itoa(util.NowYear()) + "-" + fixTimeString(util.NowMonth()) + "-" + fixTimeString(util.NowDay())+ " " + fixTimeString(util.NowHour()) + ":" + fixTimeString(util.NowMinute()) + ":"+ fixTimeString(util.NowSecond())
}

func fixTimeString(timeIndex int) string {
	timeString := strconv.Itoa(timeIndex)
	if timeIndex < 10 {
		timeString = "0" + timeString
	}
	return timeString
}