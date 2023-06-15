package controllers

import (
	"context"
	"log"
	"net/http"
	"taskgolang/connection"
	"taskgolang/middleware"
	"taskgolang/models"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func FormLogin(c echo.Context) error {
	sess, errSession := session.Get("session", c)

	if errSession != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": errSession.Error()})
	}

	flash := map[string]interface {} {
		"flashStatus": sess.Values["status"],
		"flashMessage": sess.Values["message"],
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var tmpl, err = template.ParseFiles("views/login.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), flash)

}

func Login(c echo.Context) error {
	err := c.Request().ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	email := c.FormValue("inputEmail")
	password := c.FormValue("inputPassword")

	user := models.Users{}
	err = connection.Conn.QueryRow(context.Background(), "SELECT * FROM users WHERE email=$1", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return middleware.RedirectWithMessage(c, "Email Incorrect!", false, "/form-login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return  middleware.RedirectWithMessage(c, "Password Incorrect!", false, "/form-login")
	}

	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.ID
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}