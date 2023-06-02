package main

import (
   "net/http"

   "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat datang di dunia pemrograman web golang si golang petualang susah kodingnya seperti bahasa C# lala lala")
	})

	e.Logger.Fatal(e.Start(":5000"))
}