package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

const DB_DRIVER = "postgres"

func ConnectPostgres(host, port, password, user, database string) (*sql.DB, error) {
	fmt.Sprintf("Connecting to PostgreSQL at %s:%s...", host, port)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open(DB_DRIVER, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Connected to PostgreSQL!")
	return db, nil
}
