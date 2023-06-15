package controllers

import (
	"context"
	"fmt"
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

func Home(c echo.Context) error {
	sess, errSession := session.Get("session", c)
	var result []models.Project
	var userData = models.SessionData{}

	if sess.Values["id"] == nil {
		// query untuk menampilkan seluruh data pada table
		data, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_project.name, start_date, end_date,description, technologies, image, users.name	AS author FROM tb_project JOIN users ON tb_project.author_id = users.id ORDER by tb_project.id DESC")
		
		// kondisi jika query error
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
		}

		// deklasrasia array dari model struct project

		// looping data 
		for data.Next() {
			var each = models.Project{}

			// scan, membaca setiap baris dari value database
			err := data.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technology, &each.Image, &each.Author)

			// kondisi jika terjadi pada err baris data
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
			}

			// mengubah format date
			each.Date1 = each.StartDate.Format("2006-01-02")
			each.Date2 = each.EndDate.Format("2006-01-02")
			
			// parsing date
			parsingDate1, errParsingDate1 := time.Parse("2006-01-02", each.Date1)
			
			if errParsingDate1 != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate1.Error()})
			}

			// parsing date
			parsingDate2, errParsingDate2 := time.Parse("2006-01-02", each.Date2)

			if errParsingDate2 != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate2.Error()})
			}

			// menghitung jarak tanggal
			distance := parsingDate2.Sub(parsingDate1).Milliseconds()

			// menghitung agar menjadi bilangan bulat ke bawah dari jarak dan mendapatkan nilai hari, minggu, bulan dan tahun
			days := math.Floor(float64(distance) / (1000 * 3600 * 24) )
			fmt.Println(days)
			weeks := math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
			months := math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
			years := math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )

			// parsing float64 menjadi int
			daysParsingInt := int(days)
			weeksParsingInt := int(weeks)
			monthsParsingInt := int(months)
			yearsParsingInt := int(years)

			// pengkondisian durasi
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


		if errSession != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{"message": errSession.Error()})
		}

		if sess.Values["isLogin"] != true {
			userData.IsLogin = false
		} else {
			userData.IsLogin = true
			userData.Name = sess.Values["name"].(string)
		}
		// membuat deklarasi dan menampung value dari result ke dalam map interface
	} else {
		// query untuk menampilkan seluruh data pada table
		data, errQuery := connection.Conn.Query(context.Background(), "SELECT tb_project.id, tb_project.name, start_date, end_date,description, technologies, image, users.name	AS author FROM tb_project JOIN users ON tb_project.author_id = users.id WHERE author_id=$1 ORDER by tb_project.id DESC", sess.Values["id"])
		
		// kondisi jika query error
		if errQuery != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": errQuery.Error()})
		}
	
		// deklasrasia array dari model struct project
	
		// looping data 
		for data.Next() {
			var each = models.Project{}
	
			// scan, membaca setiap baris dari value database
			err := data.Scan(&each.ID, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technology, &each.Image, &each.Author)
	
			// kondisi jika terjadi pada err baris data
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
			}
	
			// mengubah format date
			each.Date1 = each.StartDate.Format("2006-01-02")
			each.Date2 = each.EndDate.Format("2006-01-02")
				
			// parsing date
			parsingDate1, errParsingDate1 := time.Parse("2006-01-02", each.Date1)
				
			if errParsingDate1 != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate1.Error()})
			}
	
			// parsing date
			parsingDate2, errParsingDate2 := time.Parse("2006-01-02", each.Date2)
	
			if errParsingDate2 != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": errParsingDate2.Error()})
			}
	
			// menghitung jarak tanggal
			distance := parsingDate2.Sub(parsingDate1).Milliseconds()
	
			// menghitung agar menjadi bilangan bulat ke bawah dari jarak dan mendapatkan nilai hari, minggu, bulan dan tahun
			days := math.Floor(float64(distance) / (1000 * 3600 * 24) )
				fmt.Println(days)
			weeks := math.Floor(float64(distance) / (1000 * 3600 * 24 * 7) )
			months := math.Floor(float64(distance) / (1000 * 3600 * 24 * 30) )
			years := math.Floor(float64(distance) / (1000 * 3600 * 24 * 365) )
	
			// parsing float64 menjadi int
			daysParsingInt := int(days)
			weeksParsingInt := int(weeks)
			monthsParsingInt := int(months)
			yearsParsingInt := int(years)
	
			// pengkondisian durasi
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
	
	
			if errSession != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"message": errSession.Error()})
			}
	
			if sess.Values["isLogin"] != true {
				userData.IsLogin = false
			} else {
				userData.IsLogin = true
				userData.Name = sess.Values["name"].(string)
		}
	}
		
	// membuat deklarasi dan menampung value dari result ke dalam map interface
	datas := map[string]interface{}{
		"Blogs": result,
		"FlashStatus" : sess.Values["status"],
		"FlasMessage" : sess.Values["message"],
		"DataSession" : userData,
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())
	
	
	// mendapatkan halaman yang akan ditampilkan
	var template, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return template.Execute(c.Response(), datas)
	
}