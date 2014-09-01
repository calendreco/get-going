package main

import (
	"fmt"
	"net/http"
	"os"

	// "github.com/calendreco/get-going/lib/greetings"
	// "github.com/calendreco/get-going/lib/greetings"
	"github.com/calendreco/get-going/lib/greetings"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s to the home page!", greetings.Friendly)
	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"greeting": greetings.Friendly})
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	// Example middleware
	// n.Use(gzip.Gzip(gzip.DefaultCompression))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	host := os.Getenv("HOST")
	n.Run(host + ":" + port)
}
