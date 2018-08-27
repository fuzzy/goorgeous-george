package main

import (
	"os"
	"fmt"
	"log"
	"html"
	"io/ioutil"
	"net/http"
	
	"github.com/chaseadamsio/goorgeous"
)

func ServeTheOrgs(w http.ResponseWriter, r *http.Request) {
	if len(html.EscapeString(r.URL.Path)) == 1 {
		fmt.Fprintf(w, "<html><body><h1>Index</h1><hr /><ul>")
		fn, er := ioutil.ReadDir("./")
		if er != nil {
			log.Fatal(er)
		}
		for _, f := range fn {
			if f.Name()[len(f.Name())-4:] == ".org" {
				fmt.Fprintf(w, "<li><a href=\"/%s\">%s</a></li>", f.Name(), f.Name())
			}
		}
		fmt.Fprintf(w, "</ul></body></html>")
	} else {
		fmt.Fprintf(w, "<html><body>")
		if _, err := os.Stat(fmt.Sprintf(".%s", r.URL.Path)); err == nil {
			data, der := ioutil.ReadFile(fmt.Sprintf(".%s", r.URL.Path))
			if der != nil {
				log.Fatal(der)
			}
			output := goorgeous.OrgCommon([]byte(data))
			fmt.Fprintf(w, string(output))
		}
		fmt.Fprintf(w, "</body></html>")
	}
}

func main() {

	log.Println("Registering path handlers.")
	http.HandleFunc("/", ServeTheOrgs)
	log.Println("Serving on 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
