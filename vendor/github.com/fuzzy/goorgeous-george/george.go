package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/src-d/go-git.v4"
)

func AppStart(c *cli.Context) error {
	cfg := Config{
		Interface: c.String("host"),
		Port:      c.String("port"),
		Git: GitConfig{
			Repo:    c.String("repo"),
			WebHook: c.Bool("webhook"),
		},
		Content: ContentConfig{
			Base:      "/data",
			OrgDir:    "org",
			StaticDir: "static",
			Template: TemplateConfig{
				Dir:  "template",
				Name: "george.html",
			},
		},
	}

	// Register the route functions
	http.HandleFunc("/", Router)

	// Clone the repo
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/src-d/go-git",
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output and start serving
	log.Println(fmt.Sprintf("Serving on %s:%s", cfg.Interface, cfg.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Interface, cfg.Port), nil))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "GoOrgEous George"
	app.Usage = "Render Org files as HTML, live."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "host", Value: "0.0.0.0", Usage: "Set the interface to listen on."},
		cli.StringFlag{Name: "port", Value: "8080", Usage: "Port to listen on."},
		cli.StringFlag{Name: "repo", Value: "", Usage: "Specify a git repo to clone for content."},
		cli.BoolFlag{Name: "webhook", Usage: "Enable the webhook for github style integration."},
	}
	app.Action = AppStart
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
