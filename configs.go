package main

import (
  "os"
  "path/filepath"
  "io/ioutil"
)

type Gits struct {
  url string
  dir string
}

func rcs(dir string) []string {
  var rcfiles []string
  var files   []os.FileInfo

  files, err := ioutil.ReadDir(dir)
  if err != nil {
    panic(err)
  }

  for _, file := range files {
    rcfiles = append(rcfiles, file.Name())
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

func setup_commands() []string {
  commands := []string {
  }

  return commands
}
