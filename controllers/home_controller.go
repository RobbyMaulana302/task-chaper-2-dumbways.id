package controllers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"os"
	"strconv"
	"taskgolang/config"
	"taskgolang/models"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	// sql untuk mendapatkan value dari tabel tb_project
	data, x := config.Conn.Query(context.Background(), "SELECT * FROM tb_project")
	if x != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", x)
	}

	var result []models.Project
	for data.Next() {
		var each = models.Project{}

		err := data.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technology, &each.Image)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}
		each.Date1 = each.StartDate.Format("2006-01-02")
		each.Date2 = each.EndDate.Format("2006-01-02")
		
		parsingDate1, errParsingDate1 := time.Parse("2006-01-02", each.Date1)
		
		if errParsingDate1 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate1.Error()})
		}

		parsingDate2, errParsingDate2 := time.Parse("2006-01-02", each.Date2)

		if errParsingDate2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate2.Error()})
		}

		distance := parsingDate2.Sub(parsingDate1).Milliseconds()

		days := math.Floor(float64(distance) / (1000 * 3600 * 24) )
		fmt.Println(days)
		weeks := math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
		months := math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
		years := math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )

		daysParsingInt := int(days)
		weeksParsingInt := int(weeks)
		monthsParsingInt := int(months)
		yearsParsingInt := int(years)

		if days < 7 {
			each.Duration = strconv.Itoa(daysParsingInt) + " days"
		} else if days < 30 {
			each.Duration = strconv.Itoa(weeksParsingInt) + " Weeks"
		} else if days < 365 {
			each.Duration = strconv.Itoa(monthsParsingInt) + " Months"
		} else {
			each.Duration = strconv.Itoa(yearsParsingInt) + " Years"
		}

		result = append(result, each)
	}


	blogs := map[string]interface{}{
		"Blogs": result,
	}
	
	var template, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return template.Execute(c.Response(), blogs)
}