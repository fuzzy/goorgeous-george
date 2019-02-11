package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type GitConfig struct {
	Repo    string `yaml:"repo"`
	WebHook bool   `yaml:"webhook"`
}

type ContentConfig struct {
	Base      string         `yaml:"base_dir"`
	OrgDir    string         `yaml:"org_dir"`
	StaticDir string         `yaml:"static_dir"`
	Template  TemplateConfig `yaml:"template"`
}

type TemplateConfig struct {
	Dir  string `yaml:"dir"`
	Name string `yaml:"name"`
}

type Config struct {
	Interface string        `yaml:"interface"`
	Port      string        `yaml:"port"`
	Git       GitConfig     `yaml:git`
	Content   ContentConfig `yaml:"content"`
}

func ReadConfig() *Config {
	retv := &Config{}
	fn := fmt.Sprintf("/config/george.yml")
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

	if retv.Git.Repo != "" {
		data := strings.Split(retv.Git.Repo, "/")
		tdir := fmt.Sprintf("%s/%s", retv.Content.Base, data[len(data)-1])
		retv.Content.Base = tdir
	}

	return retv
}
