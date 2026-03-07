package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Check for common Railway and standard Postgres connection variables
	databaseURL := os.Getenv("DATABASE_URL")
	postgresURL := os.Getenv("POSTGRES_URL")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")
	pgUser := os.Getenv("PGUSER")
	pgPass := os.Getenv("PGPASSWORD")
	pgDB := os.Getenv("PGDATABASE")

	log.Println("--- Environment Variable Check ---")
	log.Printf("DATABASE_URL: %t (len=%d)", databaseURL != "", len(databaseURL))
	log.Printf("POSTGRES_URL: %t (len=%d)", postgresURL != "", len(postgresURL))
	log.Printf("PGHOST: %s", pgHost)
	log.Printf("PGPORT: %s", pgPort)
	log.Printf("PGUSER: %s", pgUser)
	log.Printf("PGDATABASE: %s", pgDB)
	log.Printf("PGPASSWORD set: %t", pgPass != "")
	log.Println("---------------------------------")

	var dsn string
	if databaseURL != "" {
		dsn = databaseURL
		log.Println("Using DATABASE_URL for connection")
	} else if postgresURL != "" {
		dsn = postgresURL
		log.Println("Using POSTGRES_URL for connection")
	} else if pgHost != "" && pgUser != "" && pgDB != "" {
		if pgPort == "" {
			pgPort = "5432"
		}
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", pgHost, pgUser, pgPass, pgDB, pgPort)
		log.Println("Using PG* variables for connection")
	} else {
		log.Println("WARNING: No remote database environment variables found. Falling back to localhost.")
		dsn = "host=localhost user=password= dbname= port=5432 sslmode=disable TimeZone=UTC"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v. DSN attempted: %s", err, dsn)
	}

	log.Println("Database connection successfully opened")
	DB = db
}
