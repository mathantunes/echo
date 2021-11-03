package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	port := flag.Int("port", 7, "echo tcp port")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf(":%v", *port))
		if err != nil {
			log.Fatal(err)
		}

		text, err := reader.ReadBytes('\n')
		fatalIfError(err)

		_, err = conn.Write(text)
		fatalIfError(err)

		if tcpcon, ok := conn.(*net.TCPConn); ok {
			tcpcon.CloseWrite()
		}
		_, err = io.Copy(os.Stdout, conn)
		fatalIfError(err)

		err = conn.Close()
		fatalIfError(err)
	}
}

func fatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
