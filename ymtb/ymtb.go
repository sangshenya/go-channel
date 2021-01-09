package ymtb

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
	URL = "https://publisher-api.deepleaper.com/goods"
)

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	uid := getReq.Imei
	uidtype := "imei"
	os := "android"
	if len(uid) == 0 {
		uid = getReq.Mac
		uidtype = "mac"
	}
	if getReq.Os == "2" {
		uid = getReq.Idfa
		uidtype = "idfa"
		os = "ios"
	}

	network := "unknown"
	switch getReq.Network {
	case "wifi":
		network = "wifi"
	case "2g":
		network = "2G"
	case "3g":
		network = "3G"
	case "4g":
		network = "4G"
	default:
		network = "unknown"
	}

	pid := getReq.ChannelReq.Adid
	channelid := getReq.ChannelReq.Appid

	if len(pid) == 0 || len(channelid) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数部分参数为空")
		return util.ResMsg{}, channelError
	}

	// pid=**&channelid=**

	postdata := adreq{
		Version:    "1",
		Id:         util.GetRandom(),
		Pid:        pid,
		Channel_id: channelid,
		User:       _user{
			Uid:            uid,
			Uid_type:       uidtype,
			Uid_encryption: "NA",
		},
		Device:     _device{
			Ipv4:        getReq.Ip,
			Device_type: "phone",
			Device_make: getReq.Vendor,
			Device_os:   os,
			Network:     network,
		},
	}

	ma, error := json.Marshal(postdata)
	if error != nil {
		channelError := util.NewChannelRequestFailErrorError(error)
		return util.ResMsg{}, channelError
	}

	req, err := http.NewRequest("POST", URL, bytes.NewReader(ma))
	if err != nil {
		channelError := util.NewChannelRequestFailErrorError(err)
		return util.ResMsg{}, channelError
	}
	//req.Header.Set("X-Forwarded-For", getReq.Ip)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")


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

	resData := &adres{}
	err = json.Unmarshal(data, resData)
	if err != nil {
		channelError := util.NewChannelRequestNoError(err)
		return util.ResMsg{}, channelError
	}

	if resData.Status != 0 {
		channelError := util.NewChannelRequestNoErrorWithText("Status不为0")
		return util.ResMsg{}, channelError
	}

	ad := resData.Creative

	imgurl := ad.Img
	if len(ad.Img) == 0 && len(ad.Imgs) != 0 {
		imgurl = ad.Imgs[0]
	}

	if len(imgurl) == 0 {
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	if len(ad.Clk_url) == 0 {
		channelError := util.NewChannelNoUrlErrorWithText("落地页链接长度为0")
		return util.ResMsg{}, channelError
	}

	postData := util.ResMsg{
		Id:       				 util.Md5(string(data) + time.Now().String()),
		Title:                   ad.Title,
		Content:                 ad.Title,
		ImageUrl:                imgurl,
		Uri:                     ad.Clk_url,
		Scheme:                  ad.Deeplink_url,
		ImageList:               ad.Imgs,
	}

	if postData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return postData, nil
}