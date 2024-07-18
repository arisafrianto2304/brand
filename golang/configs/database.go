package configs

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	host := "db"
	port := "5432"
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	log.Printf("Connecting to database with the following details:\n")
	log.Printf("Host: %s\n", host)
	log.Printf("Port: %s\n", port)
	log.Printf("User: %s\n", user)
	log.Printf("Password: %s\n", password) // Hati-hati dalam menampilkan password di log pada produksi
	log.Printf("Database Name: %s\n", dbname)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			break
		}
		log.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")
	return db
}
