package middleware

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"taskgolang/connection"
	"taskgolang/models"
	"time"

	"github.com/labstack/echo/v4"
)

func ReadProject(c echo.Context) ([]models.Project){
	data, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_project ORDER by id DESC")

	var result []models.Project

	for data.Next(){
		var each =  models.Project{}
		
		err := data.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technology, &each.Image, &each.Author)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		parsingStartDate, _ := time.Parse("2006-01-02", each.StartDate.Format("2006-01-02"))

		parsingEndDate, _ := time.Parse("2006-01-02", each.StartDate.Format("2006-01-02"))

		distance := parsingEndDate.Sub(parsingStartDate).Milliseconds()
		days :=  math.Floor(float64(distance) / (1000 * 3600 * 24) )
		weeks :=  math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
		months :=  math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
		years :=  math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )

		parseDays := int(days)
		parseWeeks := int(weeks)
		parseMonths := int(months)
		parseYears := int(years)

		if days < 7 {
			each.Duration = strconv.Itoa(parseDays) + " days"
		} else if days < 30 {
			each.Duration = strconv.Itoa(parseWeeks) + " Weeks"
		} else if days < 365 {
			each.Duration = strconv.Itoa(parseMonths) + " Months"
		} else {
			each.Duration = strconv.Itoa(parseYears) + " Years"
		}

		result = append(result, each)
	}

	return result

}
