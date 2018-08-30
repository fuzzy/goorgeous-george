package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	fmt.Fprintf(w, "<html><body><h1>Index</h1><hr /><ul>")
	fn, er := ioutil.ReadDir(cfg.OrgDir)
	if er != nil {
		log.Println(er)
	}
	for _, f := range fn {
		if f.Name()[len(f.Name())-4:] == ".org" {
			fmt.Fprintf(w, "<li><a href=\"/org/%s\">%s</a></li>", f.Name(), f.Name())
		}
	}
	fmt.Fprintf(w, "</ul></body></html>")
}

func main() {
	// Register the route functions
	http.HandleFunc("/", Router)

	// Output and start serving
	log.Println("Serving on 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
