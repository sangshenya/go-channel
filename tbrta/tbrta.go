package tbrta

import (
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const(
	//TAOBAO_URL = "http://gw.api.taobao.com/router/resimt?"
	URL = "http://gw.api.taobao.com/router/rest?"
	TAOBAO_URL_TEST = "http://gw.api.tbsandbox.com/router/rest?"
	TAOBAO_URL = "http://gw.api.taobao.com/router/rest?"

	CHANNEL = "4183462252"
	ADID_ECOOK_A_S = "9635"
	APPKEY_ECOOK = "28334935"
	APPSECRET_ECOOK = "964b93b2e1ec319bc6fdfc3ee457f76e"
)

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	if getReq.ChannelReq.Adtype != "startup" && getReq.ChannelReq.Adtype != "splashad" {
		channelError := util.NewChannelRequestFailErrorWithText("不支持的广告请求类型")
		return util.ResMsg{}, channelError
	}

	reqFunc(getReq)

	if TaobaoTarget(getReq) {
		channelError := util.NewChannelRequestNoErrorWithText("淘宝rta定向不匹配")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    "广告",
		Content:  "广告",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/tb_zsp.png",
		Uri:      "https://star-link.taobao.com?bc_fl_src=growth_dhh_4183462252_100-12768-32896&dpa_Inid=3289610158&dpa_material_id=635763217806&dpa_material_type=1&dpa_source_code=10158&force_no_smb=true&itemIds=635763217806&slk_actid=100000000207&spm=2014.ugdhh.4183462252.100-12768-32896&wh_biz=tm",
		Scheme:   "taobao://m.taobao.com/tbopen/index.html?action=ali.open.nav&bc_fl_src=growth_dhh_4183462252_100-12768-32896&bootImage=0&dpa_Inid=3289610158&dpa_material_id=635763217806&dpa_material_type=1&dpa_source_code=10158&force_no_smb=true&itemIds=635763217806&module=h5&slk_actid=100000000207&source=auto&spm=2014.ugdhh.4183462252.100-12768-32896&wh_biz=tm&h5Url=https://star-link.taobao.com?bc_fl_src%3Dgrowth_dhh_4183462252_100-12768-32896%26dpa_Inid%3D3289610158%26dpa_material_id%3D635763217806%26dpa_material_type%3D1%26dpa_source_code%3D10158%26force_no_smb%3Dtrue%26itemIds%3D635763217806%26slk_actid%3D100000000207%26spm%3D2014.ugdhh.4183462252.100-12768-32896%26wh_biz%3Dtm",
	}

	return resultData, nil
}

func Base2(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	if getReq.ChannelReq.Adtype != "flow" {
		channelError := util.NewChannelRequestFailErrorWithText("不支持的广告请求类型")
		return util.ResMsg{}, channelError
	}

	reqFunc(getReq)

	if TaobaoTarget(getReq) {
		channelError := util.NewChannelRequestNoErrorWithText("淘宝rta定向不匹配")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    "广告",
		Content:  "广告",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/tbdhhwflow_rta.jpg",
		Uri:      "https://star-link.taobao.com?bc_fl_src=growth_dhh_4183462252_100-21253-32896&dpa_Inid=3289610158&dpa_material_id=635763217806&dpa_material_type=1&dpa_source_code=10158&force_no_smb=true&itemIds=635763217806&slk_actid=100000000207&spm=2014.ugdhh.4183462252.100-21253-32896&wh_biz=tm",
		Scheme:   "tbopen://m.taobao.com/tbopen/index.html?action=ali.open.nav&bc_fl_src=growth_dhh_4183462252_100-21253-32896&bootImage=0&dpa_Inid=3289610158&dpa_material_id=635763217806&dpa_material_type=1&dpa_source_code=10158&force_no_smb=true&itemIds=635763217806&module=h5&slk_actid=100000000207&source=auto&spm=2014.ugdhh.4183462252.100-21253-32896&wh_biz=tm&h5Url=https://star-link.taobao.com?bc_fl_src%3Dgrowth_dhh_4183462252_100-21253-32896%26dpa_Inid%3D3289610158%26dpa_material_id%3D635763217806%26dpa_material_type%3D1%26dpa_source_code%3D10158%26force_no_smb%3Dtrue%26itemIds%3D635763217806%26slk_actid%3D100000000207%26spm%3D2014.ugdhh.4183462252.100-21253-32896%26wh_biz%3Dtm",
	}

	return resultData, nil
}

