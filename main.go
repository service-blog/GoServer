package main

import (
	"os"

	"blog/service"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "7777"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for http to listen")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}