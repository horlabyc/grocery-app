package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/horlabyc/grocery-app/internal/handlers"
	"github.com/horlabyc/grocery-app/internal/middleware"
	"github.com/horlabyc/grocery-app/internal/services"
	"github.com/horlabyc/grocery-app/internal/storage/postgres"
	"github.com/horlabyc/grocery-app/internal/utils"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// initialize db config
	dbConfig := postgres.Config{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnvAsInt("DB_PORT", 5432),
		User:     utils.GetEnv("DB_USER", "admin"),
		Password: utils.GetEnv("DB_PASSWORD", "password"),
		DBName:   utils.GetEnv("DB_NAME", "grocery_app"),
		SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
	}

	db, err := postgres.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()
	fmt.Println("Connected to database")

	// Initialize Router
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	//Apply Middleware
	router.Use(middleware.LoggerMiddleware)

	//Apply CORS Middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Initialize Repositories
	shopRepo := postgres.NewShopRepo(db)

	// Services
	shopService := services.NewShopService(shopRepo)

	// Handlers
	shopHandler := handlers.NewShopHandler(shopService)

	//Shop Routes
	router.HandleFunc("/shops", shopHandler.CreateShop).Methods(http.MethodPost)
	router.HandleFunc("/shops", shopHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/shops/{id:[0-9]+}", shopHandler.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/shops/{id:[0-9]+}", shopHandler.Update).Methods(http.MethodPut)

	serverPort := utils.GetEnv("SERVER_PORT", "8080")

	// Create Server
	srv := &http.Server{
		Addr:         "localhost:" + serverPort,
		Handler:      c.Handler(router),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start Server in a goroutine
	log.Println("Starting server on port", serverPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // blocks until signal is received

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exited properly")
}
