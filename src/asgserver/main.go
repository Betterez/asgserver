package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("starting asg tester")
	server := &TestServer{}
	var version int64 = 1
	if os.Getenv("VERSION") != "" {
		version, _ = strconv.ParseInt(os.Getenv("VERSION"), 10, 32)
	}
	server.Version = version
	server.StartServer()
	log.Println("server running")
}