func Base3(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	if getReq.ChannelReq.Adtype != "flow" {
		channelError := util.NewChannelRequestFailErrorWithText("不支持的广告请求类型")
		return util.ResMsg{}, channelError
	}

	reqFunc(getReq)

	if TaobaoTarget(getReq) {
		channelError := util.NewChannelRequestNoErrorWithText("淘宝rta定向不匹配")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    "广告",
		Content:  "广告",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_320_f.jpg",
		Uri:      "https://star-link.taobao.com?bc_fl_src=growth_dhh_4183462252_100-13678-32896&dpa_Inid=3289610118&dpa_material_id=622110332794&dpa_material_type=1&dpa_source_code=10118&force_no_smb=true&itemIds=622110332794&slk_actid=100000000207&spm=2014.ugdhh.4183462252.100-13678-32896&wh_biz=tm",
		Scheme:   "tbopen://m.taobao.com/tbopen/index.html?action=ali.open.nav&bc_fl_src=growth_dhh_4183462252_100-13678-32896&bootImage=0&dpa_Inid=3289610118&dpa_material_id=622110332794&dpa_material_type=1&dpa_source_code=10118&force_no_smb=true&itemIds=622110332794&module=h5&slk_actid=100000000207&source=auto&spm=2014.ugdhh.4183462252.100-13678-32896&wh_biz=tm&h5Url=https://star-link.taobao.com?bc_fl_src%3Dgrowth_dhh_4183462252_100-13678-32896%26dpa_Inid%3D3289610118%26dpa_material_id%3D622110332794%26dpa_material_type%3D1%26dpa_source_code%3D10118%26force_no_smb%3Dtrue%26itemIds%3D622110332794%26slk_actid%3D100000000207%26spm%3D2014.ugdhh.4183462252.100-13678-32896%26wh_biz%3Dtm",
	}

	return resultData, nil
}

func Base4(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {
	if getReq.ChannelReq.Adtype != "startup" && getReq.ChannelReq.Adtype != "splashad" {
		channelError := util.NewChannelRequestFailErrorWithText("不支持的广告请求类型")
		return util.ResMsg{}, channelError
	}

	reqFunc(getReq)

	if TaobaoTarget(getReq) {
		channelError := util.NewChannelRequestNoErrorWithText("淘宝rta定向不匹配")
		return util.ResMsg{}, channelError
	}

	resultData := util.ResMsg{
		Id:       util.Md5(util.GetRandom() + time.Now().String()),
		Weight:   0,
		State:    0,
		Title:    "广告",
		Content:  "广告",
		ImageUrl: "https://img.admobile.top/admobile-adRequest/dhh_320_s.jpg",
		Uri:      "https://star-link.taobao.com?bc_fl_src=growth_dhh_4183462252_100-12768-32896&dpa_Inid=3289610118&dpa_material_id=622110332794&dpa_material_type=1&dpa_source_code=10118&force_no_smb=true&itemIds=622110332794&slk_actid=100000000207&spm=2014.ugdhh.4183462252.100-12768-32896&wh_biz=tm",
		Scheme:   "tbopen://m.taobao.com/tbopen/index.html?action=ali.open.nav&bc_fl_src=growth_dhh_4183462252_100-12768-32896&bootImage=0&dpa_Inid=3289610118&dpa_material_id=622110332794&dpa_material_type=1&dpa_source_code=10118&force_no_smb=true&itemIds=622110332794&module=h5&slk_actid=100000000207&source=auto&spm=2014.ugdhh.4183462252.100-12768-32896&wh_biz=tm&h5Url=https://star-link.taobao.com?bc_fl_src%3Dgrowth_dhh_4183462252_100-12768-32896%26dpa_Inid%3D3289610118%26dpa_material_id%3D622110332794%26dpa_material_type%3D1%26dpa_source_code%3D10118%26force_no_smb%3Dtrue%26itemIds%3D622110332794%26slk_actid%3D100000000207%26spm%3D2014.ugdhh.4183462252.100-12768-32896%26wh_biz%3Dtm",
	}

	return resultData, nil
}

func TaobaoTarget(getReq *util.ReqMsg) bool {
	adres := TaobaoRta(getReq, CHANNEL, ADID_ECOOK_A_S, APPKEY_ECOOK, APPSECRET_ECOOK)

	if !adres.Usergrowth_dhh_delivery_ask_response.Result {
		return false
	}
	return true
}

func TaobaoRta(getReq *util.ReqMsg, channel, adid, appkey, appsecret string) newAdres {
	if len(getReq.Imei) == 0 && len(getReq.Idfa) == 0 && len(getReq.Oaid) == 0 {
		return newAdres{}
	}

	v := CommonParam("taobao.usergrowth.dhh.delivery.ask", appkey)

	v.Set("advertising_space_id", adid)
	v.Set("channel", channel)
	if getReq.Os == "2" {
		v.Set("idfa_md5", util.Md5(getReq.Idfa))
		v.Set("os", "1")
	} else {
		v.Set("os", "0")
		if len(getReq.Imei) != 0 {
			v.Set("imei_md5", util.Md5(getReq.Imei))
		} else if len(getReq.Oaid) != 0 {
			v.Set("oaid_md5", util.Md5(getReq.Oaid))
		}
	}

	sign := CreateTaobaoSign(v, appsecret)
	v.Set("sign", sign)


	request, error := http.NewRequest("GET", URL + v.Encode(), nil)
	if error != nil {
		return newAdres{}
	}
	response, error := util.Client.Do(request)
	if error != nil {
		return newAdres{}
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		return newAdres{}
	}

	//fmt.Println("base2:", string(data))
	if response.StatusCode != 200 {
		return newAdres{}
	}

	resData := newAdres{}
	json.Unmarshal(data, &resData)

	return resData
}


func CommonParam(method, appkey string) url.Values {
	v := url.Values{}
	v.Set("method", method)
	v.Set("app_key", appkey)
	v.Set("sign_method", "md5")
	v.Set("timestamp", GetTaobaoTimeString())

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

func ChangeTime(timeString string) int {
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	loc, _ := time.LoadLocation("Asia/Shanghai")
	stamp,_ := time.ParseInLocation(timeTemplate1, timeString, loc)
	return int(stamp.Unix())
}