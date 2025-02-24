package bootstrap

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/horlabyc/grocery-app/internal/middleware"
	"github.com/rs/cors"
)

func InitializeRouter() http.Handler {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	// Apply Middleware
	router.Use(middleware.LoggerMiddleware)

	// Apply CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	return c.Handler(router)
}
