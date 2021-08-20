package main

import (
	"go-image/route"
	"go-image/server"
	"log"
	"os"
)

func init() {
	logFilePath := "log.txt"

	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	return
}

func main() {
	route.RegisterRoute()
	server.RunServer()
}
