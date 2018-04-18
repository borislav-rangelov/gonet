package handlers

import (
	"log"
	"net/http"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: [%s] %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
	})
}
