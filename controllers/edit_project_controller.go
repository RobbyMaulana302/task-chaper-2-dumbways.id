package controllers

import (
	"context"
	"net/http"
	"strconv"
	"taskgolang/connection"
	"taskgolang/models"
	"text/template"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func EditProject(c echo.Context) error {
	// parsing id dari query params
	id, errParse := strconv.Atoi(c.Param("id"))

	if errParse != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string {"message": errParse.Error()})
	}

	var result = models.Project{}

	// Query mengambil 1 data dari database
	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&result.ID, &result.ProjectName, &result.StartDate, &result.EndDate, &result.Description, &result.Technology, &result.Image, &result.Author_id)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParse.Error()})
	}

	// mengubah format tanggal
	parseLayout := "2006-01-02"
	result.Date1 = result.StartDate.Format(parseLayout)
	result.Date2 = result.EndDate.Format(parseLayout)

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
	
	// membuat data baru map interface dari data result
	datas := map[string]interface{}{
		"Data": result,
		"DataSession" : userData,
		"Author_id": author_id,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	var template, errTemplate = template.ParseFiles("views/edit_project.html")
	if errTemplate != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"message": errTemplate.Error()})
	}

	return template.Execute(c.Response(), datas)

}