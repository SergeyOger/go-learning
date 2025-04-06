package di

import (
	"database/sql"
	"user-management/internal/api"
	"user-management/internal/repository"
	"user-management/internal/service"

	"github.com/gorilla/mux"
)

func initUserAPi(router *mux.Router, database *sql.DB) {
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.CreateUser).Methods("DELETE")
}

func InitRoutes(router *mux.Router) {
	database := OpenDbConnection()
	initUserAPi(router, database)
}
