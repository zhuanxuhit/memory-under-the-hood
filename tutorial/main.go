package main

import (
	"net/http"
	"fmt"
	"html"
	"os"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	fmt.Printf("pid: %d\n", os.Getpid())
	http.ListenAndServe(":8080", nil)
}
