package popcount

import "testing"

func TestPopCount(t *testing.T) {
	result := PopCount(14)
	if result != 3 {
		t.Fatalf("want=3, got=%v", result)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
