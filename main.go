package main

import (
  "fmt"
)

const (
  Monday    = iota + 1
  Tuesday
  Wednesday
  Thursday
  Friday
  Saturday
  Sunday
)

func add(i, j int) int {
  return i + j
}

func swap(i, j int) (int, int) {
  return j, i
}

func foo() func() int {
  return func() int {
    return 100
  }
}

func main() {
  foobar := foo()
  fmt.Printf("%T\n", foobar)
  fmt.Println(foobar())

  foo := "Hello"
  bar := 100

  fmt.Println(add(2, 3))

  a := 1
  b := 2
  a, b = swap(a, b)
  fmt.Println("a:", a)
  fmt.Println("b:", b)

  a, b = b, a
  fmt.Println("a:", a)
  fmt.Println("b:", b)

  fmt.Println(foo)
  fmt.Println(bar)
  fmt.Println("done !!")

  fmt.Println(Monday)
  fmt.Println(Tuesday)
  fmt.Println(Wednesday)
  fmt.Println(Thursday)
  fmt.Println(Friday)
  fmt.Println(Saturday)
  fmt.Println(Sunday)
}
