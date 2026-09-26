package handlers

import (
	"net/http"
)	

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + " " + r.URL.Path + " is OK"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + " " + r.URL.Path + " is OK\n"))
	w.Write([]byte("Hello From Private Knowledge Vault"))
}
