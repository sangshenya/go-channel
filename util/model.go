package util


type ReqMsg struct {
	Screenwidth  	string 					// 屏幕宽度
	Screenheight 	string 					// 屏幕高度
	Adtype 			string 					// 客户端请求的广告类型
	Os           	string 					// 设备操作系统，1: Android，2：iOS，3：Windows，4：其它
	Osversion    	string 					// 系统版本号
	Imsi         	string 					// 网络型号1：中国移动、2：中国联通、3：中国电信
	Ip           	string 					// ip地址
	Mac          	string 					// 设备 MAC 地址（大写，保留冒号分隔符）
	Network      	string 					// 网络类型
	Device       	string 					// 设备型号
	Androidid    	string 					// 安卓设备广告标识符
	Idfa         	string 					// 苹果设备广告标识符
	Idfv 		 	string 					//
	Openudid     	string 					// iOS 版本 6 以下的操作系统提供 OpenUDID
	Appversion   	string 					// app版本
	Ua           	string 					// 用户代理
	Version      	string 					// app版本
	Sd           	string 					// 分辨率或密度比（/160）
	Machine     	string 					// 机器编码
	Vendor       	string 					// 设备生产商
	Imei         	string 					// 设备 IMEI 号
	Model        	string 					// 设备型号
	Lat          	string 					// 经度
	Lng          	string 					// 纬度
	Sdkid        	string 					// 渠道号
	Imsi_long    	string 					// 新版的imsi
	AndroidApiLevel string 					// 安卓 api level
	Carrier         string 					// 运营商信息，前五位
	Orientation     string 					// 设备当前屏幕状态，0：竖屏，1：横屏
	Channel         string 					// 当前渠道号
	Oaid			string 					//
	SdkVersion		string 					// sdkVersion
	Package   		string 					// 媒体包名
	Wifiname		string 					// wifi名称
	Wifimac			string 					// wifi mac地址
	Romversion		string 					// rom版本号（安卓）
	Comptime		string 					// 系统编译时间
	ChannelReq		ChannelMsg 				// 渠道相关的请求参数
}

type ChannelMsg struct {
	Appid 			string					// 渠道所需的appid
	Adid 			string					// 渠道所需的adid
	Token 			string					// 渠道所需的token
	Adtype 			string 					// 渠道所需的adtype
	Pkg 			string					// 渠道所需的媒体包名
	Appname			string					// 渠道所需的媒体名称
	Appurl			string					// 渠道所需的媒体下载地址
	Params 			string 					// 渠道所需的其他请求参数
	Loginfo			string 					// 统计使用到的参数
	Errorinfo		string 					// 错误信息
	Ext 			map[string]interface{} 	// 拓展字段，防止临时需求
}

type ResMsg struct {
	Id            			string   	`json:"id"`                     				// "0"
	Weight        			int      	`json:"weight"`                 				// 0
	State         			int      	`json:"state"`                  				// 1
	Title         			string   	`json:"title"`                   				// 标题
	Content       			string   	`json:"content"`                 				// 描述
	ImageUrl      			string   	`json:"imageUrl"`               				// 图片url
	Uri           			string   	`json:"uri"`                     				// 落地页
	Displayreport 			[]string 	`json:"displayreport"`           				// 展示上报地址
	Clickreport   			[]string 	`json:"clickreport"`             				// 点击上报地址
	StartDownload 			[]string 	`json:"startDownload,omitempty"` 				// 开始下载
	Downloaded    			[]string 	`json:"downloaded,omitempty"`    				// 下载完成
	Installed     			[]string 	`json:"installed,omitempty"`     				// 安装完成
	Open          			[]string 	`json:"open,omitempty"`          				// 打开app
	Json          			bool     	`json:"json"`                    				// 下载类返回json格式标示
	Scheme        			string   	`json:"scheme"`                  				// scheme地址
	Schemereport  			[]string 	`json:"schemereport,omitempty"`  				// scheme上报地址
	Pkg           			string   	`json:"pkg"`                     				// 包名
	ImageList     			[]string 	`json:"imageList"`               				// 图组
	// 视频Video
	VideoUrl 				string 		`json:"videoUrl,omitempty"`						//视频播放链接
	VideoEndHtml			string 		`json:"videoEndHtml,omitempty"`					//视频播放完成最后展示的卡片html
	VideoLoadedReport		[]string 	`json:"videoLoadedReport,omitempty"` 			//视频加载完成上报，判断视频资源可以播放后上报
	VideoStartReport 		[]string 	`json:"videoStartReport,omitempty"`				//播放开始上报
	VideoQuarterReport 		[]string 	`json:"videoQuarterReport,omitempty"`			//播放四分之一上报
	VideoHalfReport			[]string 	`json:"videoHalfReport,omitempty"`				//播放二分之一上报
	VideoThreeQuarterReport []string 	`json:"videoThreeQuarterReport,omitempty"` 		//播放四分之三上报
	VideoFinishReport		[]string 	`json:"videoFinishReport,omitempty"`			//播放完成上报
	VideoSkipReport 		[]string 	`json:"videoSkipReport,omitempty"`				//视频跳过上报
	VideoCacheErrorReport   []string 	`json:"videoCacheErrorReport,omitempty"`		//视频缓存错误上报
	VideoPlayingErrorReport []string 	`json:"videoPlayingErrorReport,omitempty"`		//视频播放错误上报
	VideoCloseReport 		[]string 	`json:"videoCloseReport,omitempty"`				//视频关闭上报
	VideoMuteReport         []string 	`json:"videoMuteReport,omitempty"`				//视频静音上报
	VideoUnmuteReport 		[]string 	`json:"videoUnmuteReport,omitempty"`			//视频非静音播放
	VideoCallBackReport		[]string 	`json:"videoCallBackReport,omitempty"`			//视频激励验证回调，此字段只有在使用服务器激励回调时会 返回。
	Duration   				int 		`json:"duration,omitempty"`						//视频时长，单位为：s
	VideoSize 				int64 		`json:"videoSize,omitempty"`					//视频大小

}

type ReqFailFunc func(msg *ReqMsg)

type ReqFunc func(msg *ReqMsg)

type ReqNoFunc func(msg *ReqMsg)

type ReqTimeoutFunc func(msg *ReqMsg)

type ReqNoimgFunc func(msg *ReqMsg)

type ReqNourlFunc func(msg *ReqMsg)


type BaseObj struct {
	Req ReqMsg
	NoFunc ReqNoFunc
	TimeoutFunc ReqTimeoutFunc
	NoimgFunc ReqNoimgFunc
	NourlFunc ReqNourlFunc

}

func (res *ResMsg) ResponseDataIsEmpty(adtype string) bool {
	if adtype == "noticead" || adtype == "innernoticead" {
		if (res.Uri == "" && res.Scheme == "") || res.ImageUrl == "" {
			return true
		}
	} else if adtype == "vod" {
		if res.ImageUrl == "" || res.Uri == "" || len(res.VideoUrl) == 0 {
			return true
		}
	} else {
		if res.ImageUrl == "" || res.Uri == "" {
			return true
		}
	}

	return false
}