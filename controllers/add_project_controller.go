package controllers

import (
	"context"
	"net/http"
	"taskgolang/config"
	"time"

	"github.com/labstack/echo/v4"
)

func AddProject(c echo.Context) error {

	// mengambil nilai dari inputan form
	ProjectName := c.FormValue("input-project-name")
	StartDate := c.FormValue("input-start-date")
	EndDate := c.FormValue("input-end-date")
	Description := c.FormValue("input-description")
	Technology := c.Request().Form["checkbox-technology"]

	// parsing format tanggal
	layout := "2006-01-02"
	parse1, _ := time.Parse(layout, StartDate)
	parse2, _ := time.Parse(layout, EndDate)

	// query tambah data ke database
	_, errQuery := config.Conn.Exec(context.Background(), "INSERT INTO tb_project (name, start_date, end_date, description, technologies, image) VALUES ($1, $2, $3, $4, $5, $6)", ProjectName, parse1, parse2, Description, Technology, "image-1.png" )

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")
}
