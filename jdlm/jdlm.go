package jdlm

import (
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
)

const(
	URL = "https://router.jd.com/api?"

	// ecook
	APPID_ECOOK_I = "4100249276"
	APPKEY_ECOOK_I = "48dd082fdb5841c95409e52c1b9db083"
	APPSECRET_ECOOK_I = "b3350050d63d4c438cb1e0728221cd2d"

	PID_ECOOK_I_F = "1003249735_4100249276_3003262349"
	SUBPID_ECOOK_I_F = "3003262349"

	PID_ECOOK_I_AF = "1003249735_4100249276_3003271120"
	SUBPID_ECOOK_I_AF = "3003271120"

	APPID_ECOOK_A = "4100249275"
	APPKEY_ECOOK_A = "48dd082fdb5841c95409e52c1b9db083"
	APPSECRET_ECOOK_A = "b3350050d63d4c438cb1e0728221cd2d"

	PID_ECOOK_A_F = "1003249735_4100249275_3003263356"
	SUBPID_ECOOK_A_F = "3003263356"

	PID_ECOOK_A_AF = "1003249735_4100249275_3003263355"
	SUBPID_ECOOK_A_AF = "3003263355"
)

func Base(getReq *util.ReqMsg, failFunc util.ReqFailFunc, reqFunc util.ReqFunc, noFunc util.ReqNoFunc, timeoutFunc util.ReqTimeoutFunc, noimgFunc util.ReqNoimgFunc, nourlFunc util.ReqNourlFunc) util.ResMsg {

	appid := getReq.ChannelReq.Appid
	subpid := getReq.ChannelReq.Adid
	appsecret := getReq.ChannelReq.Token

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)
	appkey, _ := paramsMap["appkey"]

	if len(appid) == 0 || len(subpid) == 0 || len(appsecret) == 0 || len(appkey) == 0 {
		getReq.ChannelReq.Errorinfo = "请求配置参数不完整"
		failFunc(getReq)
		return util.ResMsg{}
	}

	v := CommonParam("jd.union.open.goods.material.query", appkey, "")

	uid := ""
	uidType := 0

	if getReq.Os == "2" {
		if len(getReq.Idfa) != 0 {
			uid = getReq.Idfa
			uidType = 32
		} else if len(getReq.Openudid) != 0 {
			uid = getReq.Openudid
			uidType = 16
		}
	} else {
		if len(getReq.Imei) != 0 {
			uid = getReq.Imei
			uidType = 8
		} else if len(getReq.Oaid) != 0 {
			uid = getReq.Oaid
			uidType = 32768
		}
	}

	if len(uid) == 0 {
		getReq.ChannelReq.Errorinfo = "流量携带参数不完整"
		failFunc(getReq)
		return util.ResMsg{}
	}
	likeReq := likeJson{GoodsReq:LikeReq{
		EliteId:    1,
		PageIndex:  1,
		PageSize:   1,
		SiteId:     appid,
		PositionId: subpid,
		UserIdType: uidType,
		UserId:     uid,
	}}

	ma, error := json.Marshal(likeReq)
	if error != nil {
		getReq.ChannelReq.Errorinfo = error.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}

	//fmt.Println(string(ma), appkey, appsecret)

	v.Set("param_json", string(ma))
	sign := CreateJingdongSign(v, appsecret)
	v.Set("sign", sign)

	request, error := http.NewRequest("GET", URL + v.Encode(), nil)
	if error != nil {
		getReq.ChannelReq.Errorinfo = error.Error()
		failFunc(getReq)
		return util.ResMsg{}
	}
	response, error := util.Client.Do(request)
	reqFunc(getReq)
	if error != nil {
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		noFunc(getReq)
		return util.ResMsg{}
	}

	//fmt.Println("base2:", string(data))
	if response.StatusCode != 200 {
		timeoutFunc(getReq)
		return util.ResMsg{}
	}

	likeResult := LikeRes{}
	json.Unmarshal(data, &likeResult)

	likeResultData := LikeResultData{}
	if len(likeResult.Jd_union_open_goods_material_query_response.Result) != 0 {
		//fmt.Println(likeResult.Jd_union_open_goods_material_query_response.Result)
		json.Unmarshal([]byte(likeResult.Jd_union_open_goods_material_query_response.Result), &likeResultData)
	} else {
		noFunc(getReq)
		return util.ResMsg{}
	}
	if likeResultData.Code != 200 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	if len(likeResultData.Data) == 0 {
		noFunc(getReq)
		return util.ResMsg{}
	}

	likeData := likeResultData.Data[0]

	if len(likeData.ImageInfo.ImageList) == 0 {
		noimgFunc(getReq)
		return util.ResMsg{}
	}

	postData := util.ResMsg{
		Title:                   likeData.BrandCode,
		Content:                 likeData.SkuName,
		ImageUrl:                likeData.ImageInfo.ImageList[0].Url,
		Uri:                     likeData.PromotionInfo.ClickURL,
	}

	if postData.ResponseDataIsEmpty(getReq.Adtype) {
		getReq.ChannelReq.Errorinfo = "数据不完整"
		noFunc(getReq)
		return util.ResMsg{}
	}

	return postData
}