package main

import (
	"fmt"
	"gopl/ch07/eval"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Handler(w http.ResponseWriter, req *http.Request) {
	// TODO: URLEncode & URLDecode
	formula := req.URL.Query().Get("f")
	if len(formula) == 0 {
		http.Error(w, "f is required: ", http.StatusBadRequest)
		return
	}
	expr, err := eval.Parse(formula)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
	}
	env := make(eval.Env)
	fmt.Fprintf(w, "%f", expr.Eval(env))
}
