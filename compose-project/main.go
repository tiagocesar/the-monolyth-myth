package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbHost   = "localhost"
	dbUser   = "usuario"
	dbPass   = "senha"
	dbSchema = "postgres"
)

type Pageview struct {
	Id        int       `db:"id"`
	User      int       `db:"userId"`
	Url       string    `db:"url"`
	Timestamp time.Time `db:"timestamp"`
}

func main() {
	// Building our connection string to connect to the database
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbSchema)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	// Total rows on the "Pageviews" table
	var totalRows int
	if err := db.QueryRow(`SELECT COUNT(id) AS total FROM "Pageviews"`).Scan(&totalRows); err != nil {
		log.Fatal(err)
	}
	log.Printf("Total rows on the Pageviews table: %d\n", totalRows)

	// Adding a new row for user '3'
	var insertedId int
	query := `
		INSERT INTO "Pageviews" ("userId", url)
	    VALUES ($1, $2) RETURNING id`
	if err := db.QueryRow(query, 3, "googl.com").Scan(&insertedId); err != nil {
		log.Fatal(err)
	}
	log.Printf("New row added to Pageviews table with id %d\n", insertedId)

	// Ooooops, the URL in the previous block was wrong, let's fix it
	query = `
		UPDATE "Pageviews"
		   SET url = $1
		 WHERE id = $2`
	if _, err := db.Exec(query, "google.com", insertedId); err != nil {
		log.Fatal(err)
	}

	// Now let's get all data for user 3
	query = `SELECT id, "userId", url, timestamp FROM "Pageviews" WHERE "userId" = $1`
	rows, err := db.Query(query, 3)
	if err != nil {
		log.Fatal(err)
	}

	var results []Pageview
	for rows.Next() {
		var row Pageview
		if err := rows.Scan(&row.Id, &row.User, &row.Url, &row.Timestamp); err != nil {
			log.Fatal(err)
		}

		results = append(results, row)
	}

	for _, result := range results {
		log.Printf("Id = %d, User = %d, URL = %s, Timestamp = %s\n",
			result.Id, result.User, result.Url, result.Timestamp)
	}
}
