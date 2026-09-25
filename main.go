package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Private Knowledge Vault started")
	http.HandleFunc("/health", healthCheckHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}


func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Status : OK"))
		log.Println(r.Method)
		log.Println(r.URL.Path)
	}