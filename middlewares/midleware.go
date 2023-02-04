package middlewares

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func GetCorsHandler(router *mux.Router) http.Handler {
	corsConfig := handlers.AllowedOrigins([]string{"*"})
	corsHandler := handlers.CORS(corsConfig)
	return corsHandler(router)
}
