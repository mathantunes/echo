package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mathantunes/echo/echo"
)

func main() {
	port := flag.Int("port", 7, "echo tcp port")
	flag.Parse()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port %v", *port)
	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		if err := echo.Do(con); err != nil {
			log.Println(err)
		}
	}
}
