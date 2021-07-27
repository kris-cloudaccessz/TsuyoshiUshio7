package filehandler

import (
	"go-image/config"
	"log"
	"regexp"
	"strings"
)

var regexpURLParse *regexp.Regexp
var imageSavePath string
var fileSavePath string
var imageTypes []string

func init() {
	imageSavePath = config.GetSetting("image.path")
	fileSavePath = config.GetSetting("file.path")
	imageTypes = strings.Split(config.GetSetting("image.type"), ",")

	var err error
	regexpURLParse, err = regexp.Compile("[a-z0-9]{32}")
	if err != nil {
		log.Fatalln("regexpUrlParse:", err)
	}
}
