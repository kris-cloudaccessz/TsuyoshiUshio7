package main

import (
	"go-image/route"
	"go-image/server"
)

func main() {

	route.RegisterRoute()

	server.RunServer()
}
