package ex01

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	tt := struct {
		input string
		want  int
	}{
		"hello world",
		2,
	}

	var c WordCounter
	_, err := c.Write([]byte(tt.input))
	if err != nil {
		t.Fatalf("%v", err)
	}
	if int(c) != tt.want {
		t.Fatalf("got=%d, want=%d", c, tt.want)
	}
}

func TestFprintf(t *testing.T) {
	tt := struct {
		input string
		want  int
	}{
		"Dolly",
		3,
	}
	var c WordCounter
	fmt.Fprintf(&c, "hello, %s bye", tt.input)
	if int(c) != tt.want {
		t.Fatalf("got=%d, want=%d", c, tt.want)
	}
}
