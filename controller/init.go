package controller

import (
	"go-image/config"
	"log"
	"regexp"
	"strings"
)

var regexpURLParse *regexp.Regexp
var imageTypes []string
var imagePath string
var adminIPs []string

func init() {
	var err error

	imagePath = config.GetSetting("image.path")

	if len(imagePath) == 0 {
		imagePath = "image/"
	} else {
		if imagePath[len(imagePath):] != "/" {
			imagePath = imagePath + "/"
		}
	}

	regexpURLParse, err = regexp.Compile("[a-z0-9]{32}")
	if err != nil {
		log.Println("regexpUrlParse:", err)
	}

	imageTypes = strings.Split(config.GetSetting("image.type"), ",")
	adminIPs = strings.Split(config.GetSetting("server.admin_ips"), ",")
}
