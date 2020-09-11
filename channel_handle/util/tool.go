package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Sha1(str string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(str))
	data := sha1Ctx.Sum(nil)
	return hex.EncodeToString(data[:])
}

func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	data := md5Ctx.Sum(nil)
	return hex.EncodeToString(data[:])
}

func GetRandom() string {
	unixStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	str := strconv.Itoa(rand.Intn(99))
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(unixStr + str))
	data := md5Ctx.Sum(nil)
	return hex.EncodeToString(data[:])
}

func ParamsEncode(paramsStr, adtype string) map[string]string {
	valueMap := map[string]string{}
	if len(paramsStr) == 0 {
		return valueMap
	}

	if strings.Contains(paramsStr, "&") {
		paramsStrArray := strings.Split(paramsStr, "&")
		if len(paramsStrArray) >= 2 {
			for _, paramItem := range paramsStrArray {
				if strings.Contains(paramItem, "=") {
					itemArray := strings.Split(paramItem, "=")
					if len(itemArray) >= 2 {
						valueMap[itemArray[0]] = itemArray[1]
					}
				}
			}
		}
	} else {
		if strings.Contains(paramsStr, "=") {
			itemArray := strings.Split(paramsStr, "=")
			if len(itemArray) >= 2 {
				valueMap[itemArray[0]] = itemArray[1]
			}
		}
	}

	_, hasW := valueMap["w"]
	_, hasH := valueMap["h"]
	wStr := ""
	hStr := ""
	switch adtype {
	case "startup":
		wStr = "640"
		hStr = "960"
	case "flow":
		wStr = "1280"
		hStr = "720"
	case "banner":
		wStr = "640"
		hStr = "100"
	}
	if !hasW && len(wStr) > 0 {
		valueMap["w"] = wStr
	}

	if !hasH && len(hStr) > 0 {
		valueMap["h"] = hStr
	}

	return valueMap
}

