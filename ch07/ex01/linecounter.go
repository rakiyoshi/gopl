package ex01

import (
	"bufio"
	"bytes"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	sc := bufio.NewScanner(buf)

	var counter int
	for sc.Scan() {
		counter++
	}
	*c += LineCounter(counter)
	return counter, nil
}
