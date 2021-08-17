package server

import (
	"go-image/config"
	"log"
	"net/http"
	"strconv"
	"time"
)

var serveMux *http.ServeMux = http.NewServeMux()

//HandleFunc Register route from HandleFunc.
func HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	serveMux.HandleFunc(pattern, handler)
}

//Handle Register route from Handle.
func Handle(pattern string, handler http.Handler) {
	serveMux.Handle(pattern, handler)
}

//RunServer start HTTP server.
func RunServer() {

	webPath := config.GetSetting("http.webPath")

	if len(webPath) != 0 {
		serveMux.Handle("/index/", http.StripPrefix("/index/", http.FileServer(http.Dir(webPath))))
	}

	readTimeout, err := strconv.Atoi(config.GetSetting("http.readTimeout"))
	if err != nil {
		readTimeout = 0
	}

	writeTimeout, err := strconv.Atoi(config.GetSetting("http.writeTimeout"))
	if err != nil {
		writeTimeout = 0
	}

	serv := &http.Server{
		Addr:         config.GetSetting("http.addr"),
		Handler:      serveMux,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,  // 读超时
		WriteTimeout: time.Duration(writeTimeout) * time.Second, // 写超时
		//ReadHeaderTimeout: 5 * time.Second,
		//IdleTimeout:       5 * time.Second,
	}

	err = serv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
