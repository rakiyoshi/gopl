package ex10

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input sort.IntSlice
		want  bool
	}{
		{
			input: sort.IntSlice{1, 2, 3, 2, 1},
			want:  true,
		},
		{
			input: sort.IntSlice{1, 2, 2, 1},
			want:  true,
		},
		{
			input: sort.IntSlice{1, 1},
			want:  true,
		},
		{
			input: sort.IntSlice{1},
			want:  true,
		},
		{
			input: sort.IntSlice{1, 2, 3, 4},
			want:  false,
		},
		{
			input: sort.IntSlice{1, 2, 3},
			want:  false,
		},
	}

	for _, tt := range tests {
		if got := IsPalindrome(tt.input); got != tt.want {
			t.Fatalf("IsPalindrome(%v)=%v, want=%v", tt.input, got, tt.want)
		}
	}
}
