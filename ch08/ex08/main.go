package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
	input := bufio.NewScanner(c)

	done := make(chan struct{})
	go func() {
		for input.Scan() {
			go echo(c, input.Text(), 1*time.Second)
			done <- struct{}{}
		}
	}()

	ticker := time.NewTicker(1 * time.Second)
	var t int
	for {
		select {
		case <-ticker.C:
			if t > 10 {
				fmt.Fprintln(c, "connection timeout")
				c.Close()
				return
			}
			t++
		case <-done:
			t = 0
		}
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
