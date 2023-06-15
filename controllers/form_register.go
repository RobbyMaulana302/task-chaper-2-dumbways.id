package controllers

import (
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func FormRegister(c echo.Context) error {

	tmpl, err := template.ParseFiles("views/register.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"mesage": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

