package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	handler := NewServerHandler()
	e.GET("/get/:key", handler.GetKey)
	e.POST("/set", handler.SetKey)
	e.GET("/search", handler.Search)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
