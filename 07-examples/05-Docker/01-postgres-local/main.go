package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Define PostgreSQL connection parameters
	connStr := "user=postgres password=123456 dbname=postgres host=127.0.0.1 sslmode=disable"

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a table (if not exists)
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS public.data (
            id SERIAL PRIMARY KEY,
            name TEXT
        )
    `)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the table
	_, err = db.Exec("INSERT INTO public.data (name) VALUES ($1)", "Sample Data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data inserted successfully!")
}
