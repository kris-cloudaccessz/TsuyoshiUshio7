package controller

import (
	"bufio"
	"fmt"
	"go-image/filehandler"
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

	for i := 0; i < len(files); i++ {
		file, err := files[i].Open()
		if err != nil {
			return
		}
		defer file.Close()

		bufferFile := bufio.NewReader(file)

		md5Str := filehandler.GetFileHash(bufferFile)
		md5Path := filehandler.SavePath(md5Str)

		file.Seek(0, 0)

		dirPath := "upload/" + md5Path + "/"

		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}

		b, err := ioutil.ReadAll(bufferFile)
		if err != nil {
			log.Println(err)
		}

		err = filehandler.CompressionImage(b, dirPath+"0_0", 75)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Fprint(w, "ok!")
}
