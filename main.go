package main

import (
	"log"
	"net/http"

	"github.com/LucasRMP/golang-studies-rest-api/database"
	"github.com/LucasRMP/golang-studies-rest-api/middlewares"
	"github.com/LucasRMP/golang-studies-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.Connect()

	startServer()
}

func startServer() {
	router := mux.NewRouter()
	router.Use(middlewares.ContentTypeMiddleware)

	routes.RegisterRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", middlewares.GetCorsHandler(router)))
}
