package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string
type EnterOrLeave struct {
	name string
	cli  client
}

var (
	entering = make(chan EnterOrLeave)
	leaving  = make(chan EnterOrLeave)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[string]client)
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				cli <- msg
			}

		case ent := <-entering:
			clients[ent.name] = ent.cli
			ent.cli <- "Current Users:"
			for name := range clients {
				ent.cli <- name
			}
			ent.cli <- ""

		case lv := <-leaving:
			delete(clients, lv.name)
			close(lv.cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- EnterOrLeave{who, ch}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- EnterOrLeave{who, ch}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
