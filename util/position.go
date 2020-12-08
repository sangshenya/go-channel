package util

import (
	"math/rand"
	"strconv"
	"time"
)

/*
渠道的宏替换
	__TS__:当前时间,单位:毫秒
	__TS_S__:当前时间，单位秒

	__DOWN_X__:相对于广告位的按下x坐标
	__DOWN_Y__:相对于广告位的按下y坐标
	__UP_X__:相对于广告位的抬起x坐标
	__UP_Y__:相对于广告位的抬起y坐标

	__RE_DOWN_X__:相对于屏幕的按下x坐标
	__RE_DOWN_Y__:相对于屏幕的按下y坐标
	__RE_UP_X__:相对于屏幕的抬起x坐标
	__RE_UP_Y__:相对于屏幕的抬起y坐标

	__WIDTH__:在手机上真实展示的宽度，与手机屏幕宽度相关
	__HEIGHT__:在手机上真实展示的高度，与手机屏幕宽度、广告类型相关

	__CLICK_ID__:广点通下载id

	其中请求宽高应该当在请求是进行替换

*/

const(
	TS = "EVENT_TIME"
	TSS = "EVENT_TIME"

	DX = "IT_CLK_PNT_DOWN_X"
	DY = "IT_CLK_PNT_DOWN_Y"
	UX = "IT_CLK_PNT_UP_X"
	UY = "IT_CLK_PNT_UP_Y"

	RDX = "IT_CLK_PNT_DOWN_SX"
	RDY = "IT_CLK_PNT_DOWN_SY"
	RUX = "IT_CLK_PNT_UP_SX"
	RUY = "IT_CLK_PNT_UP_SY"

	CLKID = "__CLICK_ID__"
)

// 获取各事件的时间戳 0:当前时间；1：展示时间，2：点击时间，3：跳转成功时间
func GetTime() [4]string {
	nowTime :=  time.Now().Unix() * 1000

	index := (rand.Int() % 5) + 3

	resTime := nowTime + int64(index * 100)

	disTime := resTime + int64(100)

	clickIndex := rand.Int() % 25 + 10

	clickTime := disTime + int64(clickIndex * 100)

	schemeIndex := rand.Int() % 5 + 5

	schemeTime := clickTime + int64(schemeIndex * 100)

	return [4]string{
		strconv.Itoa(int(nowTime)),
		strconv.Itoa(int(disTime)),
		strconv.Itoa(int(clickTime)),
		strconv.Itoa(int(schemeTime)),
	}
}

// 根据屏幕宽度生成相对于广告位的坐标和相对屏幕坐标和宽高
// 小图信息流按照信息流设置
func CreateAbScreenWHPos(width, height, adtype string) [10]string {
	if adtype == "banner" {
		adtype = "1"
	} else if adtype == "flow" {
		adtype = "2"
	} else if adtype == "startup" {
		adtype = "3"
	}
	w, err := strconv.Atoi(width)
	if err != nil {
		w = 0
	}
	sh, err := strconv.Atoi(height)
	if err != nil {
		sh = 0
	}
	h := 0
	switch adtype {
	case "1":
		h = w * 100 / 640
	case "2":
		h = w * 720 / 1280 + 100
	case "3":
		h = sh
	}
	pos := [10]string{}
	aw := strconv.Itoa(createN(w))
	// 宽度不变
	pos[0] = aw
	pos[2] = aw

	pos[4] = aw
	pos[6] = aw

	// 获取高度
	// 根据广告位高度获取随机高度,在广告位中
	th := createN(h)
	ah := strconv.Itoa(th)
	// 相对于广告位高度确定
	pos[1] = ah
	pos[3] = ah
	// 相对于屏幕高度
	if adtype == "3" {
		// 开屏是全屏广告，相对于广告位高度和相对于屏幕相同
		pos[5] = ah
		pos[7] = ah
	} else if adtype == "1" {
		// banner部分是固定在底部
		index := rand.Int() % 5
		if index > 3 {
			// 固定底部
			ah := strconv.Itoa(sh - h - 59 + th)
			pos[5] = ah
			pos[7] = ah
		} else {
			// 不固定底部
			ah := strconv.Itoa(createN(sh - h) + th)
			pos[5] = ah
			pos[7] = ah
		}

	} else {
		ah := strconv.Itoa(createN(sh - h) + th)
		pos[5] = ah
		pos[7] = ah
	}

	pos[8] = width
	pos[9] = strconv.Itoa(h)
	return pos
}

// 生成随机点（0.2~0.8）
func createN(n int) int {
	maxN := n * 50 / 100
	if maxN < 1 {
		maxN = 200
	}
	rN := rand.Intn(maxN)
	return n*30/100 + rN
}


// 根据屏幕宽度生成相对于广告位的坐标和相对屏幕坐标
// 小图信息流按照信息流设置
func CreateAbScreenPosWHWithOs(width, height, adtype, os string) [10]string {
	if adtype == "banner" {
		adtype = "1"
	} else if adtype == "flow" {
		adtype = "2"
	} else if adtype == "startup" {
		adtype = "3"
	}
	w := 0
	sh := 0
	if len(width) == 0 {
		w = 1080
	}
	if len(height) == 0 {
		sh = w * 2 / 3
	}
	indexw := 0
	indexh := 0
	if os == "1" {
		indexw = rand.Int() % 12 - 6
		indexh = rand.Int() % 8 - 4
	}

	w, err := strconv.Atoi(width)
	if err != nil {
		w = 1080
	}
	sh, err = strconv.Atoi(height)
	if err != nil {
		sh = w * 2 / 3
	}
	h := 0
	switch adtype {
	case "1":
		h = w * 100 / 640
	case "2":
		h = w * 360 / 640
	case "3":
		h = sh
	}
	pos := [10]string{}
	cw := createN(w)
	aw := strconv.Itoa(cw)
	aw1 := strconv.Itoa(cw + indexw)
	// 宽度不变
	pos[0] = aw
	pos[2] = aw1

	pos[4] = aw
	pos[6] = aw1

	// 获取高度
	// 根据广告位高度获取随机高度,在广告位中
	th := createN(h)
	ah := strconv.Itoa(th)
	ah1 := strconv.Itoa(th + indexh)
	// 相对于广告位高度确定
	pos[1] = ah
	pos[3] = ah1
	// 相对于屏幕高度
	if adtype == "3" {
		// 开屏是全屏广告，相对于广告位高度和相对于屏幕相同
		pos[5] = ah
		pos[7] = ah1
	} else if adtype == "1" {
		// banner部分是固定在底部
		index := rand.Int() % 5
		if index > 3 {
			// 固定底部
			ah = strconv.Itoa(sh - h - 59 + th)
			ah1 = strconv.Itoa(sh - h - 59 + th + indexh)
			pos[5] = ah
			pos[7] = ah1
		} else {
			// 不固定底部
			ah = strconv.Itoa(createN(sh - h) + th)
			ah1 = strconv.Itoa(createN(sh - h) + th + indexh)
			pos[5] = ah
			pos[7] = ah1
		}

	} else {
		ah = strconv.Itoa(createN(sh - h) + th)
		ah1 = strconv.Itoa(createN(sh - h) + th + indexh)
		pos[5] = ah
		pos[7] = ah1
	}

	pos[8] = width
	pos[9] = strconv.Itoa(h)
	return pos
}