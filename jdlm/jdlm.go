package jdlm

import (
	"encoding/json"
	"github.com/sangshenya/go-channel/util"
	"io/ioutil"
	"net/http"
	"strconv"
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

func Base(getReq *util.ReqMsg, reqFunc util.ReqFunc) (util.ResMsg, util.ChannelErrorProtocol) {

	appid := getReq.ChannelReq.Appid
	subpid := getReq.ChannelReq.Adid
	appsecret := getReq.ChannelReq.Token

	paramsMap := util.ParamsEncode(getReq.ChannelReq.Params, getReq.ChannelReq.Adtype)
	appkey, _ := paramsMap["appkey"]

	if len(appid) == 0 || len(subpid) == 0 || len(appsecret) == 0 || len(appkey) == 0 {
		channelError := util.NewChannelRequestFailErrorWithText("请求必需参数部分参数为空")
		return util.ResMsg{}, channelError
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
		channelError := util.NewChannelRequestFailErrorWithText("流量携带参数不完整")
		return util.ResMsg{}, channelError
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
		channelError := util.NewChannelRequestFailErrorError(error)
		return util.ResMsg{}, channelError
	}

	//fmt.Println(string(ma), appkey, appsecret)

	v.Set("param_json", string(ma))
	sign := CreateJingdongSign(v, appsecret)
	v.Set("sign", sign)

	request, error := http.NewRequest("GET", URL + v.Encode(), nil)
	if error != nil {
		channelError := util.NewChannelRequestFailErrorError(error)
		return util.ResMsg{}, channelError
	}
	response, error := util.Client.Do(request)
	reqFunc(getReq)
	if error != nil {
		channelError := util.NewChannelRequestTimeoutError(error)
		return util.ResMsg{}, channelError
	}

	data, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if error != nil {
		channelError := util.NewChannelRequestNoError(error)
		return util.ResMsg{}, channelError
	}

	//fmt.Println("base2:", string(data))
	if response.StatusCode != 200 {
		code := response.StatusCode
		channelError := util.NewChannelRequestNoErrorWithText("状态码为:"+ strconv.Itoa(int(code)))
		return util.ResMsg{}, channelError
	}

	likeResult := LikeRes{}
	error = json.Unmarshal(data, &likeResult)
	if error != nil {
		channelError := util.NewChannelRequestNoError(error)
		return util.ResMsg{}, channelError
	}

	likeResultData := LikeResultData{}
	if len(likeResult.Jd_union_open_goods_material_query_response.Result) != 0 {
		//fmt.Println(likeResult.Jd_union_open_goods_material_query_response.Result)
		error = json.Unmarshal([]byte(likeResult.Jd_union_open_goods_material_query_response.Result), &likeResultData)
		if error != nil {
			channelError := util.NewChannelRequestNoError(error)
			return util.ResMsg{}, channelError
		}
	} else {
		channelError := util.NewChannelRequestNoErrorWithText("result长度为0")
		return util.ResMsg{}, channelError
	}
	if likeResultData.Code != 200 {
		channelError := util.NewChannelRequestNoErrorWithText("code错误")
		return util.ResMsg{}, channelError
	}

	if len(likeResultData.Data) == 0 {
		channelError := util.NewChannelRequestNoErrorWithText("result.data长度为0")
		return util.ResMsg{}, channelError
	}

	likeData := likeResultData.Data[0]

	if len(likeData.ImageInfo.ImageList) == 0 {
		channelError := util.NewChannelNoImageErrorWithText("图片链接长度为0")
		return util.ResMsg{}, channelError
	}

	postData := util.ResMsg{
		Title:                   likeData.ShopInfo.ShopName,
		Content:                 likeData.SkuName,
		ImageUrl:                likeData.ImageInfo.ImageList[0].Url,
		Uri:                     likeData.PromotionInfo.ClickURL,
	}

	if postData.ResponseDataIsEmpty(getReq.Adtype) {
		channelError := util.NewChannelRequestNoErrorWithText("数据不完整")
		return util.ResMsg{}, channelError
	}

	return postData, nil
}