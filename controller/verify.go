package controller

import (
	"strings"
)

func isMd5Str(str string) bool {
	return regexpURLParse.MatchString(str)
}

//IsType 判断类型是否允许上传
func IsType(typeStr string) bool {
	for _, v := range imageTypes {
		if strings.Contains(typeStr, strings.ToLower(v)) {
			return true
		}
	}

	return false
}

func IsAllow(ip string) bool {
	for _, v := range adminIPs {
		if v == ip {
			return true
		}
	}

	return false
}
