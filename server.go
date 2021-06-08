package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	// Route => handler
	e.GET("/get/:key", func(c echo.Context) error {
		key := c.Param("key")
		val, err := rdb.Get(key).Result()
		if err != nil {
			return c.JSON(400, err.Error())
		}
		res := map[string]interface{}{
			"key":   key,
			"value": val,
		}
		return c.JSON(http.StatusOK, res)
	})

	e.POST("/set", func(c echo.Context) error {
		req := new(Request)
		if err = c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		err = rdb.Set(req.Key, req.Value, 0).Err()
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(http.StatusCreated, "Success")
	})

	e.GET("/search", func(c echo.Context) error {
		prefix := c.QueryParam("prefix")
		suffix := c.QueryParam("suffix")
		pattern := ""

		if prefix == "" && suffix == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "prefix or suffix query is missing",
			})
		} else if prefix != "" && suffix != "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Only prefix or suffix query is needed",
			})
		} else if prefix != "" {
			pattern = prefix + "*"
		} else {
			pattern = "*" + suffix
		}

		keys := rdb.Keys(pattern).Val()

		var searchResults []map[string]interface{}
		for _, key := range keys {
			val, err := rdb.Get(key).Result()
			if err != nil {
				return c.JSON(400, err.Error())
			}
			res := map[string]interface{}{
				"key":   key,
				"value": val,
			}
			searchResults = append(searchResults, res)
		}
		return c.JSON(http.StatusOK, searchResults)
	})

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
