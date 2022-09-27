package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
)

var tracks = []*Track{
	{"Go", "Moby", "Moby", 1992, parseTimeLength("3m37s")},
	{"Go", "Delilah", "From the Roots Up", 2012, parseTimeLength("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, parseTimeLength("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, parseTimeLength("4m24s")},
}

var htmlBody = template.Must(template.New("tracks").Parse(`
<head><title>Tracks</title></head>
<h1>Tracks</h1>
<table>
	<tr style ='text-align: left'>
		<th><a href="/?sort={{range .Sort}}{{.}},{{end}}byTitle">Title</a></th>
		<th><a href="/?sort={{range .Sort}}{{.}},{{end}}byArtist">Artist</a></th>
		<th><a href="/?sort={{range .Sort}}{{.}},{{end}}byAlbum">Album</a></th>
		<th><a href="/?sort={{range .Sort}}{{.}},{{end}}byYear">Year</a></th>
		<th><a href="/?sort={{range .Sort}}{{.}},{{end}}byLength">Length</a></th>
	</tr>
	{{range .Tracks}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Artist}}</td>
		<td>{{.Album}}</td>
		<td>{{.Year}}</td>
		<td>{{.Length}}</td>
	</tr>
	{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()["sort"]
	var sorts []string
	if len(qs) > 0 {
		sorts = strings.Split(qs[0], ",")
	}
	s := NewMultiSort(tracks)
	for _, sort := range sorts {
		switch sort {
		case "byTitle":
			s = s.ByTitle()
		case "byArtist":
			s = s.ByArtist()
		case "byAlbum":
			s = s.ByAlbum()
		case "byYear":
			s = s.ByYear()
		case "byLength":
			s = s.ByLength()
		}
	}
	sort.Sort(s)
	if err := htmlBody.Execute(w, struct {
		Tracks []*Track
		Sort   []string
	}{tracks, sorts}); err != nil {
		log.Fatal(err)
	}
}
