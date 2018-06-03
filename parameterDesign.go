package main

import (
  "fmt"
  "strings"
)

func getUserListSQL(username, email string) string {
  sql := "select * from user"
  where := []string{}

  if username != "" {
    where = append(where, fmt.Sprintf("username = '%s'", username))
  }

  if email != "" {
    where = append(where, fmt.Sprintf("email = '%s'", email))
  }

  return sql + " where " + strings.Join(where, " or ")
}

func main() {
  fmt.Println(getUserListSQL("guitarbien", ""))
  fmt.Println(getUserListSQL("guitarbien", "guitarbien@gmail.com"))
}
