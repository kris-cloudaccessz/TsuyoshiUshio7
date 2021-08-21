package controller

import (
	"fmt"
	"go-image/cache"
	"go-image/filehandler"
	"go-image/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Index 处理首页路径
func Index(w http.ResponseWriter, r *http.Request) {

	urlStr := r.URL.String()

	if urlStr == "/favicon.ico" {
		return
	}

	parse, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	path := ParseUrlPath(parse.Path[1:])
	if path == "" {
		http.NotFound(w, r)
		return
	}

	grayscale := StringToBool(r.FormValue("g"))

	var g int8
	if grayscale {
		g = 1
	} else {
		g = 0
	}

	rotate := StringToFloat64(r.FormValue("r"))
	width := StringToInt(r.FormValue("w"))
	height := StringToInt(r.FormValue("h"))

	dirPath := imagePath + path

	if width == 0 && height == 0 {
		//优先从缓存中读取
		key := dirPath + "/0_0"
		cacheValue := cache.Get(key)
		if cacheValue != nil {
			w.Write(*cacheValue)
			return
		}

		//未找到缓存从磁盘读取
		file, err := os.Open(key)
		if err != nil {
			log.Println(err)
			http.Error(w, "服务器内部错误", http.StatusInternalServerError)
			return
		}

		b, err := ioutil.ReadAll(file)
		cache.RedisClient.Set(key, b, time.Second*600)
		w.Write(b)
		file.Close()
		return
	}

	filePath := fmt.Sprintf("%s/%d_%d_g%d_r%.0f", dirPath, width, height, g, rotate)

	//从缓存读取

	//从硬盘读取
	file, err := os.Open(filePath)
	if err == nil {
		io.Copy(w, file)
		file.Close()
		return
	}

	b, err := filehandler.ResizeImage(dirPath+"/0_0", uint(width), uint(height), rotate, grayscale, filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, "文件处理失败", http.StatusInternalServerError)
		return
	}

	w.Write(*b)
}

//Uploads upload files function.
func Upload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1024 << 14)

	files := r.MultipartForm.File["files"]

	var response []*model.ResponseModel

	for i := 0; i < len(files); i++ {
		resp := model.NewResponseModel()
		// fileInfo := new(model.FileInfoModel)
		// resp.Data = fileInfo

		file, err := files[i].Open()
		if err != nil {
			resp.Success = false
			resp.Message = "读取file数据失败"
			response = append(response, resp)
			break
		}
		defer file.Close()

		resp.Data.FileName = files[i].Filename

		b, err := ioutil.ReadAll(file)
		if err != nil {
			resp.Success = false
			resp.Message = "读取file数据失败"
			response = append(response, resp)
			break
		}

		resp.Data.Mime = http.DetectContentType(b)

		if !IsType(resp.Data.Mime) {
			resp.Success = false
			resp.Message = "图片类型不符合"
			response = append(response, resp)
			break
		}

		md5Str := filehandler.GetHash(&b)
		md5Path := SavePath(md5Str)

		file.Seek(0, 0)

		dirPath := imagePath + md5Path + "/"

		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			resp.Success = false
			resp.Message = "创建目录失败"
			response = append(response, resp)
			break
		}

		err = filehandler.CompressionImage(b, dirPath+"0_0", 75, resp.Data)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			response = append(response, resp)
			break
		}

		resp.Success = true
		resp.Message = "OK"
		resp.Data.FileID = md5Str
		response = append(response, resp)
	}

	w.Write(model.ResponseJson(response))
}

func responseError(w http.ResponseWriter, code int) {
	html := fmt.Sprintf("<html><body><h1>%s</h1></body></html>", model.StatusText(code))
	w.WriteHeader(code)
	fmt.Fprintln(w, html)
}
