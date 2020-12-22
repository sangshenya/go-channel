package jdlm

// 猜你喜欢
type LikeRes struct {
	Jd_union_open_goods_material_query_response LikeResult
}

type LikeResult struct {
	Result string
}

type LikeResultData struct {
	Code int
	Message string
	Data []LikeMaterial
	TotalCount int
}

type LikeMaterial struct {
	ResourceInfo resourceInfo
	SkuName string
	ImageInfo imageInfo
	ShopInfo shopInfo
	PromotionInfo promotionInfo
}

type pinGouInfo struct {
	PingouTmCount string
	PingouEndTime string
	PingouPrice string
	PingouStartTime string

}

type categoryInfo struct {
	Cid3 string
	Cid2Name string
	Cid2 string
	Cid3Name string
	Cid1Name string
	Cid1 string
}

type resourceInfo struct {
	EliteName string

}

type seckillInfo struct {
	SeckillEndTime string
	SeckillPrice string
	SeckillOriPrice string
	SeckillStartTime string

}

type bookInfo struct {
	Isbn string

}

type videoInfo struct {

}

type imageInfo struct {
	ImageList []urlInfo
	WhiteImage string
}

type imageList struct {
	UrlInfo urlInfo
}

type urlInfo struct {
	Url string
}

type priceInfo struct {
	Price string
	LowestCouponPrice string
	LowestPriceType string
	LowestPrice string

}

type shopInfo struct {
	ShopName string
}

type commissionInfo struct {
	CommissionShare float64
	CouponCommission float64
	PlusCommissionShare float64
	Commission float64
}

type promotionInfo struct {
	ClickURL string
}

type couponInfo struct {
	CouponList []couponList
}

type couponList struct {
	Coupon coupon
}

type coupon struct {
	GetEndTime string
	GetStartTime string
	Quota string
	PlatformType string
	UseEndTime string
	UseStartTime string
	BindType string
	IsBest string
	Link string
	Discount string
}