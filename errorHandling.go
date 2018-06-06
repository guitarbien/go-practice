package main

import (
  "fmt"
)

func checkUsernameExists(username string) (bool, error) {
  if username == "foo" {
    return true, fmt.Errorf("username %s already exists", username)
  }

  return false, nil
}

func main() {
  if _, err := checkUsernameExists("foo"); err != nil {
    fmt.Println(err)
  }
}
