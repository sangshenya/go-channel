package util

import (
	"math/rand"
	"strconv"
	"time"
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