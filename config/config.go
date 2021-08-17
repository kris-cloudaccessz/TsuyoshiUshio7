package config

import (
	"log"
	"strings"

	"github.com/go-ini/ini"
)

var conf *ini.File

func init() {
	var err error
	conf, err = ini.Load("config.ini")
	if err != nil {
		log.Println(err)
	}
}

// GetSetting str section.key.
func GetSetting(str string) string {
	strArr := strings.Split(str, ".")

	if len(strArr) < 2 {
		log.Println("section.key error")
		return ""
	}

	value := conf.Section(strArr[0]).Key(strArr[1]).String()

	return strings.Trim(value, " ")
}
