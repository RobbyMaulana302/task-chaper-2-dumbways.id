package models

import "time"

type Project struct {
	ID          int
	ProjectName string
	StartDate   time.Time
	EndDate     time.Time
	Duration 	string
	Description string
	Technology  []string
	Image       string
	Date1       string
	Date2       string
}