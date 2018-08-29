package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/chaseadamsio/goorgeous"
	"gopkg.in/yaml.v2"
)

type Config struct {
	OrgDir    string `yaml:"org_dir"`
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
	log.Printf("1: %s\n", r.URL.Path)
	fmt.Fprintf(w, "data")
	// if _, e := os.Stat()
}

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

func main() {
	// Register the route functions
	http.HandleFunc("/", Router)

	// Output and start serving
	log.Println("Serving on 0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
