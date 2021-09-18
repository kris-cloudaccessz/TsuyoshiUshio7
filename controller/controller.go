package controller

import (
	"fmt"
	"go-image/cache"
	"go-image/filehandler"
	"go-image/model"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Index 处理首页路径
func Index(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	if urlStr == "/favicon.ico" {
		return
	}

	var req = new(model.Goimg_req_t)

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

	model.ParamHandler(req, r)

	dirPath := imagePath + path
	sourceFilePath := dirPath + "/0_0"
	md5Str := parse.Path[1:]
	var cacheKey string

	//参数d=1时，直接下载文件。
	if req.Download == 1 {
		w.Header().Set("Content-Disposition", "attachment;filename="+md5Str+"."+req.Format)
	}

	if req.P == 1 {
		file, err := os.Open(sourceFilePath)
		if err != nil {
			log.Println(err)
			http.Error(w, "未找到文件", http.StatusNotFound)
			return
		}
		io.Copy(w, file)
		file.Close()
		return
	}

	//从缓存读取
	if cache.IsCache {
		cacheKey = fmt.Sprintf("%s:%d_%d_g%d_r%.f_p%d_x%d_y%d_q%d.%s", md5Str, req.Width, req.Height, req.Grayscale, req.Rotate, req.P, req.X, req.Y, req.Quality, req.Format)
		cacheValue := cache.Get(cacheKey)
		if *cacheValue != nil {
			w.Write(*cacheValue)
			return
		}
	}

	//从硬盘读取
	filePath := fmt.Sprintf("%s/%d_%d_g%d_r%.f_p%d_x%d_y%d_q%d.%s", dirPath, req.Width, req.Height, req.Grayscale, req.Rotate, req.P, req.X, req.Y, req.Quality, req.Format)
	file, err := os.Open(filePath)
	if err == nil {
		b, _ := ioutil.ReadAll(file)
		if cache.IsCache {
			cache.Set(cacheKey, b)
		}
		w.Write(b)
		file.Close()
		return
	}

	if _, err = os.Stat(sourceFilePath); err != nil && !os.IsExist(err) {
		http.Error(w, "文件不存在", http.StatusNotFound)
		return
	}

	b, err := filehandler.ResizeImage(sourceFilePath, req, filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cache.IsCache {
		cache.Set(cacheKey, *b)
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
		resp.Data.Size = uint(len(b))

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

		err = ioutil.WriteFile(dirPath+"0_0", b, 0660)
		if err != nil {
			resp.Success = false
			resp.Message = err.Error()
			response = append(response, resp)
			break
		}

		// err = filehandler.CompressionImage(b, dirPath+"0_0", 75, resp.Data)
		// if err != nil {
		// 	resp.Success = false
		// 	resp.Message = err.Error()
		// 	response = append(response, resp)
		// 	break
		// }

		resp.Success = true
		resp.Message = "OK"
		resp.Data.FileID = md5Str
		response = append(response, resp)
	}

	w.Write(model.ResponseJson(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	remoteIP := getRemoteIp(r)
	if !IsAllow(remoteIP) {
		http.Error(w, "禁止访问", http.StatusForbidden)
		return
	}

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
	md5Str := parse.Path[strings.LastIndex(parse.Path, "/")+1:]

	if !isMd5Str(md5Str) {
		http.NotFound(w, r)
		return
	}

	md5Path := SavePath(md5Str)

	if _, err = os.Stat(imagePath + md5Path); err == nil || os.IsExist(err) {
		err = os.RemoveAll(imagePath + md5Path)
		if err != nil {
			http.Error(w, "删除失败", http.StatusInternalServerError)
			return
		}
		cache.Del(md5Str)
		fmt.Fprintln(w, "ok")
		return
	}

	fmt.Fprintln(w, "文件不存在")
}

func responseError(w http.ResponseWriter, code int) {
	html := fmt.Sprintf("<html><body><h1>%s</h1></body></html>", model.StatusText(code))
	w.WriteHeader(code)
	fmt.Fprintln(w, html)
}

func getRemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
