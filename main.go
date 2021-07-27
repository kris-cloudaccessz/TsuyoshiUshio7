package main

import (
	"goimg/route"
	"goimg/server"
)

func main() {

	route.RegisterRoute()

	server.RunServer()
}
