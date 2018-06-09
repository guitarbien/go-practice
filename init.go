package main

import (
  "fmt"
  _ "github.com/guitarbien/go-practice/foo"
  _ "github.com/guitarbien/go-practice/bar"
)

func init() {
  fmt.Println("init 1")
}

func init() {
  fmt.Println("init 2")
  global = 0
}

var global = convert()

func convert() int {
  return 100
}

func main() {
  fmt.Println("run main")
  fmt.Println("global is", global)
}
