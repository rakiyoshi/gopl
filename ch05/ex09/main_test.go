package main

import "testing"

type Arg struct {
	s string
	f func(string) string
}

func TestExpandTestExpand(t *testing.T) {
	tests := []struct {
		input Arg
		want  string
	}{
		{
			input: Arg{
				"This is a $pen.",
				func(s string) string {
					return s
				},
			},
			want: "This is a pen.",
		},
		{
			input: Arg{
				"This is a $pen.",
				func(s string) string {
					return "`" + s + "`"
				},
			},
			want: "This is a `pen.`",
		},
	}

	for _, tt := range tests {
		if got := expand(tt.input.s, tt.input.f); got != tt.want {
			t.Fatalf("got=%s, want=%s\n", got, tt.want)
		}
	}
}
