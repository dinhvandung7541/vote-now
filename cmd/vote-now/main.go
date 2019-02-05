package main

import (
	"fmt"

	"github.com/dinhvandung7541/vote-now/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	//load env
	if err := godotenv.Load(); err != nil {
		fmt.Println("erorr")
		fmt.Println(err)
		panic(err)
	}

	// Setup DB
	db, closeDB := db.InitDBConnection()
	defer closeDB()

	if err := db.Ping(); err != nil {
		fmt.Println("Init DB error")
	} else {
		fmt.Println("Hello Docker")
	}

	// Echo instance
	e := echo.New()
	e = routes(e)
	// Start server
	e.Logger.Fatal(e.StartTLS(":8080", "cert.pem", "key.pem"))
	// e.Logger.Fatal(e.Start(":8080"))

}
