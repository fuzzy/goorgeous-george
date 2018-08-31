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
	fn, er := ioutil.ReadDir(cfg.Content.OrgDir)
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
	// Read in the config
	cfg := ReadConfig()

	// Register the route functions
	http.HandleFunc("/", Router)

	// Output and start serving
	log.Println(fmt.Sprintf("Serving on %s:%s", cfg.Network.Interface, cfg.Network.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Network.Interface, cfg.Network.Port), nil))
}
