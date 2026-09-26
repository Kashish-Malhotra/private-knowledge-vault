package routes

import (
	"net/http"
	"private-knowledge-vault/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/hello", handlers.Hello)
}