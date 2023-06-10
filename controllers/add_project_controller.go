package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func AddProject(c echo.Context) error {
	// ProjectName := c.FormValue("input-project-name")
	StartDate := c.FormValue("input-start-date")
	EndDate := c.FormValue("input-end-date")
	// Description := c.FormValue("input-description")
	// c.Request().ParseForm()
	// var Technology = c.Request().Form["checkbox-technology"]
	// for _, v := range Technology {
	// 	fmt.Println("v:", v)
	// }
	layout := "2006-01-02"
	parse1, _ := time.Parse(layout, StartDate)
	parse2, _ := time.Parse(layout, EndDate)

	// println("Title :", ProjectName)
	fmt.Println("StartDate :", parse1.Nanosecond())
	println("EndDate :", parse2.Nanosecond())
	// println("Description :", Description)
	// println("Technology :", Technology)

	// var newProject = Project{
	// 	ProjectName: ProjectName,
	// }

	return c.Redirect(http.StatusMovedPermanently, "/")
}
