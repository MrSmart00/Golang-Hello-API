package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routing(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func routing(e *echo.Echo) {
	e.GET("/", hello)
	e.GET("/:name", greeting)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "hello " + c.Param("name")})
}
