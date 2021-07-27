package imghandler

import (
	"strings"
)

//IsType 判断类型是否允许上传
func IsType(typeStr string) bool {
	for _, v := range imageTypes {
		if strings.ToLower(v) == strings.ToLower(typeStr) {
			return true
		}
	}

	return false
}
