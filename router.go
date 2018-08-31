package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Main view\n")
	_path := strings.Split(r.URL.Path, "/")[1:]
	log.Printf("Request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
	switch _path[0] {
	case "static":
		fmt.Fprintf(w, "Static file: %V\n", _path[1:])
	case "org":
		if len(_path) >= 2 {
			ServeTheOrgs(w, r)
		} else {
			Index(w, r)
		}
	default:
		Index(w, r)
	}
	// path := strings.Join(_path[:len(_path)-1], "/")
	// fmt.Fprintf(w, "%+V", path)
}
