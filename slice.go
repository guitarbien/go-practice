package main

import "fmt"

// pass slice as function argument
func addValue(foo []string) []string {
  foo = append(foo, "c")
  fmt.Println("addValue foo", foo)
  return foo
}

func main()  {
  foo := []string{"a", "b"}
  fmt.Println("before foo:", foo)

  foo = addValue(foo)
  fmt.Println("after foo:", foo)
}
