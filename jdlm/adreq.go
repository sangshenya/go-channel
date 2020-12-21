package jdlm

import (
	"github.com/sangshenya/go-channel/util"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func CommonParam(method, appkey, token string) url.Values {
	v := url.Values{}
	v.Set("method", method)
	v.Set("app_key", appkey)
	if len(token) != 0 {
		v.Set("access_token", token)
	}
	v.Set("timestamp", GetTaobaoTimeString())
	v.Set("format", "json")
	v.Set("v", "1.0")
	v.Set("sign_method", "md5")

	return v
}

func CreateJingdongSign(v url.Values, appsecret string) string {
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

func GetTaobaoTimeString() string {
	return strconv.Itoa(util.NowYear()) + "-" + fixTimeString(util.NowMonth()) + "-" + fixTimeString(util.NowDay())+ " " + fixTimeString(util.NowHour()) + ":" + fixTimeString(util.NowMinute()) + ":"+ fixTimeString(util.NowSecond())
}

func fixTimeString(timeIndex int) string {
	timeString := strconv.Itoa(timeIndex)
	if timeIndex < 10 {
		timeString = "0" + timeString
	}
	return timeString
}

// 猜你喜欢
type likeJson struct {
	GoodsReq LikeReq `json:"goodsReq"`
}
type LikeReq struct {
	EliteId int `json:"eliteId"`
	PageIndex int `json:"pageIndex"`
	PageSize int `json:"pageSize"`
	Pid string `json:"pid"`
	SiteId string `json:"siteId"`
	PositionId string `json:"positionId"`
	UserIdType int `json:"userIdType"`
	UserId string `json:"userId"`
}