package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
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

	if r.URL.Path == "/favicon.ico" {
		http.ServeFile(w, r, fmt.Sprintf("%s%s", cfg.Content.OrgDir, r.URL.Path))
		return
	}

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
	_tmpl, ter := ioutil.ReadFile(fmt.Sprintf(
		"%s/%s/%s.html",
		cfg.Content.Base,
		cfg.Content.Template.Dir,
		cfg.Content.Template.Name,
	))
	if !check(ter, false) {
		Index(w, r)
	}

	// Setup the template
	tmpl, err := template.New("GoOrgEous George").Parse(string(_tmpl))
	_ = check(err, true)

	_path := fmt.Sprintf("%s/%s%s", cfg.Content.Base, cfg.Content.OrgDir, r.URL.Path)
	fn, er := ioutil.ReadDir(_path)
	_ = check(er, true)

	output := ""
	for _, f := range fn {
		if f.Name() == "index.org" {
			http.Redirect(w, r, fmt.Sprintf("/org%s/index.org", r.URL.Path), http.StatusFound)
			return
		}

		_path = strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
		if f.Name()[len(f.Name())-4:] == ".org" {
			output = fmt.Sprintf("%s<li><a href=\"/org/%s/%s\">%s</a></li>", output, r.URL.Path, f.Name(), f.Name())
		} else if f.Name() != "favicon.ico" {
			output = fmt.Sprintf("%s<li><a href=\"%s\">%s</a></li>", output, f.Name(), f.Name())
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
