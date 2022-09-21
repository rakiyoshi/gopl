package ex02

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var buf bytes.Buffer
	w, c := CountingWriter(&buf)
	tt := struct {
		inputs []string
		want   int64
	}{
		inputs: []string{
			"hello",
			"world",
		},
		want: 10,
	}

	_, err := w.Write([]byte(tt.inputs[0]))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if buf.String() != tt.inputs[0] {
		t.Fatalf("failed to write to buf: got=%s, want=%s", buf.String(), tt.inputs[0])
	}

	_, err = w.Write([]byte(tt.inputs[1]))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if buf.String() != tt.inputs[0]+tt.inputs[1] {
		t.Fatalf("failed to write to buf: got=%s, want=%s", buf.String(), tt.inputs[0]+tt.inputs[1])
	}

	if *c != tt.want {
		t.Fatalf("got=%d, want=%d", *c, tt.want)
	}
}
