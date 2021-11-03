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

		text, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(text))
		if tcpcon, ok := conn.(*net.TCPConn); ok {
			tcpcon.CloseWrite()
		}
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, conn)
		if err != nil {
			log.Fatal(err)
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
