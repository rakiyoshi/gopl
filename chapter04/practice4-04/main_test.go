package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		s    []int
		n    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{3, 4, 1, 2},
		},
	}

	for _, tt := range tests {
		rotate(&tt.s, tt.n)
		if len(tt.s) != len(tt.want) {
			t.Errorf("len(tt.s) != len(tt.want)")
		}
		for i := 0; i < len(tt.s); i++ {
			if tt.s[i] != tt.want[i] {
				t.Errorf("tt.s[%d] = %d, want=%d", i, tt.s[i], tt.want[i])
			}
		}
	}
}
