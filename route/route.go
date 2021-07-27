package route

import (
	"goimg/controller"
	"goimg/server"
)

// RegisterRoute Register Route.
func RegisterRoute() {
	server.HandleFunc("/index/", controller.Index)
	server.HandleFunc("/upload", controller.Upload)
}
