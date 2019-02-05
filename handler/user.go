package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/volatiletech/sqlboiler/boil"

	"tutorial/vote-now/db"
	"tutorial/vote-now/models"

	// "github.com/dinhvandung7541/vote-now/models"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// Signup create new user
func Signup(c echo.Context) error {
	//To do
	username := c.FormValue("username")
	password := c.FormValue("password")
	if isExist, err := isUserExist(username); isExist != false || err != nil {
		return c.String(http.StatusOK, "Some thing error!")
	}
	newUser := &models.User{Username: username, Password: password}
	err := newUser.Insert(context.Background(), db.DB, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK, "Signup successful "+username)
}

func Signin(c echo.Context) error {
	//To do
	username := c.FormValue("username")
	password := c.FormValue("password")

	fmt.Println("Username: " + username + " Password: " + password)

	user, err := models.Users(qm.Where("username= ? ", username)).One(context.Background(), db.DB)
	if err != nil || user == nil || user.Password != password {
		fmt.Println("---------------------------------------")
		return c.String(http.StatusOK, "Something error!")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func ChangePassword(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	newPassword := c.FormValue("new_pasword")

	user, err := models.Users(qm.Where("username= ? ", username)).One(context.Background(), db.DB)
	if err != nil || user == nil || user.Password != password {
		fmt.Println(err)
		return c.String(http.StatusOK, "Something error!")
	}
	user.Password = newPassword
	_, err = user.Update(context.Background(), db.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, "Something error!")
	}
	return c.String(http.StatusOK, "Change password success!")
}

func isUserExist(username string) (bool, error) {
	isExist, err := models.Users(qm.Where("username= ? ", username)).Exists(context.Background(), db.DB)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return false, err
	}
	return isExist, nil
}

func validateData(c echo.Context) {

}
