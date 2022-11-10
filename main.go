package main

import (
	"fmt"
	"net/http"
)

func RegisterHttpHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you've requested path: %s\n", r.URL.Path)
	})
}

func main() {
	RegisterHttpHandlers()
	print("\nlistening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
