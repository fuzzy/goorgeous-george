package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
)

func ServeTheStatic(w http.ResponseWriter, r *http.Request) {
	cfg := ReadConfig()
	_path := strings.Join(strings.Split(r.URL.Path, "/")[2:], "/")
	_fname := fmt.Sprintf("%s/%s/%s", cfg.Content.Base, cfg.Content.StaticDir, _path)
	if _, e := os.Stat(_fname); e == nil {
		http.ServeFile(w, r, _fname)
	} else {
		log.Printf("File not found: %s", _fname)
	}
}

/*
 * Please understand, I am deeply ashamed of this. But don't worry, it will be going
 * away. I'm working on webhook triggers to accept notifications of updates from github,
 * gogs/gitea, gitlab etc. Please stay tuned.
 */
func UpdateTheRepo() {
	cfg := ReadConfig()
	repo, err := git.PlainOpen(cfg.Content.Base)
	if err != nil {
		log.Fatal(err)
	}
	wdir, err := repo.Worktree()
	err = wdir.Pull(&git.PullOptions{})
	if err != nil {
		log.Printf("%s: %s", cfg.Content.Base, err)
	} else {
		log.Printf("%s: updated successfully", cfg.Content.Base)
	}
}

func UpdateTheRepoThread() {
	for {
		time.Sleep(time.Minute * 1)
		UpdateTheRepo()
	}
}

// </this_particular_bit_of_shame>
