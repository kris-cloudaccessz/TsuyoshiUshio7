package controller

import (
	"fmt"
	"go-image/filehandler"
	"go-image/model"
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
	}

	md5Str := parse.Path

	// fmt.Println(md5Str)
	// f := md5Str[1:4]

	fmt.Println(md5Str)

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

	var response = new(model.ResponseModel)

	var fileInfos []*model.FileInfoModel

	for i := 0; i < len(files); i++ {

		fileInfo := new(model.FileInfoModel)
		file, err := files[i].Open()
		if err != nil {
			return
		}
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println(err)
		}

		filetype := filehandler.GetFileType(&b)

		if !IsType(filetype) {
			w.Write(response.ResponseJson(model.StatusImgIsType, false, nil))
			return
		}

		md5Str := filehandler.GetHash(&b)
		md5Path := SavePath(md5Str)

		file.Seek(0, 0)

		dirPath := imagePath + md5Path + "/"

		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
			w.Write(response.ResponseJson(model.StatusMkdir, false, nil))
			return
		}

		err = filehandler.CompressionImage(b, dirPath+"0_0", 75, fileInfo)
		if err != nil {
			log.Println(err)
			w.Write(response.ResponseJson(model.StatusImgCompression, false, nil))
			return
		}
		fileInfo.FileID = md5Str

		fileInfos = append(fileInfos, fileInfo)
	}

	w.Write(response.ResponseJson(model.StatusOK, true, fileInfos))
}
