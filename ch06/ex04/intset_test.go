package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	want := "{1 9 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}

func TestUnionWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	y.Add(9)
	y.Add(42)

	x.UnitonWith(&y)
	want := "{1 9 42 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}
func TestElems(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	got := x.Elems()
	want := []int{1, 9, 42, 144}
	if len(got) != len(want) {
		t.Fatalf("invalid length: len(got)=%d, len(want)=%d", len(got), len(want))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("invalid value: got[%d] = %d, want[%d] = %d", i, got[i], i, want[i])
		}
	}
}

func TestHas(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if !x.Has(9) {
		t.Fatalf("got=%v, want=%v", x.Has(9), true)
	}
	if x.Has(123) {
		t.Fatalf("got=%v, want=%v", x.Has(123), false)
	}
}
