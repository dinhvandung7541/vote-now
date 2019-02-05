package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dinhvandung7541/vote-now/db"
	"github.com/dinhvandung7541/vote-now/models"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/labstack/echo"
)

// List to get all question and option
func List(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "List all vote!---------")
}

// AddOption add 1 option to question
func AddOption(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	userID := claims["userID"].(float64)
	questionID := claims["questionID"].(float64)

	option := new(models.Option)
	option.UserID.Int = int(userID)
	option.QuestionID.Int = int(questionID)
	if err := c.Bind(option); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "can't parse data")
	}

	err := option.Insert(context.Background(), db.DB, boil.Infer())
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusOK, "Can't create question")
	}

	return c.String(http.StatusOK, option.Content)
}

//RemoveOption remove 1 option in question
func RemoveOption(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	userID := claims["userID"].(float64)
	questionID := claims["questionID"].(float64)

	option := new(models.Option)
	option.UserID.Int = int(userID)
	option.QuestionID.Int = int(questionID)
	if err := c.Bind(option); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "can't parse data")
	}

	err := option.Insert(context.Background(), db.DB, boil.Infer())
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusOK, "Can't create question")
	}

	return c.String(http.StatusOK, option.Content)
	return c.String(http.StatusOK, "Remove option--------")
}

//CreateQuestion add 1 question
func CreateQuestion(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	userID := claims["userID"].(float64)

	question := new(models.Question)
	question.UserID.Int = int(userID)
	if err := c.Bind(question); err != nil {
		log.Fatal(err)
		return c.String(http.StatusOK, "can't parse data")
	}

	err := question.Insert(context.Background(), db.DB, boil.Infer())
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusOK, "Can't create question")
	}

	return c.String(http.StatusOK, question.Content)
}

// RemoveQuestion remove question with exact ID
func RemoveQuestion(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Removw question--------")
}

//Vote user chose option in Question
func Vote(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Vote--------")
}

//Unvote user remove option that vote before
func Unvote(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Unvote-----")
}
