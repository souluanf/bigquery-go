package models

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
)

type ExampleDTO struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address struct {
		Street   bigquery.NullString `json:"street"`
		Number   bigquery.NullInt64  `json:"number"`
		District civil.Date          `json:"district"`
		City     bigquery.NullString `json:"city"`
	} `json:"address"`
}
