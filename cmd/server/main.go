package main

import (
    "fmt"
    "net/http"
    "database/sql"
    "os"

    _ "github.com/lib/pq"
)

var db *sql.DB

const (
    dbhost = "localhost"
    dbport = "5432"
    dbuser = "postgres"
    dbpass = "example"
    dbname = "vote-now"

)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
    initDb()
    defer db.Close()

    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}

func initDb() {
    config := dbConfig()
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
        "password=%s dbname=%s sslmode=disable",
        config[dbhost], config[dbport],
        config[dbuser], config[dbpass], config[dbname])

    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
    conf := make(map[string]string)
    host, ok := os.LookupEnv(dbhost)
    if !ok {
        panic("DBHOST environment variable required but not set")
    }
    port, ok := os.LookupEnv(dbport)
    if !ok {
        panic("DBPORT environment variable required but not set")
    }
    user, ok := os.LookupEnv(dbuser)
    if !ok {
        panic("DBUSER environment variable required but not set")
    }
    password, ok := os.LookupEnv(dbpass)
    if !ok {
        panic("DBPASS environment variable required but not set")
    }
    name, ok := os.LookupEnv(dbname)
    if !ok {
        panic("DBNAME environment variable required but not set")
    }
    conf[dbhost] = host
    conf[dbport] = port
    conf[dbuser] = user
    conf[dbpass] = password
    conf[dbname] = name
    return conf
}