package routes

import (
	"log"
	"net/http"
	"private-knowledge-vault/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/health", firstMiddleware(secondMiddleware(http.HandlerFunc(handlers.HealthCheckHandler))))
	mux.Handle("/hello", firstMiddleware(secondMiddleware(http.HandlerFunc(handlers.Hello))))
	mux.Handle("/ws", http.HandlerFunc(handlers.WebSocketHandler))
	return mux
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("In Middleware - Before handler") // Pre-processing logic here (e.g., logging, authentication)
		next.ServeHTTP(w, r)
		log.Print("In Middleware - After handler")// Post-processing logic here (e.g., logging, cleanup)
	})
}

func secondMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("In Middleware - Before handler") // Pre-processing logic here (e.g., logging, authentication)
		next.ServeHTTP(w, r)
		log.Print("In Middleware - After handler") // Post-processing logic here (e.g., logging, cleanup)
	})

}

