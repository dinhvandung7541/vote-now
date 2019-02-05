package main

import (
	// "github.com/dinhvandung7541/vote-now/handler"
	"tutorial/vote-now/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// type Template struct {
// 	templates *template.Template
// }

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func routes(e *echo.Echo) *echo.Echo {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	// User route
	user := e.Group("/users")
	user.POST("/signup", handler.Signup)
	user.POST("/signin", handler.Signin)
	user.POST("/change_password", handler.ChangePassword)

	e.GET("/list", handler.List)

	//vote
	vote := e.Group("/votes")
	vote.Use(middleware.JWT([]byte("secret")))
	vote.POST("/vote", handler.Vote)
	vote.POST("/unvote", handler.Unvote)
	vote.POST("/create", handler.CreateQuestion)
	vote.POST("/addOption", handler.AddOption)
	vote.POST("/removeQuestion", handler.RemoveQuestion)
	vote.POST("/removeOption", handler.RemoveOption)

	return e
}
