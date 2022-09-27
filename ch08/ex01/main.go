package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var portFlag int

func init() {
	flag.IntVar(&portFlag, "port", 8000, "port number")
}

func main() {
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", portFlag)
	fmt.Println(address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
