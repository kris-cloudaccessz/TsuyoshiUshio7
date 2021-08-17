package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// SortPath 排序
func SortPath(str []byte) string {
	strLen := len(str)

	for i := 0; i < strLen; i++ {
		for j := 1 + i; j < strLen; j++ {
			if str[i] > str[j] {
				str[i], str[j] = str[j], str[i]
			}
		}
	}

	// 对 byte 依次组成数字符串
	var ret = strings.Builder{}

	for i := 0; i < strLen; i++ {
		ret.WriteString(strconv.Itoa(int(str[i])))
	}

	return ret.String()
}

func SavePath(md5Str string) string {
	firstDir, err := strconv.ParseUint(md5Str[:3], 16, 32)
	if err != nil {
		log.Println(err)
		return ""
	}

	secondDir, err := strconv.ParseUint(md5Str[3:6], 16, 32)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf("%d/%d/%s", firstDir/4, secondDir/4, md5Str)
}

func ParseUrlPath(urlPath string) string {
	if !isMd5Str(urlPath) {
		return ""
	}

	return SavePath(urlPath)
}
