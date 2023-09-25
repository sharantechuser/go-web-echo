package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(loggingMiddleware)

	// Routes
	e.GET("/", hello)
	// Routes with query params
	e.GET("/helloworld", helloWorld)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	log.Println("hello...")
	return c.String(http.StatusOK, "Hello!")
}

func helloWorld(c echo.Context) error {
	// Get team and member from the query string
	hello := c.QueryParam("hello")
	world := c.QueryParam("world")
	return c.String(http.StatusOK, hello+", "+world)
}

func loggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Request received for")
		return next(c)
	}
}
