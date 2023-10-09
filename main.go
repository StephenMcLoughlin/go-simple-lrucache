package main

import (
	"net/http"

	"github.com/StephenMcLoughlin/go-simple-lrucache/cache"
	"github.com/labstack/echo"
)

var lruCache = cache.NewLRUCache()

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func addItem(c echo.Context) error {
	request := new(Data)

	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON")
	}

	key := request.Key
	value := request.Value

	lruCache.Put(key, value)
	lruCache.PrintCurrentCache()

	return c.String(http.StatusOK, "Added")
}

func getItem(c echo.Context) error {
	key := c.Param("key")

	cacheItem := lruCache.Get(key)

	if cacheItem == "" {
		return echo.NewHTTPError(http.StatusNotFound, "Item not found in cache")
	}

	lruCache.PrintCurrentCache()
	return c.String(http.StatusOK, cacheItem)
}

func main() {
	if lruCache == nil {
		panic("Error: Cache failed to initialised")
	}

	e := echo.New()

	e.POST("/cache/item", addItem)
	e.GET("/cache/item/:key", getItem)

	e.Start(":8080")
}
