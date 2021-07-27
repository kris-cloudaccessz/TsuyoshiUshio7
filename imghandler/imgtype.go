package imghandler

import (
	"strings"

	"goimg/config"
)

var imageTypes []string

//IsType 判断类型是否允许上传
func IsType(typeStr string) bool {
	if len(imageTypes) == 0 {
		imageTypeStr := config.GetSetting("image.type")
		if len(imageTypeStr) == 0 {
			return false
		}
		imageTypes = strings.Split(imageTypeStr, ",")
	}

	for _, v := range imageTypes {
		if strings.ToLower(v) == strings.ToLower(typeStr) {
			return true
		}
	}

	return false
}
