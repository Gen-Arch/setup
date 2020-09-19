package main

import (
  "os"
  "path/filepath"
)

type Gits struct {
  url string
  dir string
}

func rcs() []string {
  rcfiles := []string {
    "vimrc",
    "zshrc",
    "bashrc",
    "aliasrc",
    "tmux.conf",
    "hyper.js",
  }

  return rcfiles
}

func gits() []Gits {
  gitfiles := []Gits {
    Gits{"https://github.com/anyenv/anyenv",      filepath.Join(os.Getenv("HOME"), ".anyenv") },
    Gits{"https://github.com/mollifier/anyframe", filepath.Join(os.Getenv("HOME"), ".temp/anyframe")},
  }

  return gitfiles

}
