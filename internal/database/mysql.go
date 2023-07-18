package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewMySQLDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	return db
}

// func RunMigrations(db *sql.DB) error {
// 	migrationPath := "file://migrations"
// 	m, err := migrate.New(migrationPath, os.Getenv("DB_CONNECTION_STRING"))
// 	if err != nil {
// 		return fmt.Errorf("failed to initialize migrations: %v", err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		return fmt.Errorf("failed to run migrations: %v", err)
// 	}

// 	return nil
// }
