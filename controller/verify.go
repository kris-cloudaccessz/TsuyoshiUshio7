package controller

import (
	"strconv"
	"strings"
)

func isMd5Str(str string) bool {
	return regexpURLParse.MatchString(str)
}

//IsType 判断类型是否允许上传
func IsType(typeStr string) bool {
	for _, v := range imageTypes {
		if strings.ToLower(v) == strings.ToLower(typeStr) {
			return true
		}
	}

	return false
}

func StringToInt(str string) int {
	if len(str) == 0 {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	if i <= 0 {
		return 0
	}

	return i
}
