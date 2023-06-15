package main

import (
	"taskgolang/connection"
	"taskgolang/controllers"
	"taskgolang/middleware"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)
func main() {

	// import koneksi database
	connection.DatabaseConnect()
	
	// package echo
	e := echo.New()

	// file static direktori public
	e.Static("/public", "public")
	e.Static("/uploads", "uploads")

	// untuk menggunakan session echo
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	// routing
	// get
	e.GET("/", controllers.Home)
	e.GET("/add-project", controllers.FormProject)
	e.GET("/testimonials", controllers.Testimonials)
	e.GET("/contact", controllers.Contact)
	e.GET("/detail-project/:id", controllers.DetailProject)
	e.GET("/edit-project/:id", controllers.EditProject)

	// login
	e.GET("/form-login", controllers.FormLogin)
	e.POST("/login", controllers.Login)

	// logout
	e.POST("/logout", controllers.Logout)

	// register
	e.GET("/form-register", controllers.FormRegister)
	e.POST("/register", controllers.Register)
	
	// post
	e.POST("/add-project", middleware.UploadFile(controllers.AddProject))
	e.POST("/update-project", middleware.UploadFile(controllers.UpdateProject))
	e.POST("/delete-project/:id", controllers.DeleteProject)

	// port
	e.Logger.Fatal(e.Start(":5000"))
}




