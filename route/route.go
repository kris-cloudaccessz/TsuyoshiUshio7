package route

import (
	"go-image/controller"
	"go-image/server"
)

// RegisterRoute Register Route.
func RegisterRoute() {
	server.HandleFunc("/index/", controller.Index)
	server.HandleFunc("/upload", controller.Upload)
}
