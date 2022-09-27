package ex08

import (
	"fmt"
	"sort"
	"testing"
)

func TestEachMultiSort(t *testing.T) {
	input := []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}

	sort.Sort(NewMultiSort(input).ByTitle())
	want := []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}

	sort.Sort(NewMultiSort(input).ByArtist())
	want = []*Track{
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
	}
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}

	sort.Sort(NewMultiSort(input).ByAlbum())
	want = []*Track{
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}

	sort.Sort(NewMultiSort(input).ByYear())
	want = []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
	}
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}

	sort.Sort(NewMultiSort(input).ByLength())
	want = []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
	}
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestCombinedMultiSort(t *testing.T) {
	input := []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}
	want := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}

	sort.Sort(NewMultiSort(input).ByArtist().ByTitle())
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}

	sort.Sort(NewMultiSort(input).ByTitle().ByArtist().ByAlbum().ByYear().ByLength().ByArtist().ByTitle())
	if err := checkTrack(input, want); err != nil {
		t.Fatalf("%v", err)
	}
}

func BenchmarkAllMultiSort(b *testing.B) {
	input := []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}
	for i := 0; i < b.N; i++ {
		sort.Sort(NewMultiSort(input).ByTitle().ByArtist().ByAlbum().ByYear().ByLength())
	}
}

func BenchmarkAllStableSort(b *testing.B) {
	input := []*Track{
		{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
	}
	for i := 0; i < b.N; i++ {
		sort.Stable(customSort{input, func(x, y *Track) bool { return x.Length < y.Length }})
		sort.Stable(customSort{input, func(x, y *Track) bool { return x.Year < y.Year }})
		sort.Stable(customSort{input, func(x, y *Track) bool { return x.Album < y.Album }})
		sort.Stable(customSort{input, func(x, y *Track) bool { return x.Artist < y.Artist }})
		sort.Stable(customSort{input, func(x, y *Track) bool { return x.Title < y.Title }})
	}
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func checkTrack(input, want []*Track) error {
	if len(input) != len(want) {
		return fmt.Errorf("invalid length: len(got)=%d, len(want)=%d", len(input), len(want))
	}
	for i := range input {
		if *input[i] != *want[i] {
			return fmt.Errorf("*input[%d]=%v, *want[%d]=%v", i, *input[i], i, *want[i])
		}
	}
	return nil
}
