package middlewares

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Entering route: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) 
	})
}
