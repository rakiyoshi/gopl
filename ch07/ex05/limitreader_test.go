package ex05

import (
	"strings"
	"testing"
)

func TestReadString(t *testing.T) {
	tests := []struct {
		input string
		n     int64
		want  string
	}{
		{
			input: "hello",
			n:     4,
			want:  "hell",
		},
		{
			input: "hello",
			n:     6,
			want:  "hello",
		},
	}

	for _, tt := range tests {
		r_ := strings.NewReader(tt.input)
		r := LimitReader(r_, tt.n)
		buf := make([]byte, tt.n)
		n, err := r.Read(buf)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if n != len(tt.want) {
			t.Fatalf("failed to read. read size = %d, want = %d", n, len(tt.want))
		}
		if string(buf[0:n]) != tt.want {
			t.Fatalf("failed to read. got=%s, want=%s", string(buf[0:n]), tt.want)
		}
	}
}
