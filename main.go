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

func add(i int, j int) int {
  return i + j
}

func main() {
  foo := "Hello"
  bar := 100

  fmt.Println(add(2, 3))

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
