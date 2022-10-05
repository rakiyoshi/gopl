package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch(url string) string {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-done:
			cancel()
		default:
		}
	}()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ""
	}

	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
	}
	return string(b)
}

var done = make(chan struct{})

func mirroredQuery(urls []string) string {
	resp := make(chan string)
	for _, url := range urls {
		go func(url string) {
			r := fetch(url)
			if r != "" {
				resp <- r
				done <- struct{}{}
			}
		}(url)
	}
	return <-resp
}

func main() {
	body := mirroredQuery(os.Args[1:])
	fmt.Println(body)
}
