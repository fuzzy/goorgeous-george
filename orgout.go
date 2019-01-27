package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/donaldh/goorgeous"
)

func ServeTheOrgsRaw(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_fname := strings.Split(r.URL.Path, "/")[2:][0]
	fname := fmt.Sprintf("%s/%s", cfg.Content.OrgDir, _fname)
	log.Println("Rendering:", fname)
	if _, err := os.Stat(fname); err == nil {
		data, der := ioutil.ReadFile(fname)
		if der != nil {
			log.Println(der)
		}
		fmt.Fprintf(w, string(data))
	}
}

func ServeTheOrgs(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_fname := strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
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
			Path    string
			Title   string
			Payload string
		}{
			Path:    fmt.Sprintf(_fname),
			Title:   "GoOrgEous George",
			Payload: string(output),
		}

		// and fire this shit off to the browser
		if !check(tmpl.Execute(w, _data), false) {
			ServeTheOrgsRaw(w, r)
		}
	}
}
