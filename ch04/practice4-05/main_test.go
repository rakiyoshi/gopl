package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		s    []string
		want []string
	}{
		{
			[]string{"apple", "apple", "orange"},
			[]string{"apple", "orange"},
		},
	}

	for _, tt := range tests {
		got := uniq(tt.s)
		fmt.Printf("%q", got)
		if len(got) != len(tt.want) {
			t.Errorf("len(got) = %d, want %d", len(got), len(tt.want))
		}
		for i := 0; i < len(tt.want); i++ {
			if tt.s[i] != tt.want[i] {
				t.Errorf("tt.s[%d] = %s, want=%s", i, tt.s[i], tt.want[i])
			}
		}
	}
}
