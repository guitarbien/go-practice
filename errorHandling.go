package main

import (
  "fmt"
  "errors"
)

func checkUsernameExists(username string) (bool, error) {
  if username == "foo" {
    return true, fmt.Errorf("username %s already exists", username)
  }

  if username == "bar" {
    return true, errors.New("username bar already exists")
  }

  return false, nil
}

func main() {
  if _, err := checkUsernameExists("foo"); err != nil {
    fmt.Println(err)
  }

  if _, err := checkUsernameExists("bar"); err != nil {
    fmt.Println(err)
  }
}
