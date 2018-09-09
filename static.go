package main

import (
	"os"
	"fmt"
	"log"
	"strings"
	"net/http"
)

func ServeTheStatic(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_path := strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
	_fname := fmt.Sprintf("%s/%s", cfg.Content.StaticDir, _path)
	if _, e := os.Stat(_fname); e == nil {
		http.ServeFile(w, r, _fname)
	} else {
		log.Printf("File not found: %s", _fname)
	}
}

