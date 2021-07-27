package imghandler

import (
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

func SavePath() string {

}
