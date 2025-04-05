package main

import (
	"log"
	"net/http"
	"user-management/di"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	di.InitRoutes(router)
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
