package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request entered | PATH :%v | Method :%v | address: %v\n", r.URL.Path, r.Method, r.RemoteAddr)
		next.ServeHTTP(w, r)
		log.Printf("Request Completed | PATH :%v | Method :%v | in %v\n", r.URL.Path, r.Method, time.Since(start))
	})
}

func SecurityHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'")

		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}
