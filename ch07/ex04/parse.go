package ex04

import (
	"io"

	"golang.org/x/net/html"
)

type easyReader struct {
	s string
	i int64
}

func parse(htmlStr string) (doc *html.Node, err error) {
	htmlReader := &easyReader{htmlStr, 0}
	doc, err = html.Parse(htmlReader)
	return
}

func (r *easyReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	if n < len(b) {
		err = io.EOF
	}
	r.i += int64(n)
	return
}
