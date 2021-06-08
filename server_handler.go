package main

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

type ServerHandler struct {
	Redis *redis.Client
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{
		Redis: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func NewTestHandler() *ServerHandler {
	return &ServerHandler{
		Redis: redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}),
	}
}

func (h *ServerHandler) GetKey(c echo.Context) error {
	key := c.Param("key")
	val, err := h.Redis.Get(key).Result()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	res := map[string]interface{}{
		"key":   key,
		"value": val,
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ServerHandler) SetKey(c echo.Context) error {
	req := new(Request)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.Redis.Set(req.Key, req.Value, 0).Err(); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Saved successfully",
	})
}

func (h *ServerHandler) Search(c echo.Context) error {
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

	keys := h.Redis.Keys(pattern).Val()

	var searchResults []map[string]interface{}
	for _, key := range keys {
		val, err := h.Redis.Get(key).Result()
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
}
