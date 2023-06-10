package controllers

import (
	"context"
	"net/http"
	"strconv"
	"taskgolang/config"
	"taskgolang/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

func EditProject(c echo.Context) error {
	id, errParse := strconv.Atoi(c.Param("id"))

	if errParse != nil {
		c.JSON(http.StatusInternalServerError, map[string]string {"message": errParse.Error()})
	}

	var result = models.Project{}
	errQuery := config.Conn.QueryRow(context.Background(), "SELECT * FROM tb_project WHERE id=$1", id).Scan(&result.ID, &result.ProjectName, &result.StartDate, &result.EndDate, &result.Description, &result.Technology, &result.Image)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParse.Error()})
	}

	parseLayout := "2006-01-02"
	result.Date1 = result.StartDate.Format(parseLayout)
	result.Date2 = result.EndDate.Format(parseLayout)
	
	data := map[string]interface{}{
		"Data": result,
	}

	var template, errTemplate = template.ParseFiles("views/edit_project.html")
	if errTemplate != nil {
		return c.JSON(http.StatusMovedPermanently, map[string]string{"message": errTemplate.Error()})
	}

	return template.Execute(c.Response(), data)

}