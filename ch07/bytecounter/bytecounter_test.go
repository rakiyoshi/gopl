package bytecounter

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	tt := struct {
		input string
		want  int
	}{
		"hello",
		5,
	}

	var c ByteCounter
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
		12,
	}
	var c ByteCounter
	fmt.Fprintf(&c, "hello, %s", tt.input)
	if int(c) != tt.want {
		t.Fatalf("got=%d, want=%d", c, tt.want)
	}
}
