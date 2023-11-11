package main

import (
	"github.com/gorilla/mux"
	"github.com/office-sports/odds-api/handlers"
	"github.com/office-sports/odds-api/middlewares"
	"github.com/office-sports/odds-api/models"
	"os"
)

func main() {
	var configFileParam string

	if len(os.Args) > 1 {
		configFileParam = os.Args[1]
	}

	config, err := models.GetConfig(configFileParam)
	if err != nil {
		panic(err)
	}
	models.InitDB(config.GetMysqlConnectionString())

	router := mux.NewRouter()

	// secured routes and tokens
	router.HandleFunc("/token", handlers.GenerateToken).Methods("POST")
	router.Handle("/auth", middlewares.AuthMiddleware(handlers.CheckAuth()))
}
