package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Network struct {
		Interface string `yaml:"interface"`
		Port      string `yaml:"port"`
	} `yaml:"network"`
	Content struct {
		OrgDir    string `yaml:"org_dir"`
		StaticDir string `yaml:"static_dir"`
	} `yaml:"content"`
	Template struct {
		Dir  string `yaml:"dir"`
		Name string `yaml:"name"`
	} `yaml:"template"`
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
