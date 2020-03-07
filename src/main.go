package main

import (
  "fmt"
  "os"
  "os/user"
  "repl"
)

func main() {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("hello %s. enter commands pls\n", user.Username)
  repl.Start(os.Stdin, os.Stdout)
}
