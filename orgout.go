package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chaseadamsio/goorgeous"
)

func ServeTheOrgs(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_fname := strings.Split(r.URL.Path, "/")[2:][0]
	fmt.Fprintf(w, "<html><body>")
	fname := fmt.Sprintf("%s/%s", cfg.OrgDir, _fname)
	log.Println("Rendering:", fname)
	if _, err := os.Stat(fname); err == nil {
		data, der := ioutil.ReadFile(fname)
		if der != nil {
			log.Println(der)
		}
		output := goorgeous.OrgCommon([]byte(data))
		fmt.Fprintf(w, string(output))
	}
	fmt.Fprintf(w, "</body></html>")
}
