package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	return forEachNode(
		doc,
		func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				if contains(name, n.Data) {
					return true
				}
			}
			return false
		},
		nil,
	)
}

func contains(elems []string, str string) bool {
	for _, elem := range elems {
		if elem == str {
			return true
		}
	}
	return false
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

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) []*html.Node {
	var nodes []*html.Node
	if pre != nil && pre(n) {
		nodes = append(nodes, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cn := forEachNode(c, pre, post); len(cn) > 0 {
			nodes = append(nodes, cn...)
		}
	}

	if post != nil && post(n) {
		nodes = append(nodes, n)
	}
	return nodes
}

func main() {
	url := os.Args[1]
	doc, err := getURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Println("### img")
	images := ElementsByTagName(doc, "img")
	for _, img := range images {
		fmt.Println(img.Attr)
	}
	fmt.Println("### headings")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, hdg := range headings {
		fmt.Println(hdg.Attr)
	}
}
