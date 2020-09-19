package main

import (
  "os"
  "os/exec"
  "fmt"
  "strings"
  "path/filepath"
  "gopkg.in/src-d/go-git.v4"
)


func  git_clone(repo Gits) string{
  var err error
  fmt.Printf(repo.dir + "\n")

  if _, err := os.Stat(repo.dir); err == nil {
    return "Directory already exists\n"
  }

  _, err = git.PlainClone(repo.dir, false, &git.CloneOptions{
    URL: repo.url,
  })

  if err != nil {
    fmt.Printf("error -> " + repo.url + "\n")
    os.Exit(1)
  }
  return "clone repository\n"
}

func symlink(file string) {
  var symfile  string = filepath.Join(os.Getenv("HOME"), "." + file)
  var realfile string = filepath.Join(os.Getenv("HOME"), "dotfiles", "configs", file)

  if _, err := os.Stat(symfile); err == nil {
    os.Remove(symfile)
    fmt.Printf("remove symlink -> " + symfile + "\n")
  }

  os.Symlink(realfile, symfile)
  fmt.Printf("create symlink -> " + symfile + "\n")
}


func run(comand string) {
  coms := strings.Split(comand, " ")
  com  := coms[0]
  opt  := coms[1:]

  out, err := exec.Command(com, opt...).Output()

  if err != nil {
    fmt.Printf("command error -> "+ com + " " + strings.Join(opt, " ") + "\n")
    fmt.Printf("output: %s", out)
    os.Exit(1)
  }
  fmt.Printf("%s", out)
}

func main() {
  dotfile_dir := filepath.Join(os.Getenv("HOME"),"dotfiles", "configs")
  repos       := gits()
  rcfiles     := rcs(dotfile_dir)
  commands    := setup_commands()

  for _, repo := range repos {
    res := git_clone(repo)
    fmt.Printf(res)
  }

  for _, file := range rcfiles {
    symlink(file)
  }

  for _, command := range commands {
    run(command)
  }
}
