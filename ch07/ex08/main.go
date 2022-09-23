package ex08

import (
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type MultiSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func NewMultiSort(t []*Track) MultiSort {
	return MultiSort{t, func(_, _ *Track) bool { return false }}
}

func (m MultiSort) Len() int           { return len(m.t) }
func (m MultiSort) Less(i, j int) bool { return m.less(m.t[i], m.t[j]) }
func (m MultiSort) Swap(i, j int)      { m.t[i], m.t[j] = m.t[j], m.t[i] }

func (m MultiSort) ByTitle() MultiSort {
	return MultiSort{
		m.t,
		func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			return m.less(x, y)
		},
	}
}

func (m MultiSort) ByArtist() MultiSort {
	return MultiSort{
		m.t,
		func(x, y *Track) bool {
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
			return m.less(x, y)
		},
	}
}

func (m MultiSort) ByAlbum() MultiSort {
	return MultiSort{
		m.t,
		func(x, y *Track) bool {
			if x.Album != y.Album {
				return x.Album < y.Album
			}
			return m.less(x, y)
		},
	}
}

func (m MultiSort) ByYear() MultiSort {
	return MultiSort{
		m.t,
		func(x, y *Track) bool {
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			return m.less(x, y)
		},
	}
}

func (m MultiSort) ByLength() MultiSort {
	return MultiSort{
		m.t,
		func(x, y *Track) bool {
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return m.less(x, y)
		},
	}
}

func parseTimeLength(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
