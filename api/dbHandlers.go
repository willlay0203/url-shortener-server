package api

import (
	"fmt"
	"log"
	"server/lib"
	"server/types"

	_ "github.com/go-sql-driver/mysql"
)

func postNewShortenedUrl(code string, link string) error {

	db := lib.OpenDb()

	// SQL query
	res, err := db.Exec("INSERT INTO Urls(code, link) VALUES (?, ?)", code, link)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Check if it has worked
	_, err = res.RowsAffected()

	if err != nil {
		return err
	}

	// Close the sql
	defer db.Close()

	fmt.Printf("\ncode: %v, link: %v", code, link)

	return nil
}

func getUrl(code string) (string, error) {
	db := lib.OpenDb()

	// SQL query
	url := types.ShortenedUrl{}
	err := db.QueryRow("SELECT * FROM Urls WHERE code = ?", code).Scan(&url.Code, &url.Link)

	if err != nil {
		return "", err
	}

	if url.Code == "" || url.Link == "" {
		return "Code and link is missing", err
	}

	// Close the sql
	defer db.Close()

	return url.Link, nil
}
