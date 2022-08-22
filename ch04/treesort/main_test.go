package main

import "testing"

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
