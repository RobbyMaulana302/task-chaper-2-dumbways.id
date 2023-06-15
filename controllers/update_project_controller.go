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
	author_id := c.FormValue("author")
	layout := "2006-01-02"
	parse1, _ := time.Parse(layout, StartDate)
	parse2, _ := time.Parse(layout, EndDate)
	image := c.Get("dataFile").(string)


	// Query update data ke database
	_, errQuery := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET name = $1, start_date = $2, end_date = $3, description = $4, technologies = $5, image = $6, author_id = $7 WHERE id = $8", ProjectName, parse1, parse2, Description, Technology, image, author_id, id)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}