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

  urlParams   := mux.Vars(r)
  id       := urlParams["id"]
  ReadUser := User{}

  db, err := sql.Open("postgres", "user=elliottchavis dbname=gohttp sslmode=disable")
  
  err = db.Ping()	

  if err != nil { 
    panic(err.Error()) 
  }

   err = db.QueryRow("select * from users where users_id=$1",id).Scan(&ReadUser.ID, &ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email )
 switch {
      case err == sql.ErrNoRows:
              fmt.Fprintf(w,"No such user")
      case err != nil:
              log.Fatal(err)
  fmt.Fprintf(w, "Error")
      default:
        output, _ := json.Marshal(ReadUser)
        fmt.Fprintf(w,string(output))
  }
}

 func main() {

  gorillaRoute := mux.NewRouter()
  gorillaRoute.HandleFunc("/api/user/create", CreateUser).Methods("GET")
  gorillaRoute.HandleFunc("/api/user/read/{id:[0-9]+}", GetUser).Methods("GET")  
  http.Handle("/", gorillaRoute)
  http.ListenAndServe(":8080", nil)
}




