package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/chaseadamsio/goorgeous"
)

func ServeTheOrgsRaw(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_fname := strings.Split(r.URL.Path, "/")[2:][0]
	fmt.Fprintf(w, "<html><body>")
	fname := fmt.Sprintf("%s/%s", cfg.Content.OrgDir, _fname)
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

func ServeTheOrgs(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_fname := strings.Split(r.URL.Path, "/")[2:][0]
	fname := fmt.Sprintf("%s/%s", cfg.Content.OrgDir, _fname)
	if _, err := os.Stat(fname); err == nil {

		// setup check()
		check := func(err error, fatal bool) bool {
			if err != nil {
				if fatal {
					log.Fatal(err)
				} else {
					log.Println(err)
				}
				return false
			}
			return true
		}

		// read in template
		_tmpl, ter := ioutil.ReadFile(fmt.Sprintf("%s/%s.html", cfg.Template.Dir, cfg.Template.Name))
		if !check(ter, false) {
			ServeTheOrgsRaw(w, r)
		}

		// read in org doc
		data, der := ioutil.ReadFile(fname)
		_ = check(der, false)

		// parse org document
		output := goorgeous.OrgCommon([]byte(data))

		// Setup the template
		tmpl, err := template.New("GoOrgEous George").Parse(string(_tmpl))
		_ = check(err, true)

		// Create the payload
		_data := struct {
			Title   string
			Payload string
		}{
			Title:   "GoOrgEous George",
			Payload: string(output),
		}

		// and fire this shit off to the browser
		if !check(tmpl.Execute(w, _data), false) {
			ServeTheOrgsRaw(w, r)
		}
	}
}
