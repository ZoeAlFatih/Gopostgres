package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetPostgresDB() (*sql.DB, error) {
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "qwords")
	os.Setenv("POSTGRES_DB_NAME", "go_postgres")

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB_NAME")

	desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)

	db, err := createConecction(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConecction(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
