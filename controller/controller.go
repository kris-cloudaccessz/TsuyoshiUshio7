package controller

import (
	"fmt"
	"io"
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

	fmt.Println(md5Str)
	f := md5Str[1:4]

	fmt.Println(f)

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

	r.ParseMultipartForm(1024 << 12)

	formFile, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Get form file failed: %s \n", err)
		return
	}
	defer formFile.Close()

	destFile, err := os.Create("./upload/" + header.Filename)
	if err != nil {
		log.Printf("Create failed: %s \n", err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, formFile)
	if err != nil {
		log.Printf("Write file failed: %s \n", err)
		return
	}

	fmt.Fprint(w, "ok!")
	// r.ParseForm()
	// fmt.Println(r.PostForm["id"])

	// r.ParseMultipartForm(1024 << 12)
	// if r.MultipartForm != nil {
	// 	fmt.Fprintln(w, r.MultipartForm.Value["id"][0])
	// }

}
