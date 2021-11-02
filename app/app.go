package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mathantunes/echo/echo"
)

func main() {
	port := *flag.Int("port", 7, "echo tcp port")
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(err)
	}
	for err != nil {
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
