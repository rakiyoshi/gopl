package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := htmlBody.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500: Internal Server Error")
		return
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

var htmlBody = template.Must(template.New("tracks").Parse(`
<head><title>Tracks</title></head>
<h1>Tracks</h1>
<table>
	<tr style ='text-align: left'>
		<th>Item</th>
		<th>Price</th>
	</tr>
	{{range $item, $price := .}}
	<tr>
		<td>{{ $item }}</td>
		<td>{{ $price }}</td>
	</tr>
	{{end}}
</table>
`))
