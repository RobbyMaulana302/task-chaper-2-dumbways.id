package controllers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func FormProject(c echo.Context) error {
	var template, error = template.ParseFiles("views/add_project.html")

	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}

	return template.Execute(c.Response(), nil)
}