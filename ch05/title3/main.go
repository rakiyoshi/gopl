package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			// EXPECTED panic
			err = fmt.Errorf("multiple title elements")
		default:
			// unexpected (normal) panic. keep panic
			panic(p)
		}
	}()

	forEachNode(
		doc,
		func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
				if title != "" {
					panic(bailout{}) // doc has multiple title
				}
				title = n.FirstChild.Data
			}
		},
		nil,
	)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
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

func getURL(url string) (*html.Node, error) {
	resp, err := http.Get(url)
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
	return doc, nil
}

func main() {
	doc, err := getURL(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	title, err := soleTitle(doc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Println(title)
}
