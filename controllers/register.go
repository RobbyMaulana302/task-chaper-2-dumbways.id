package controllers

import (
	"context"
	"log"
	"taskgolang/connection"
	"taskgolang/middleware"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	err := c.Request().ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	name := c.FormValue("inputName")
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err = connection.Conn.Exec(context.Background(), "INSERT INTO users(name, email, password) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		middleware.RedirectWithMessage(c, "Register failed, please try again", false, "/form-register")
	}

	return middleware.RedirectWithMessage(c, "Register success", true, "/form-login")
}