package main

import (
	"cloud.google.com/go/bigquery"
	"context"
	"encoding/json"
	"fmt"
	"github.com/souluanf/bigquery-go/models"
	"google.golang.org/api/iterator"
	"log"
)

func main() {
	projectID := "="
	dataSetName := ""
	tableName := ""
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()
	data := fmt.Sprintf("SELECT * FROM `%s.%s.%s`", projectID, dataSetName, tableName)
	rows, err := query(ctx, client, data)
	if err != nil {
		log.Fatal(err)
	}
	if err := printResults(rows, models.ExampleDTO{}); err != nil {
		log.Fatal(err)
	}

}

func query(ctx context.Context, client *bigquery.Client, queryString string) (*bigquery.RowIterator, error) {
	query := client.Query(queryString)
	return query.Read(ctx)
}

func printResults(iter *bigquery.RowIterator, row models.ExampleDTO) error {
	for {
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		marshal, err := json.Marshal(row)
		fmt.Println(string(marshal))
	}
}
