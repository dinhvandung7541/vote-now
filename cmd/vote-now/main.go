package main

import (
    "fmt"
    "net/http"
    // "database/sql"
    // "os"

    // _ "github.com/lib/pq"
)

func main() {
    // initDb()
    // defer db.Close()

    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!")
}

