package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	name, address string
}

func main() {
	var clocks []Clock
	for _, arg := range os.Args[1:] {
		sp := strings.Split(arg, "=")
		if len(sp) != 2 {
			fmt.Fprintf(os.Stderr, "invalid arg: %s", arg)
			os.Exit(1)
		}
		clocks = append(clocks, Clock{sp[0], sp[1]})
	}

	for _, clock := range clocks {
		fmt.Printf("%s%s", clock.name, strings.Repeat(" ", 9-len(clock.name)))
	}
	fmt.Println()

	for i, clock := range clocks {
		conn, err := net.Dial("tcp", clock.address)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn, i)
	}

	for {
		time.Sleep(time.Minute)
	}
}

func mustCopy(dst io.Writer, src io.Reader, n int) {
	sc := bufio.NewScanner(src)
	for sc.Scan() {
		if n == 0 {
			fmt.Fprintf(dst, "\r%s", sc.Text())
		} else {
			fmt.Fprintf(dst, "\r\033[%dC%s", 9*n, sc.Text())
		}
	}
}
