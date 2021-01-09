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

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {

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
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数adid不能为空")
		return util.ResMsg{}, channelError
	}
	appkey := getReq.ChannelReq.Appid
	if len(appkey) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数appid不能为空")
		return util.ResMsg{}, channelError
	}
	appsecret := getReq.ChannelReq.Token
	if len(appsecret) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数appsecret不能为空")
		return util.ResMsg{}, channelError
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

	ma, err := json.Marshal(&req)
	if err != nil{
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}

	v := CommonParam("taobao.tbk.thor.creative.launch", appkey)
	v.Set("req", string(ma))

	sign := CreateTaobaoSign(v, appsecret)
	v.Set("sign", sign)

	header := http.Header{}
	response, error := util.HttpGET(TAOBAO_URL + v.Encode(), &header)
	reqFunc(getReq)
	if error != nil {
		channelError := util.NewChannelRequestFailErrorError(error)
		return util.ResMsg{}, channelError
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		channelError := util.NewChannelRequestNoError(error)
		return util.ResMsg{}, channelError
	}

	if response.StatusCode != 200 {
		code := response.StatusCode
		channelError := util.NewChannelRequestNoErrorWithText("状态码为:"+ strconv.Itoa(int(code)))
		return util.ResMsg{}, channelError
	}

	//fmt.Println("GetTaobaoQrqmData:", string(data))

	resData := &tbadres{}
	err = json.Unmarshal(data, resData)
	if err != nil {
		channelError := util.NewChannelRequestNoError(err)
		return util.ResMsg{}, channelError
	}

	if len(resData.Tbk_thor_creative_launch_response.Result.Click_through_url) == 0 {
		channelError := util.NewChannelRequestNoErrorWithText("Click_through_url长度为0")
		return util.ResMsg{}, channelError
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
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(postData.Uri) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if postData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return postData, nil
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