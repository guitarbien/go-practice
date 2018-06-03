package main

import (
  "fmt"
  "strings"
)

func getUserListSQL(username, email string, gender int) string {
  sql := "select * from user"
  var where []string

  if username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", username))
  }

  if email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", email))
  }

  if gender != 0 {
    where = append(where, fmt.Sprintf("gender = '%d'", gender))
  }

  return sql + " where " + strings.Join(where, " or ")
}

func main() {
  fmt.Println(getUserListSQL("guitarbien", ""))
  fmt.Println(getUserListSQL("guitarbien", "guitarbien@gmail.com"))
}
