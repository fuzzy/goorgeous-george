package main

import (
	"log"
	"fmt"
	"text/template"
	"net/http"
	"io/ioutil"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	fn, er := ioutil.ReadDir(cfg.Content.OrgDir)
	if er != nil {
		log.Println(er)
	}
	fmt.Fprintf(w, "<html><body><h1>Index</h1><hr /><ul>")
	for _, f := range fn {
		if f.Name()[len(f.Name())-4:] == ".org" {
			fmt.Fprintf(w, "<li><a href=\"/org/%s\">%s</a></li>", f.Name(), f.Name())
		}
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func ServeTheIndex(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()

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
		Index(w, r)
	}

	// Setup the template
	tmpl, err := template.New("GoOrgEous George").Parse(string(_tmpl))
	_ = check(err, true)

	fn, er := ioutil.ReadDir(cfg.Content.OrgDir)
	_ = check(er, true)
	
	output := ""
	for _, f := range fn {
		if f.Name()[len(f.Name())-4:] == ".org" {
			output = fmt.Sprintf("%s<li><a href=\"/org/%s\">%s</a></li>", output, f.Name(), f.Name())
		}
	}
	
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
		Index(w, r)
	}
}
