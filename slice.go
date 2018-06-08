package main

import "fmt"

// pass slice as function argument
func addValue(foo []string) []string {
  foo = append(foo, "c")
  fmt.Println("addValue foo", foo)
  return foo
}

func main() {
  foo := []string{"a", "b"}
  fmt.Println("before foo:", foo)

  addValue(foo)
  fmt.Println("after foo:", foo)

  bar := foo[:1]
  fmt.Println("bar:", bar)

  s1 := append(bar, "c")
  fmt.Println("foo:", foo)
  fmt.Println("s1:", s1)

  s2 := append(bar, "d")
  fmt.Println("foo:", foo)
  fmt.Println("s2:", s2)
}
