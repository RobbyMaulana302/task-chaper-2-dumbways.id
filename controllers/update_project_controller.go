package controllers

import (
	"context"
	"net/http"
	"taskgolang/connection"
	"time"

	"github.com/labstack/echo/v4"
)

func UpdateProject(c echo.Context) error {

	// mengambil value dari input form
	id := c.FormValue("id")
	ProjectName := c.FormValue("input-project-name")
	StartDate := c.FormValue("input-start-date")
	EndDate := c.FormValue("input-end-date")
	Description := c.FormValue("input-description")
	Technology := c.Request().Form["checkbox-technology"]
	layout := "2006-01-02"
	parse1, _ := time.Parse(layout, StartDate)
	parse2, _ := time.Parse(layout, EndDate)

	// Query update data ke database
	_, errQuery := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5, image = $6 WHERE id = $7", ProjectName, parse1, parse2, Description, Technology, "image-update-ke-1.png", id)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}