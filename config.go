package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	OrgDir       string `yaml:"org_dir"`
	StaticDir    string `yaml:"static_dir"`
	TemplateDir  string `yaml:"template_dir"`
	TemplateName string `yaml:"template"`
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
