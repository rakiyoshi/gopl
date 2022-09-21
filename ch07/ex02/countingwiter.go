package ex02

import "io"

type countingWriter struct {
	writer  io.Writer
	counter int64
}

func (w *countingWriter) Write(p []byte) (int, error) {
	n, err := w.writer.Write(p)
	w.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c int64
	cw := &countingWriter{w, c}
	return cw, &cw.counter
}
