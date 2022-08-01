package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Let's get to work, %q!", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
