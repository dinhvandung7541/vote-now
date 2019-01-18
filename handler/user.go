package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Signup create new user
func Signup(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

func Login(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

func ChangePassword(c echo.Context) error {
	//To do
	return c.String(http.StatusOK, "Hello, World!")
}

func isUserExist(username string) (bool, error) {
	//To do
	return false, false
}

func save(user models.user) {
	//To do
}

func validateData(c echo.Context) {

}
