package ex01

import (
	"fmt"
	"testing"
)

func TestLineCounterWrite(t *testing.T) {
	tt := struct {
		input string
		want  int
	}{
		"hello\nworld\n",
		2,
	}

	var c LineCounter
	_, err := c.Write([]byte(tt.input))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if int(c) != tt.want {
		t.Fatalf("got=%d, want=%d", c, tt.want)
	}
}

func TestLineCounterFprintf(t *testing.T) {
	tt := struct {
		input string
		want  int
	}{
		"Dolly",
		3,
	}
	var c LineCounter
	fmt.Fprintf(&c, "hello\n%s\nbye\n", tt.input)
	if int(c) != tt.want {
		t.Fatalf("got=%d, want=%d", c, tt.want)
	}
}
