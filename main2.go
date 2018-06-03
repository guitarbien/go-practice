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

  go func(i, j int) {
    fmt.Println(i + j)
  }(1, 2)
}
