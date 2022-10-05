package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8021")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	sc := bufio.NewScanner(c)
	defer c.Close()

	_, err := io.WriteString(c, "220 Service ready for new user.\n")
	if err != nil {
		log.Print(err)
		return
	}

	for {
		sc.Scan()
		fmt.Println(sc.Text())
		s := strings.SplitN(sc.Text(), " ", 2)

		var command, arg string
		if len(s) > 0 {
			command = s[0]
		}
		if len(s) > 1 {
			arg = s[1]
		}
		switch command {
		case "USER":
			_, err = io.WriteString(c, fmt.Sprintf("331 Hi, %s. Need password.\n", arg))
			if err != nil {
				return
			}
		case "PASS":
			_, err = io.WriteString(c, "230 User logged in, proceed.\n")
			if err != nil {
				log.Print(err)
				return
			}
		case "SYST":
			_, err = io.WriteString(c, "215 gopl FTP server.\n")
			if err != nil {
				log.Print(err)
				return
			}
		case "PASV":
			_, err = io.WriteString(c, fmt.Sprintf("227 Entering Passive Mode (127,0,0,1,%d,%d).\n", 8020/256, 8020%256))
			if err != nil {
				log.Print(err)
				return
			}
		default:
			_, err = io.WriteString(c, fmt.Sprintf("502 Command not implemented: %s\n", sc.Text()))
			if err != nil {
				log.Print(err)
				return
			}
		}
	}
}
