package ex05

import "io"

type limitReader struct {
	r io.Reader // underlying reader
	n int64     // max bytes remaining
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
