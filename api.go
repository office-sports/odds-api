package main

import (
	"github.com/gorilla/mux"
	"github.com/office-sports/odds-api/handlers"
)

func main() {
	router := mux.NewRouter()

	// secured routes and tokens
	router.HandleFunc("/token", handlers.GenerateToken).Methods("POST")
}
