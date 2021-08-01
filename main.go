package main

import (
	"go-image/filehandler"
	"go-image/route"
	"go-image/server"
	"os"
)

func main() {

	f, _ := os.Open("config.ini")

	filehandler.GetFileHash(f)

	route.RegisterRoute()

	server.RunServer()
}
