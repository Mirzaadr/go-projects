package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

// a middleware to handle unexpected panic
func RecoveryMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Defer a functino to catch any panics
		defer func() {
			if err := recover(); err != nil {
				// Log the panic and stack trace
				msg := "Caught panic: %v, Stack trace: %s"
				log.Printf(msg, err, string(debug.Stack()))

				// return a generic error message to the client
				er := http.StatusInternalServerError
				http.Error(w, "Internal Server Error", er)
			}
		}()
		// Call next handler in the chain
		next(w, r)
	}
}
