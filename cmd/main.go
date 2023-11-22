package main

import (
	"log"
	"strings"

	"github.com/hosseinmirzapur/rthnk/api/server"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	err := server.RunServer()
	handleErrLog(err)
}

func handleErrLog(err error) {
	if err != nil {
		log.Println(strings.Repeat("=", 10), "ERROR LOG", strings.Repeat("=", 10))
		log.Fatal(err.Error())
	}
}
