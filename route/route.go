package route

import (
	"go-image/controller"
	"go-image/server"
)

// RegisterRoute Register Route.
func RegisterRoute() {
	server.HandleFunc("/", controller.Index)
	server.HandleFunc("/upload", controller.Upload)
}
