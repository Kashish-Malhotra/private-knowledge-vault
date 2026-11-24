package main

import (
	"log"
	"net/http"
	"private-knowledge-vault/routes"
)

func main() {
	log.Println("Private Knowledge Vault started")
	mux := routes.SetupRoutes()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
