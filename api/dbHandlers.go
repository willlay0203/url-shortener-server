package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func postNewShortenedUrl(code string, link string) error {

	// Load connection string from .env file, opens the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)

	}

	// Open a connection to PlanetScale
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// SQL query
	res, err := db.Exec("INSERT INTO Urls(code, link) VALUES (?, ?)", code, link)

	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	// Check if it has worked
	added, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	if added != 1 {
		log.Fatalf("expected new row to be added, affected %d", added)
	}
	// Close the sql
	defer db.Close()

	fmt.Printf("\ncode: %v, link: %v", code, link)

	return nil
}
