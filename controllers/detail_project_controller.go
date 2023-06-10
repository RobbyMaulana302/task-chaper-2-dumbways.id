package controllers

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"taskgolang/config"
	"taskgolang/models"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

func DetailProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var result = models.Project{}

	errQuery := config.Conn.QueryRow(context.Background(), "SELECT name, start_date, end_date, description, technologies FROM tb_project WHERE id=$1", id).Scan(&result.ProjectName, &result.StartDate, &result.EndDate, &result.Description, &result.Technology)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}

	result.Date1 = result.StartDate.Format("02 Jan 2006")
	result.Date2 = result.EndDate.Format("02 Jan 2006")

	parsingDate1, errParsingDate1 := time.Parse("2006-01-02", result.StartDate.Format("2006-01-02"))
		
	if errParsingDate1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate1.Error()})
	}

	parsingDate2, errParsingDate2 := time.Parse("2006-01-02", result.EndDate.Format("2006-01-02"))

	if errParsingDate2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate2.Error()})
	}

	distance := parsingDate2.Sub(parsingDate1).Milliseconds()

	days := math.Floor(float64(distance) / (1000 * 3600 * 24) )
	weeks := math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
	months := math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
	years := math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )
	
	daysParsingInt := int(days)
	weeksParsingInt := int(weeks)
	monthsParsingInt := int(months)
	yearsParsingInt := int(years)

	if days < 7 {
		result.Duration = strconv.Itoa(daysParsingInt) + " days"
	} else if days < 30 {
		result.Duration = strconv.Itoa(weeksParsingInt) + " weeks"
	} else if days < 365 {
		result.Duration = strconv.Itoa(monthsParsingInt) + " months"
	} else {
		result.Duration = strconv.Itoa(yearsParsingInt) + " years"
	}

	data := map[string]interface{}{
		"Projects": result,
	}

	var template, errTemplate = template.ParseFiles("views/detail_project.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return template.Execute(c.Response(), data)
}