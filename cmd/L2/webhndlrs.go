package main

import (
	"fmt"
    "net/http"
)

// extend this with the steps shown at 	https://golang.org/doc/articles/wiki/

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}