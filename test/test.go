package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":7")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	_, err = conn.Write([]byte(text))
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
