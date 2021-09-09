package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			log.Println(cookie.Value)
			next.ServeHTTP(w, r)
		}
	})
}
