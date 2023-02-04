package routes

import (
	"github.com/LucasRMP/golang-studies-rest-api/personality"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/personalities", personality.FindAllController).Methods("GET")
	router.HandleFunc("/personalities/{id}", personality.FindByIdController).Methods("GET")
	router.HandleFunc("/personalities", personality.CreateController).Methods("POST")
	router.HandleFunc("/personalities/{id}", personality.DeleteController).Methods("DELETE")
	router.HandleFunc("/personalities/{id}", personality.UpdateController).Methods("PUT")
}
