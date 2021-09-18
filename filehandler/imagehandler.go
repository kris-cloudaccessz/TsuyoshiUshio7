package filehandler

import (
	"errors"
	"go-image/model"
	"log"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func ResizeImage(imagePath string, req *model.Goimg_req_t, outPath string) (*[]byte, error) {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImage(imagePath)
	if err != nil {
		log.Println(err)
		return nil, errors.New("读取图片错误")
	}

	imageType := mw.GetImageColorspace()
	if imageType == imagick.COLORSPACE_CMYK {
		return nil, errors.New("该图片为CMYK，暂不支持处理")
	}

	width := mw.GetImageWidth()
	height := mw.GetImageHeight()

	var x, y int
	var w1, h1 uint

	if req.X >= 0 && req.Y >= 0 {
		x = req.X
		y = req.Y
		if req.Width != 0 && req.Height != 0 {
			if req.Width >= width {
				w1 = width
			} else {
				w1 = req.Width
			}

			if req.Height >= height {
				h1 = height
			} else {
				h1 = req.Height
			}

			err = mw.CropImage(w1, h1, x, y)
			if err != nil {
				log.Println(err)
				return nil, errors.New("裁切图片错误")
			}
		}
	} else {
		if req.Width == 0 && req.Height == 0 {
			w1 = width
			h1 = height
		} else if req.Width == 0 || req.Height == 0 {
			if req.Width == 0 {
				w1 = req.Height * width / height
				h1 = req.Height
			}

			if req.Height == 0 {
				h1 = height * req.Width / width
				w1 = req.Width
			}

			err = mw.ResizeImage(w1, h1, imagick.FILTER_LANCZOS)
			if err != nil {
				log.Println(err)
				return nil, errors.New("缩放图片错误")
			}

		} else {
			if width < height {
				h1 = height * req.Width / width
				w1 = req.Width
				x = 0
				y = int((h1 - req.Height) / 2)
			} else {
				w1 = req.Height * width / height
				h1 = req.Height
				x = int((w1 - req.Width) / 2)
				y = 0
			}

			err = mw.ResizeImage(w1, h1, imagick.FILTER_LANCZOS)
			if err != nil {
				log.Println(err)
				return nil, errors.New("缩放图片错误")
			}

			err = mw.CropImage(req.Width, req.Height, x, y)
			if err != nil {
				log.Println(err)
				return nil, errors.New("裁切图片错误")
			}
		}
	}

	mw.StripImage()

	err = mw.SetImageFormat(req.Format)
	if err != nil {
		log.Println(err)
		return nil, errors.New("设置图片格式错误")
	}

	mw.SetImageCompression(imagick.COMPRESSION_JPEG)
	err = mw.SetImageCompressionQuality(req.Quality)
	if err != nil {
		log.Println(err)
		return nil, errors.New("压缩图片错误")
	}

	if req.Grayscale == 1 {
		//设置图片颜色灰度
		err = mw.SetImageType(imagick.IMAGE_TYPE_GRAYSCALE)
		if err != nil {
			log.Println(err)
		}
	}

	if req.Rotate != 0 {
		pw := imagick.NewPixelWand()
		defer pw.Destroy()
		pw.SetColor("white")
		err = mw.RotateImage(pw, req.Rotate)
		if err != nil {
			log.Println(err)
			return nil, errors.New("旋转图片错误")
		}
	}

	b := mw.GetImageBlob()

	err = mw.WriteImage(outPath)
	if err != nil {
		log.Println(err)
	}

	return &b, nil
}

func CompressionImage(imageByte []byte, outPath string, quality uint, fileInfo *model.FileInfoModel) error {
	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	err := mw.ReadImageBlob(imageByte)
	if err != nil {
		log.Println(err)
		return errors.New("压缩读取图片错误")
	}

	if mw.GetImageFormat() == "JPEG" || mw.GetImageFormat() == "JPG" {
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
