package main

import (
	"bufio"
	"fmt"
	"gopl/ch05/links"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			item = strings.TrimSuffix(item, "/")
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(urlStr string) []string {
	fmt.Println(urlStr)
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Print(err)
	}
	if u.Hostname() == "go.dev" {
		archivePage(u.String())
	}

	list, err := links.Extract(urlStr)
	if err != nil {
		log.Print(err)
	}
	return list
}

func archivePage(urlStr string) {
	// TODO: determine path
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Print(err)
	}

	var filename, dirname string
	if u.Path == "" {
		filename = "index.html"
		dirname = "archive"
	} else {
		paths := strings.Split(u.Path, "/")
		filename = paths[len(paths)-1] + ".html"
		dirname = "archive/" + filepath.Join(paths[:len(paths)-1]...)
	}

	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(dirname + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("getting %s: %s", urlStr, resp.Status)
	}

	w := bufio.NewWriter(f)
	r := bufio.NewReader(resp.Body)
	_, err = io.Copy(w, r)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
