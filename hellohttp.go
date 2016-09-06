package main

import (
    "fmt"
    "log" 
    "database/sql"
	_ "github.com/lib/pq"
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
  
)

type Users struct {
  Users []User `json:"users"`
}


type User struct {
  ID int "json:id"
  Name  string "json:username"
  Email string "json:email"
  First string "json:first"
  Last  string "json:last"
}


func CreateUser(w http.ResponseWriter, r *http.Request) {

  NewUser := User{}
  NewUser.Name = r.FormValue("user")
  NewUser.Email = r.FormValue("email")
  NewUser.First = r.FormValue("first")
  NewUser.Last = r.FormValue("last")
  output, err := json.Marshal(NewUser)
  fmt.Println(string(output))
  if err != nil {
    fmt.Println("Something went wrong!")
  }

 

   db, err := sql.Open("postgres", "user=elliottchavis dbname=gohttp sslmode=disable")
  
  err = db.Ping()	

  if err != nil { 
    panic(err.Error()) 
  }


  rows, err := db.Query("INSERT INTO users (first_name, last_name, username, email) VALUES ('" + NewUser.First + "', '"+ NewUser.Last + "', '" + NewUser.Name +"', '" + NewUser.Email + "')" )
  
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(rows)

}



func GetUser(w http.ResponseWriter, r *http.Request) {

  log.Println("starting retrieval")
  start := 0
  limit := 10

  next := start + limit

  w.Header().Set("Pragma", "no-cache")
  w.Header().Set("Link", "<http://localhost:8080/api/users?start="+string(next)+"; rel=\"next\"")

  db, err := sql.Open("postgres", "user=elliottchavis dbname=gohttp sslmode=disable")
  
  err = db.Ping()	

  if err != nil { 
    panic(err.Error()) 
  }

  /* this is to select ONE user...
 err = db.QueryRow("select * from users where users_id=$1",id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email )
*/

  
/* this is to select ALL users...    */

  rows, err := db.Query("select * from users LIMIT 10")
  Response := Users{}

  for rows.Next() {

    user := User{}
      rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)

    Response.Users = append(Response.Users, user)
  }

  output,_ := json.Marshal(Response)
  fmt.Fprintln(w,string(output))

}


 func main() {

  gorillaRoute := mux.NewRouter()
  gorillaRoute.HandleFunc("/api/users", CreateUser).Methods("POST")
  gorillaRoute.HandleFunc("/api/users", GetUser).Methods("GET")  
  http.Handle("/", gorillaRoute)
  http.ListenAndServe(":8080", nil)
}




