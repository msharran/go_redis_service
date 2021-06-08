package main

import (
	"go_redis_app/promexporter"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	e.Use(promexporter.PrometheusMiddleware())

	// Routes
	handler := NewServerHandler()
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/get/:key", handler.GetKey)
	e.POST("/set", handler.SetKey)
	e.GET("/search", handler.Search)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
