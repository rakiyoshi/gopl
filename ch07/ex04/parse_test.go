package ex04

import (
	"testing"

	"golang.org/x/net/html"
)

func TestParse(t *testing.T) {
	input := `<html>
	</html>
	`

	doc, err := parse(input)
	if err != nil {
		t.Errorf("%v", err)
	}
	if doc.Type != html.DocumentNode {
		t.Errorf("doc.Type is invalid. got=%d, want=%d", doc.Type, html.DocumentNode)
	}
}
