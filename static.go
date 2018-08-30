package main

import (
	"fmt"
	"log"
	"net/http"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	log.Printf("1: %s\n", r.URL.Path)
	fmt.Fprintf(w, "data")
	// if _, e := os.Stat()
}
