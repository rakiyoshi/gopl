package memo_test

import (
	"fmt"
	memo "gopl/ch09/memo1"
	"io"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestMemoSync(t *testing.T) {
	m := memo.New(httpGetBody)

	// cache miss
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		since := time.Since(start)
		fmt.Printf("%s, %s, %d bytes\n", url, since, len(value.([]byte)))
	}
	fmt.Println()

	// cache hit
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		since := time.Since(start)
		fmt.Printf("%s, %s, %d bytes\n", url, since, len(value.([]byte)))
		if since.Nanoseconds() > 1000 {
			t.Fatalf("%d", since)
		}
	}
}

func TestMemoAsync(t *testing.T) {
	m := memo.New(httpGetBody)

	var wg sync.WaitGroup
	// cache miss
	for url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			since := time.Since(start)
			fmt.Printf("%s, %s, %d bytes\n", url, since, len(value.([]byte)))
			wg.Done()
		}(url)
	}

	// cache hit
	for url := range incomingURLs() {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			since := time.Since(start)
			fmt.Printf("%s, %s, %d bytes\n", url, since, len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func incomingURLs() <-chan string {
	out := make(chan string)
	urls := []string{
		"https://go.dev/",
		"https://go.dev/play/",
		"https://gopl.io/",
		"https://pkg.go.dev/",
	}
	go func() {
		for _, url := range urls {
			out <- url
		}
		close(out)
	}()
	return out
}
