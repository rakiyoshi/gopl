package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	url := os.Args[1]
	id := os.Args[2]
	element, err := ex08(url, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	fmt.Println(element.Data, element.Attr)
}

func ex08(url, id string) (*html.Node, error) {
	doc, err := getURL(url)
	if err != nil {
		return nil, err
	}
	return ElementById(doc, id), nil
}

func ElementById(doc *html.Node, id string) *html.Node {
	return forEachNode(
		doc,
		func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				for _, a := range n.Attr {
					if a.Key == "id" && a.Val == id {
						return false
					}
				}
			}
			return true
		},
		nil,
	)
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

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && !pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cn := forEachNode(c, pre, post); cn != nil {
			return cn
		}
	}

	if post != nil && !post(n) {
		return n
	}
	return nil
}
