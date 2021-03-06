package inmobi

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

const (
	URL = "https://union.api.w.inmobi.cn/showad/v3.3"
)

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	ct := 2
	if getReq.Network == "4g" {
		ct = 6
	}

	swidth, err := strconv.Atoi(getReq.Screenwidth)
	if err != nil {
		swidth = 0
	}
	sheight, err := strconv.Atoi(getReq.Screenheight)
	if err != nil {
		sheight = 0
	}

	lat, err := strconv.ParseFloat(getReq.Lat, 64)
	if err != nil {
		lat = 0
	}
	lon, err := strconv.ParseFloat(getReq.Lng, 64)
	if err != nil {
		lon = 0
	}

	os := "iOS"
	ifa := getReq.Idfa
	um5 := ""
	o1 := ""
	imei_sha1 := ""
	imei_md5 := ""

	if getReq.Os == "1" {
		os = "Android"
		imei_md5 = util.Md5(getReq.Imei)
		imei_sha1 = util.Sha1(getReq.Imei)
		o1 = util.Sha1(strings.ToLower(getReq.Androidid))
		um5 = util.Md5(strings.ToLower(getReq.Androidid))
	}

	carrier := "ChinaMobile"
	if getReq.Imsi == "1" {
		carrier = "ChinaMobile"
	} else if getReq.Imsi == "2" {
		carrier = "ChinaUnicom"
	} else if getReq.Imsi == "3" {
		carrier = "ChinaTelecom"
	}

	sd, _ := strconv.Atoi(getReq.Sd)

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)

	adid := getReq.ChannelReq.Adid
	adType := getReq.ChannelReq.Adtype
	pkg := getReq.ChannelReq.Pkg
	w, _ := paramsMap["w"]
	h, _ := paramsMap["h"]

	if len(adid) == 0 || len(pkg) == 0 || len(adType) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数部分参数为空")
		return util.ResMsg{}, channelError
	}

	postData := adreq{
		App:_app{
			Id:adid,
			//Id:"1463152075376",
			Bundle:pkg,
			//Bundle:"com.inmobi.test",
			Ver:getReq.Appversion,
			Orientation:1,
			Paid:0,
		},
		Imp:_imp{
			Ext:_impext{
				AdsCount:1,
			},
		},
		Device:_device{
			Ifa:ifa,
			//Ifa:"FC0F3445-0FCE-40EE-8646-3CA8BB2663EA",
			Oaid:getReq.Oaid,
			Md5_imei:imei_md5,
			Sha1_imei:imei_sha1,
			O1:o1,
			Um5:um5,
			Type:"1",
			Ua:getReq.Ua,
			//Ua:"Mozilla/5.0 (iPhone; CPU iPhone OS 8_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, likeGecko) Version/8.0 Mobile/12D436",
			Ip:getReq.Ip,
			//Ip:"3.0.119.0",
			Os:os,
			Osv:getReq.Osversion,
			Model:getReq.Model,
			Geo:_geo{
				Lat:lat,
				Lon:lon,
				Accu:0,
			},
			Carrier:carrier,
			//Carrier:"ChinaUnicom",
			Connectiontype:ct,
			//Connectiontype:2,
			Ext:_deviceExt{
				Orientation:1,
			},
			Idfv:getReq.Idfv,
			Make:getReq.Vendor,
			H:sheight,
			W:swidth,
			Ppi:sd,
		},
		Ext:_ext{
			Responseformat:"json",
			SupportDeeplink:true,
		},
	}

	if getReq.Os == "2" {
		postData.Imp.Secure = 0
	}

	if adType == "banner" {
		bannerW, _ := strconv.Atoi(w)
		bannerH, _ := strconv.Atoi(h)
		postData.Imp.Banner = &_banner{
			W:bannerW,
			H:bannerH,
		}
	} else {
		postData.Imp.Native = &_native{
			Layout:0,
		}
		postData.Imp.Trackertype = "url_ping"
	}

	ma, err := json.Marshal(&postData)
	if err != nil{
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}

	req, err := http.NewRequest("POST", URL, bytes.NewReader(ma))
	if err != nil {
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	resp, err := util.Client.Do(req)
	reqFunc(getReq)

	if err != nil {
		channelError := util.NewChannelRequestTimeoutError(err)
		return util.ResMsg{}, channelError
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		channelError := util.NewChannelRequestNoError(err)
		return util.ResMsg{}, channelError
	}

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
	//fmt.Println("11111,",string(res))

	ads := resData.Ads
	if len(ads) == 0 {
		channelError := util.NewChannelRequestNoErrorWithText("ads长度为0")
		return util.ResMsg{}, channelError
	}

	ad := ads[0]
	img := ""
	title := ""
	content := ""
	ldp := ""
	scheme := ""

	pub := ad.PubContent
	img = pub.Screenshots.Url
	title = pub.Title
	content = pub.Description
	ldp = ad.TargetUrl
	scheme = pub.LandingURL
	if len(scheme) == 0 {
		scheme = ad.LandingURL
	}

	if len(img) == 0 {
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(ldp) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("落地页链接长度为0")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(string(data) + time.Now().String()),
		Weight:   0,
		State:    1,
		Title:    title,
		Content:  content,
		ImageUrl: img,
		Uri:      replace(ldp, w, h),
	}

	if ad.OpenExternal {
		resultData.Scheme = scheme
		for _, item := range ad.EventTracking.DplSuccess {
			resultData.Schemereport = append(resultData.Schemereport, replace(item, w, h))
		}
	}

	for _, item := range ad.EventTracking.ImpressionTrackers {
		resultData.Displayreport = append(resultData.Displayreport, replace(item, w, h))
	}

	for _,clickUrl := range ad.EventTracking.ClickTrackers {
		resultData.Clickreport = append(resultData.Clickreport, replace(clickUrl, w, h))
	}

	for _,downloadUrl := range ad.EventTracking.DownloadStart{
		resultData.StartDownload = append(resultData.StartDownload, replace(downloadUrl, w, h))
	}

	for _,downloadUrl := range ad.EventTracking.DownloadFinish{
		resultData.Downloaded = append(resultData.Downloaded, replace(downloadUrl, w, h))
	}

	for _,downloadUrl := range ad.EventTracking.InstallFinish{
		resultData.Installed = append(resultData.Installed, replace(downloadUrl, w, h))
	}

	if resultData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return resultData, nil
}

func replace(urlStr string, w, h string) string {

	urlStr = strings.Replace(urlStr, "__REQ_WIDTH__", w, -1)
	urlStr = strings.Replace(urlStr, "__REQ_HEIGHT__", h, -1)

	urlStr = strings.Replace(urlStr, "$TS", util.TS, -1)
	return urlStr
}

