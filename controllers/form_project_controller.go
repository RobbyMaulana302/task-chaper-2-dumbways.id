package controllers

import (
	"net/http"
	"taskgolang/models"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func FormProject(c echo.Context) error {
	// mendapatkan halaman yang akan ditampilkan
	var template, error = template.ParseFiles("views/add_project.html")

	var userData = models.SessionData{}

	sess, errSession := session.Get("session", c)

	if errSession != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": errSession.Error()})
	}

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = true
		userData.Name = sess.Values["name"].(string)
	}

	author_id := sess.Values["id"]

	// membuat deklarasi dan menampung value dari result ke dalam map interface
	datas := map[string]interface{}{
		"DataSession" : userData,
		"Author_id": author_id,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}

	return template.Execute(c.Response(), datas)
}