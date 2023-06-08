package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

func DetailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"Id":          id,
		"Title":       "Dumbways Web App",
		"Date":        "12 Jan 2021 - 11 Feb 2021",
		"Duration":    "1 Month",
		"Description": "Lorem ipsum dolor sit amet consectetur adipisicing elit Suscipit qui natus perspiciatis tenetur, exercitationem, expedita rerum nam illo tempora eos doloribus? Cupiditate molestiae corrupti eveniet,facere asperiores doloribus praesentium voluptate beatae velit quis repellat fugiat deleniti a inventore voluptatibus sit delectus quidem at, expedita eligendi facilis assumenda sunt iusto? Vitae, rerum assumenda. Et saepe, dolor aliquid impedit, hic id eum nemo rerum dignissimos, quo temporibus sunt cum reprehenderit nostrum facere quasi voluptatum. Temporibus explicabo eum ipsa dolores! Architecto quisquam laudantium facere sapiente eius voluptas molestias quo repudiandae labore reiciendis maiores nihil earum illo, ipsa aspernatur, consequatur odit tempora? Aut, corrupti.",
	}

	var template, error = template.ParseFiles("views/detailProject.html")

	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": error.Error()})
	}

	return template.Execute(c.Response(), data)
}