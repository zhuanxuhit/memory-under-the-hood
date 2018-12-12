package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	fmt.Printf("pid: %d\n", os.Getpid())
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			log.Println(float64(m.Sys) / 1024 / 1024)
			log.Println(float64(m.HeapAlloc) / 1024 / 1024)
			time.Sleep(10 * time.Second)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
