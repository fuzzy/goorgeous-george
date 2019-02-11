package main

import (
	"log"
	"net/http"
	"strings"
)

func Router(w http.ResponseWriter, r *http.Request) {
	_path := strings.Split(r.URL.Path, "/")[1:]

	log.Printf("Request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)

	switch _path[0] {
	case "static":
		ServeTheStatic(w, r)
	case "raw":
		ServeTheOrgsRaw(w, r)
	case "org":
		if len(_path) >= 2 {
			ServeTheOrgs(w, r)
		} else {
			ServeTheIndex(w, r)
		}
	default:
		ServeTheIndex(w, r)
	}
}
