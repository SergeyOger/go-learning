package di

import (
	"database/sql"
	"github.com/gorilla/mux"
	"user-management/internal/api"
	"user-management/internal/repository"
	"user-management/internal/service"
)

func initUserAPi(router *mux.Router, database *sql.DB) {
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
}

func InitRoutes(router *mux.Router) {
	database := OpenDbConnection()
	initUserAPi(router, database)
}
