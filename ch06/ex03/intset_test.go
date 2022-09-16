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

func TestIntersectWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(10)

	y.Add(9)
	y.Add(10)
	y.Add(42)

	x.IntersectWith(&y)
	want := "{9 10}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}

func TestDifferenceWith(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(10)

	y.Add(9)
	y.Add(10)
	y.Add(42)

	x.DifferenceWith(&y)
	want := "{1 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(10)

	y.Add(9)
	y.Add(10)
	y.Add(42)

	x.SymmetricDifference(&y)
	want := "{1 42 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
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
