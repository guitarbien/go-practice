package main

import (
  "fmt"
)

type errUserNameExist struct {
  UserName string
}

func (e errUserNameExist) Error() string {
  return fmt.Sprintf("username %s already exists", e.UserName)
}

func checkUsernameExists(username string) (bool, error) {
  if username == "foo" {
    return true, errUserNameExist{UserName: username}
  }

  return false, nil
}

func main() {
  if _, err := checkUsernameExists("foo"); err != nil {
    fmt.Println(err)
  }
}
