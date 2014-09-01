package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/phyber/negroni-gzip/gzip"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	// Example middleware
	n.Use(gzip.Gzip(gzip.DefaultCompression))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	host := os.Getenv("HOST")
	n.Run(host + ":" + port)
}
