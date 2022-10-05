package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()

	go func() {
		_, err := os.Stdin.Read(make([]byte, 1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex10: %v\n", err)
		}
		close(done)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case <-done:
					for range unseenLinks {
					}
					return
				case link, ok := <-unseenLinks:
					if !ok {
						return
					}
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if cancelled() {
				return
			}
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-tokens }()
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// Extract requests HTTP GET to specified URL, parses the response as HTML, and returns links in the HTML document
func Extract(url string) ([]string, error) {
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
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore invalid URL
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if cancelled() {
		return
	}

	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
