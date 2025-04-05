package di

import (
	"database/sql"
	"log"
	"user-management/config"
	"user-management/internal/db"
)

func OpenDbConnection() *sql.DB {

	host := config.GetEnv("DB_HOST", "localhost")
	port := config.GetEnv("DB_PORT", "5432")
	password := config.GetEnv("DB_PASSWORD", "postgres")
	user := config.GetEnv("DB_USER", "postgres")
	database := config.GetEnv("DB_NAME", "user-management")

	// Connect to the database
	db, err := db.ConnectPostgres(host, port, password, user, database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Note: The caller is responsible for closing the database connection
	return db
}
