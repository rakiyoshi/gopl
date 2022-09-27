package ex01

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	sc := bufio.NewScanner(buf)
	sc.Split(bufio.ScanWords)

	var counter int
	for sc.Scan() {
		counter++
	}
	*c += WordCounter(counter)
	return counter, nil
}
