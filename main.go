package main

import (
	"taskgolang/config"
	"taskgolang/controllers"

	"github.com/labstack/echo/v4"
)
func main() {

	// import koneksi database
	config.DatabaseConnect()

	// package echo
	e := echo.New()

	// file static direktori public
	e.Static("/public", "public")

	// routing
	// get
	e.GET("/", controllers.Home)
	e.GET("/add-project", controllers.FormProject)
	e.GET("/testimonials", controllers.Testimonials)
	e.GET("/contact", controllers.Contact)
	e.GET("/detail-project:id", controllers.DetailProject)
	
	// post
	e.POST("/add-project", controllers.AddProject)

	// port
	e.Logger.Fatal(e.Start(":5000"))
}

















