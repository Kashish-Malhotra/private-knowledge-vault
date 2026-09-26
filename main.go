package main

import (
	"log"
	"net/http"
	"private-knowledge-vault/routes"
)

func main() {
	log.Println("Private Knowledge Vault started")
	routes.SetupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
