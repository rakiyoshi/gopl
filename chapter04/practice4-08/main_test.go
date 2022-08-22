package main

import (
	"testing"
)

func TestCharcountTestCharcount(t *testing.T) {
	tests := []struct {
		input string
		want  Counts
	}{
		{
			input: "fmt.Println(\"Hello world\")",
			want: Counts{
				control:     0,
				letter:      20,
				mark:        0,
				number:      0,
				punctuation: 5,
				space:       1,
				symbol:      0,
			},
		},
		{
			input: `Donald J. Trump チャン、オッハー❗
😚今日のお弁当が美味しくて、一緒に〇〇チャンのことも、食べちゃいたいナ〜😍💕（笑）✋
ナンチャッテ😃💗`,
			want: Counts{
				control:     2,
				letter:      56,
				mark:        0,
				number:      2,
				punctuation: 7,
				space:       5,
				symbol:      7,
			},
		},
	}

	for _, tt := range tests {
		got := charcount(tt.input)
		if got != tt.want {
			t.Errorf("got = %q, want = %q", got, tt.want)
		}
	}
}
