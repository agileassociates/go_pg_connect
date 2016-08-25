package main

import (
    "fmt"
    "database/sql"

	_ "github.com/lib/pq"

  
)


func main() {

  
  db, err := sql.Open("postgres", "user=elliottchavis dbname=gohttp sslmode=disable")
  
  err = db.Ping()	

  if err != nil { 
    panic(err.Error()) 
  }

  rows, err := db.Query("INSERT INTO users (first_name, last_name, username, email) VALUES ('bob', 'jones', 'bj', 'bj@gmail')")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(rows)

}




