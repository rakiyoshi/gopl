package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	want := 3
	if x.Len() != want {
		t.Fatalf("got=%d, want=%d", x.Len(), want)
	}
}

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

func TestRemove(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Remove(1)
	want := "{9 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}

	x.Remove(1)
	want = "{9 144}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}

func TestClear(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Clear()
	want := "{}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}

	x.Add(1)
	want = "{1}"
	if x.String() != want {
		t.Fatalf("got=%s, want=%s", x.String(), want)
	}
}

func TestCopy(t *testing.T) {
	var x IntSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	y := x.Copy()
	want := "{1 9 144}"
	if y.String() != want {
		t.Fatalf("got=%s, want=%s", y.String(), want)
	}

	y.Add(25)
	want = "{1 9 25 144}"
	if y.String() != want {
		t.Fatalf("got=%s, want=%s", y.String(), want)
	}

	want = "{1 9 144}"
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
