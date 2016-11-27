package main

import (
	"fmt"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[ResponseWriter]:{\n%+v\n}\n[Request]:{\n%+v\n}\n", w, r)
	fmt.Printf("[ResponseWriter]:{\n%+v\n}\n[Request]:{\n%+v\n}\n", w, r)
}

func main() {
	http.HandleFunc("/", handlerRoot)
	http.ListenAndServe(":8080", nil)
}
