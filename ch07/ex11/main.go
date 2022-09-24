package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

// warn: create handles all method
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if len(item) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "item is required\n")
		return
	}

	priceStr := req.URL.Query().Get("price")
	if len(priceStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "price is required\n")
		return
	}
	p, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
	}
	price := dollars(p)

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	db[item] = dollars(price)
	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

// warn: update handles all method
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	priceStr := req.URL.Query().Get("price")
	if len(priceStr) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "price is required\n")
		return
	}
	p, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", priceStr)
	}
	price := dollars(p)

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	db[item] = price
	fmt.Fprintf(w, "%s: %s\n", item, price)
}

// warn: delete handles all method
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "%s has been deleted\n", item)
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
