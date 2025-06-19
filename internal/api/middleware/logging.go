package middleware

import (
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware logs the details of each incoming request and the response time.
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Log the incoming request
        log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)

        // Call the next handler
        next.ServeHTTP(w, r)

        // Log the response time
        duration := time.Since(start)
        log.Printf("Completed request: %s %s in %v", r.Method, r.URL.Path, duration)
    })
}