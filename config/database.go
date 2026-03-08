package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	databaseURL := os.Getenv("DATABASE_URL")
	postgresURL := os.Getenv("POSTGRES_URL")
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")
	pgUser := os.Getenv("PGUSER")
	pgPass := os.Getenv("PGPASSWORD")
	pgDB := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("--- Database Environment Variable Check ---")
	fmt.Printf("DATABASE_URL: %t (len=%d)\n", databaseURL != "", len(databaseURL))
	fmt.Printf("POSTGRES_URL: %t (len=%d)\n", postgresURL != "", len(postgresURL))
	fmt.Printf("PGHOST: %s, PGPORT: %s, PGUSER: %s, PGDATABASE: %s\n", pgHost, pgPort, pgUser, pgDB)
	fmt.Printf("DB_HOST: %s, DB_PORT: %s, DB_USER: %s, DB_NAME: %s\n", dbHost, dbPort, dbUser, dbName)
	fmt.Println("------------------------------------------")

	var dsn string
	
	// Determine core connection parameters
	host := pgHost
	if host == "" { host = dbHost }
	if host == "" { host = "localhost" }

	port := pgPort
	if port == "" { port = dbPort }

	user := pgUser
	if user == "" { user = dbUser }

	pass := pgPass
	if pass == "" { pass = dbPass }

	dbname := pgDB
	if dbname == "" { dbname = dbName }

	// If we have an explicit dbname, we MUST use it to avoid falling into 'railway' default DB
	if dbname != "" && (databaseURL != "" || postgresURL != "") {
		url := databaseURL
		if url == "" { url = postgresURL }
		
		log.Printf("Forcing database name '%s' on provided URL", dbname)
		
		// If it's a standard postgres DSN URL, we try to replace the path
		if strings.HasPrefix(url, "postgres://") || strings.HasPrefix(url, "postgresql://") {
			// Find the last slash before query params
			base := url
			query := ""
			if idx := strings.Index(url, "?"); idx != -1 {
				base = url[:idx]
				query = url[idx:]
			}
			
			lastSlash := strings.LastIndex(base, "/")
			if lastSlash != -1 {
				dsn = base[:lastSlash+1] + dbname + query
			} else {
				dsn = url // Fallback if format is weird
			}
		} else {
			dsn = url // Not a URL format we readily recognize for path replacement
		}
	} else if databaseURL != "" {
		dsn = databaseURL
		log.Println("Using DATABASE_URL as is")
	} else if postgresURL != "" {
		dsn = postgresURL
		log.Println("Using POSTGRES_URL as is")
	} else {
		// Construct from components
		if port == "" { port = "5432" }
		if dbname == "" { dbname = "postgres" }
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, pass, dbname, port)
		log.Printf("Using constructed DSN: host=%s user=%s dbname=%s port=%s", host, user, dbname, port)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v. DSN attempted: %s", err, dsn)
	}

	log.Println("Database connection successfully opened")
	DB = db
}
