package main

import (
	"os"
	"fmt"
	"log"
	"html"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
	"github.com/chaseadamsio/goorgeous"
)

type Config struct {
	OrgDir string `yaml:"org_dir"`
	StaticDir string `yaml:"static_dir"`
}

func ReadConfig() *Config {
	retv := &Config{}
	fn := fmt.Sprintf("%s/george.yml", os.Getenv("PWD"))
	if _, err := os.Stat(fn); err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(data), retv)
	if err != nil {
		log.Fatal(err)
	}
	return retv
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "data")
}

func ServeTheOrgs(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	if len(html.EscapeString(r.URL.Path)) == 1 {
		fmt.Fprintf(w, "<html><body><h1>Index</h1><hr /><ul>")
		fn, er := ioutil.ReadDir(cfg.OrgDir)
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
		fname := fmt.Sprintf("%s%s", cfg.OrgDir, r.URL.Path)
		if _, err := os.Stat(fname); err == nil {
			data, der := ioutil.ReadFile(fname)
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
	// Register the route functions
	http.HandleFunc("/", ServeTheOrgs)
	http.HandleFunc("/static", ServeStatic)

	// Output and start serving
	log.Println("Serving on 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
