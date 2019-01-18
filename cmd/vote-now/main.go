package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	//load env
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Setup DB
	db, closeDB := initDB()
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
	e.Logger.Fatal(e.Start(":8080"))

}
