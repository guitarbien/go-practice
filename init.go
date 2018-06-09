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
}

func main() {
  fmt.Println("run man")
}
