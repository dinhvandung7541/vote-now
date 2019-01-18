package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// List to get all question and option
func List(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

// AddOption add 1 option to question
func AddOption(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

//RemoveOption remove 1 option in question
func RemoveOption(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

//CreateQuestion add 1 question
func CreateQuestion(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

// RemoveQuestion remove question with exact ID
func RemoveQuestion(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

//Vote user chose option in Question
func Vote(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

//Unvote user remove option that vote before
func Unvote(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}
