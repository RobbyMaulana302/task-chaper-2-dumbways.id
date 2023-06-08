package controllers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func Contact(c echo.Context) error {
	var template, error = template.ParseFiles("views/contact.html")

	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": error.Error()})
	}

	return template.Execute(c.Response(), nil)
}