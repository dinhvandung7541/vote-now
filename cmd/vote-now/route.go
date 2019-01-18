package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
)

func routes(e *echo.Echo) *echo.Echo {
	// Middleware
	e.Use(mw.NewCustomContext(s))
	e.Use(middleware.Logger())
	// e.Use(mw.WithJWT())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	// User route
	user := e.Group("/users")
	user.POST("/signup", handler.Signup)
	user.POST("/login", handler.Login)
	user.POST("/changePassword", handler.ChangePassword)

	//vote
	vote := e.Group("/votes")
	vote.GET("/list", handler.List)
	vote.POST("/vote", handler.Vote)
	vote.POST("/unvote", handler.Unvote)
	vote.POST("/create", handler.CreateQuestion)
	vote.POST(("/addOption", handler.AddOption))
	vote.POST("/removeQuestion", handler.RemoveQuestion)
	vote.POST("/removeOption", handler.RemoveOption)

	return e
}
