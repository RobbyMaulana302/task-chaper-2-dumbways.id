package controllers

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"taskgolang/connection"
	"taskgolang/models"
	"text/template"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func DetailProject(c echo.Context) error {
	// menangkap id dari param lalu merubah menjadi int
	id, _ := strconv.Atoi(c.Param("id"))

	// membuat variabel dari model project
	var result = models.Project{}

	// query menampilkan 1 data dari database
	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT name, start_date, end_date, description, technologies FROM tb_project WHERE id=$1", id).Scan(&result.ProjectName, &result.StartDate, &result.EndDate, &result.Description, &result.Technology)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
	}

	// merubah format tanggal
	result.Date1 = result.StartDate.Format("02 Jan 2006")
	result.Date2 = result.EndDate.Format("02 Jan 2006")

	// parsing tanggal start date
	parsingDate1, errParsingDate1 := time.Parse("2006-01-02", result.StartDate.Format("2006-01-02"))
		
	if errParsingDate1 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate1.Error()})
	}

	// parsing tanggal end date
	parsingDate2, errParsingDate2 := time.Parse("2006-01-02", result.EndDate.Format("2006-01-02"))

	if errParsingDate2 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate2.Error()})
	}

	// menghitung jarak tanggal dan merubahnya ke milisecond
	distance := parsingDate2.Sub(parsingDate1).Milliseconds()

	// menghitung tanggal hari, minggu, bulang, dan tahun
	days := math.Floor(float64(distance) / (1000 * 3600 * 24) )
	weeks := math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
	months := math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
	years := math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )
	
	// parsing float menjadi int
	daysParsingInt := int(days)
	weeksParsingInt := int(weeks)
	monthsParsingInt := int(months)
	yearsParsingInt := int(years)

	// pengkondisian jarak tanggal
	if days < 7 {
		result.Duration = strconv.Itoa(daysParsingInt) + " days"
	} else if days < 30 {
		result.Duration = strconv.Itoa(weeksParsingInt) + " weeks"
	} else if days < 365 {
		result.Duration = strconv.Itoa(monthsParsingInt) + " months"
	} else {
		result.Duration = strconv.Itoa(yearsParsingInt) + " years"
	}

	var userData = models.SessionData{}

	sess, errSession := session.Get("session", c)

	if errSession != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": errSession.Error()})
	}

	if sess.Values["isLogin"] != true {
		userData.IsLogin = false
	} else {
		userData.IsLogin = true
		userData.Name = sess.Values["name"].(string)
	}

	// membuat data baru dari data result
	datas := map[string]interface{}{
		"Projects": result,
		"DataSession" : userData,
	}

	// mendapatkan halaman yang akan ditampilkan
	var template, errTemplate = template.ParseFiles("views/detail_project.html")

	if errTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": errTemplate.Error()})
	}

	return template.Execute(c.Response(), datas)
}