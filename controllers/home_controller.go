package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"taskgolang/config"
	"taskgolang/models"
	"text/template"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	// sql untuk mendapatkan value dari tabel tb_project
	data, x := config.Conn.Query(context.Background(), "SELECT * FROM tb_project")
	if x != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", x)
	}

	var result []models.Project
	for data.Next() {
		var each = models.Project{}

		err := data.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technology, &each.Image)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}
		each.Date1 = each.StartDate.Format("2 January 2006")
		each.Date2 = each.EndDate.Format("2 January 2006")

		result = append(result, each)
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}
	
	var template, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return template.Execute(c.Response(), blogs)
}