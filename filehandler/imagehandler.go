package filehandler

import (
	"errors"
	"go-image/model"
	"log"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func ResizeImage(imagePath string, w uint, h uint, outPath string) ([]byte, error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImage(imagePath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("读取图片错误")
	}

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var x, y int
	var w1, h1 uint

	if width < height {
		h1 = height * w / width
		w1 = w
		x = 0
		y = int((h1 - h) / 2)
	} else {
		w1 = h * width / height
		h1 = h
		x = int((w1 - w) / 2)
		y = 0
	}

	err = mw.ResizeImage(w1, h1, imagick.FILTER_LANCZOS)
	if err != nil {
		log.Println(err)
		return nil, errors.New("缩放图片错误")
	}

	err = mw.CropImage(w, h, x, y)
	if err != nil {
		log.Println(err)
		return nil, errors.New("裁切图片错误")
	}

	b := mw.GetImageBlob()

	err = mw.WriteImage(outPath)
	if err != nil {
		log.Println(err)
	}

	return b, nil
}

func CompressionImage(imageByte []byte, outPath string, quality uint, fileInfo *model.FileInfoModel) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImageBlob(imageByte)
	if err != nil {
		log.Println(err)
		return errors.New("压缩读取图片错误")
	}

	fileInfo.Mime = mw.GetImageFormat()

	if fileInfo.Mime == "JPEG" || fileInfo.Mime == "JPG" {
		err = mw.SetImageCompressionQuality(quality)
		if err != nil {
			log.Println(err)
			return errors.New("压缩图片错误")
		}
	}

	fileInfo.Size, err = mw.GetImageLength()
	if err != nil {
		log.Println(err)
		return errors.New("获取图片字节错误")
	}

	err = mw.WriteImage(outPath)
	if err != nil {
		log.Println(err)
		return errors.New("写入图片错误")
	}

	return nil
}
