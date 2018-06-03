package main

import (
  "fmt"
  "strings"
)

func getUserListSQL(username, email string) string {
  sql := "select * from user"
  var where []string

  if username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", username))
  }

  if email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", email))
  }

  return sql + " where " + strings.Join(where, " or ")
}

type searchOperation struct {
  username string
  email    string
  gender   int
}

func getUserListOperationSQL(operator searchOperation) string {
  sql := "select * from user"
  var where []string

  if operator.username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", operator.username))
  }

  if operator.email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", operator.email))
  }

  if operator.gender != 0 {
    where = append(where, fmt.Sprintf("gender = '%d'", operator.gender))
  }

  return sql + " where " + strings.Join(where, " or ")
}

func main() {
  fmt.Println(getUserListSQL("guitarbien", ""))
  fmt.Println(getUserListOperationSQL(searchOperation{
    username: "guitarbien",
  }))

  fmt.Println(getUserListSQL("guitarbien", "guitarbien@gmail.com"))
  fmt.Println(getUserListOperationSQL(searchOperation{
    username: "guitarbien",
    email:    "guitarbien@gmail.com",
  }))

  fmt.Println(getUserListOperationSQL(searchOperation{
    gender: 1,
  }))
}
