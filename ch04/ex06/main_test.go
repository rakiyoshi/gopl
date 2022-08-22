package main

import (
	"testing"
	"unicode/utf8"
)

func TestMain(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{
			[]byte("こんにちは　世界"),
			[]byte("こんにちは 世界"),
		},
		{
			[]byte("こんにちは　　世界"),
			[]byte("こんにちは 世界"),
		},
		{
			[]byte("ＭＥＬＯＤＩＣ　ＤＥＡＴＨ 　　 ＭＥＴＡＬ"),
			[]byte("ＭＥＬＯＤＩＣ ＤＥＡＴＨ ＭＥＴＡＬ"),
		},
	}

	for _, tt := range tests {
		got := []rune(string(compressSpace(tt.input)))
		wantRune := []rune(string(tt.want))
		if utf8.RuneCountInString(string(got)) != utf8.RuneCountInString(string(tt.want)) {
			t.Errorf("got=%q, want=%q\n", got, tt.want)
		}
		for i := range got {
			if got[i] != wantRune[i] {
				t.Errorf("got[%d]=%q, want[i]=%q\n", i, got[i], wantRune[i])
			}
		}
	}
}
