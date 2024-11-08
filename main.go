package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("Init provider successful")
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{
		"Service ECHO": "STARTED!",
	}).Info("Let's roll")

	e.GET("/", hello)

	e.GET("/hi/:name", func(c echo.Context) error {
		name := "Hi " + c.Param("name")
		return c.String(http.StatusOK, name)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8082"))
}

// Handler
func hello(c echo.Context) error {
	u := &User{
		Name:  "Jhon",
		Email: "jhon@labstack.com",
	}

	print("hello")

	return c.JSON(http.StatusOK, u)
}
