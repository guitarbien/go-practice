package main

import "fmt"

func main()  {
  checkValue(0)
  checkValue(1)
}

func checkValue(i int) {
  switch i {
  case 0:
    fallthrough
  case 1:
    fmt.Println("check value is ", i)
  }
}
