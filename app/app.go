package main

import (
	"log"
	"net"

	"github.com/mathantunes/echo/echo"
)

func main() {
	ln, err := net.Listen("tcp", ":7")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go echo.Do(con)
	}
	err = ln.Close()
	if err != nil {
		log.Fatal(err)
	}
}
