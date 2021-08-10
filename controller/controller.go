package controller

import (
	"fmt"
	"go-image/filehandler"
	"go-image/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

// Index 处理首页路径
func Index(w http.ResponseWriter, r *http.Request) {

	urlStr := r.URL.String()

	if urlStr == "/favicon.ico" {
		return
	}

	parse, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	path := ParseUrlPath(parse.Path[1:])
	if path == "" {
		return
	}

	width := StringToInt(r.FormValue("w"))
	height := StringToInt(r.FormValue("h"))

	dirPath := imagePath + path

	if width == 0 || height == 0 {
		file, err := os.Open(dirPath + "/0_0")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		io.Copy(w, file)
		return
	}

	filePath := fmt.Sprintf("%s/%d_%d", dirPath, width, height)

	file, err := os.Open(filePath)
	if err == nil {
		io.Copy(w, file)
		return
	}
	defer file.Close()

	b, err := filehandler.ResizeImage(dirPath+"/0_0", uint(width), uint(height), filePath)
	if err != nil {

	}

	w.Write(b)
	return

	// fmt.Println(md5Str)
	// f := md5Str[1:4]

	// n, err := strconv.ParseUint(f, 16, 32)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(n)

	// s := imghandler.SortPath([]byte(md5Str[:5]))

	// fmt.Fprintln(w, s)

	// fmt.Fprintln(w, r.URL.String())
}

//Upload upload file function.
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

		filetype := filehandler.GetFileType(&b)
		resp.Data.Mime = filetype

		if !IsType(filetype) {
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

		resp.Data.FileID = md5Str
		response = append(response, resp)
	}

	w.Write(model.ResponseJson(response))
}
