package main

import "fmt"

func main() {
  foo := func() string {
    return "Hello World"
  }

  fmt.Println(foo())

  bar := func() {
    fmt.Println("Hello World 2")
  }

  bar()

  func() {
    fmt.Println("Hello World 3")
  }()
}
