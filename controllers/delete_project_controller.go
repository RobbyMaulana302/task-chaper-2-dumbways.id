package controllers

import (
	"context"
	"net/http"
	"strconv"
	"taskgolang/config"

	"github.com/labstack/echo/v4"
)

func DeleteProject(c echo.Context) error {
	id, errParam := strconv.Atoi(c.Param("id"))

	if errParam != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParam.Error()})
	}

	_, errQuery := config.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/")

	
}