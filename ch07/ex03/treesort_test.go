package ex03

import (
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			input: []int{4, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 4},
		},
	}

	for _, tt := range tests {
		Sort(tt.input)
		if len(tt.input) != len(tt.want) {
			t.Errorf("len(tt.input) = %d, want = %d", len(tt.input), len(tt.want))
		}
		for i := range tt.input {
			if tt.input[i] != tt.want[i] {
				t.Errorf("tt.input[%d] = %d, want = %d", i, tt.input[i], tt.want[i])
			}
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		input []int
		want  string
	}{
		{
			input: []int{5, 4, 3, 2, 1},
			want:  "((((( 1 ) 2 ) 3 ) 4 ) 5 )",
		},
		{
			input: []int{4, 4, 3, 2, 1},
			want:  "(((( 1 ) 2 ) 3 ) 4 ( 4 ))",
		},
	}

	for _, tt := range tests {
		var root *tree
		for _, v := range tt.input {
			root = add(root, v)
		}
		if root.String() != tt.want {
			t.Fatalf("got=%s, want=%s", root.String(), tt.want)
		}
	}
}
