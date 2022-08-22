// FetchAll fetches contents of URL parallelly and prints time and size
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading data: %v", err)
		return
	}
	nbytes := len(data)

	resp.Body.Close()

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

	curDir, err := os.Getwd()
	if err != nil {
		ch <- fmt.Sprintf("failed to get current directory: %v", err)
	}
	filename := curDir + "/ex10" + resp.Request.URL.Path
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v", filename, err)
	}
}
