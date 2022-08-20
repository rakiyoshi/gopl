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
		if got.control != tt.want.control {
			t.Errorf("got.control = %d, want = %d", got.control, tt.want.control)
		}
		if got.letter != tt.want.letter {
			t.Errorf("got.letter = %d, want = %d", got.letter, tt.want.letter)
		}
		if got.mark != tt.want.mark {
			t.Errorf("got.mark = %d, want = %d", got.mark, tt.want.mark)
		}
		if got.number != tt.want.number {
			t.Errorf("got.number = %d, want = %d", got.number, tt.want.number)
		}
		if got.punctuation != tt.want.punctuation {
			t.Errorf("got.punctuation = %d, want = %d", got.punctuation, tt.want.punctuation)
		}
		if got.space != tt.want.space {
			t.Errorf("got.space = %d, want = %d", got.space, tt.want.space)
		}
		if got.symbol != tt.want.symbol {
			t.Errorf("got.symbol = %d, want = %d", got.symbol, tt.want.symbol)
		}
	}
}
