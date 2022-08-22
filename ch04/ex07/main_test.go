package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{
			[]byte("こんにちは"),
			[]byte("はちにんこ"),
		},
		{
			[]byte("John ペトルッチ"),
			[]byte("チッルトペ nhoJ"),
		},
		{
			[]byte("ＭＥＴＡＬ🤘"),
			[]byte("🤘ＬＡＴＥＭ"),
		},
	}

	for _, tt := range tests {
		reverse(tt.input)
		if len(tt.input) != len(tt.want) {
			t.Errorf("len(tt.input)=%d, want=%d", len(tt.input), len(tt.want))
		}
		for i := range tt.input {
			if tt.input[i] != tt.want[i] {
				t.Errorf("tt.input[%d]=0x%02x, want=0x%02x", i, tt.input[i], tt.want[i])
			}
		}
	}
}
