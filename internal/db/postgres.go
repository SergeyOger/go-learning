package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4/source/file" // 👈 REQUIRED for "file://" migrations

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
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
	runMigrations(db)
	return db, nil
}

func runMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migrated successfully!")
}
