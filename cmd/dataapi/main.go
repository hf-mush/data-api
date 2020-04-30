package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/comail/colog"
)

var exitCode = 0

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc

	colog.Register()
}

func main() {
	port, err := strconv.ParseInt(os.Getenv("API_PORT"), 10, 64)
	if err != nil {
		log.Println("error: [port] " + err.Error())
		gofmtMain()
		os.Exit(exitCode)
	}

	flag.Int64Var(&port, "port", port, "Listen port of HTTP Server")

	err = RunServer(port)
	if err != nil {
		log.Println("error: [http] " + err.Error())
		gofmtMain()
		os.Exit(exitCode)
	}
}

func gofmtMain() {
	exitCode = 2
	return
}
