package api

import (
	"fmt"
	"mirzaadr/calculator-api/utils"
	"net/http"
	"runtime/debug"
)

func (app *APIServer) handleMiddleware(next http.HandlerFunc) http.Handler {
	return utils.ChainMiddleware(
		next,
		app.realIP,
		app.recoveryMiddleware,
		app.loggerMiddleware,
		app.rateLimiterMiddleware,
	)
}

// a middleware to log information about request
func (app *APIServer) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("Receive Request", "method", r.Method, "path", r.URL.Path, "address", r.RemoteAddr)
		next.ServeHTTP(w, r)
	},
	)
}

// a middleware to handle unexpected panic
func (app *APIServer) recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer a function to catch any panics
		defer func() {
			if err := recover(); err != nil {
				// Log the panic and stack trace
				app.logger.Error(fmt.Sprintf("Caught panic: %v", err), "stack", string(debug.Stack()))
				app.internalServerError(w, r, fmt.Errorf("%v", err))
			}
		}()
		// Call next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// a middleware that sets a http.Request's RemoteAddr to the results
// of parsing either the True-Client-IP, X-Real-IP or the X-Forwarded-For headers (in that order).
func (app *APIServer) realIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rip := utils.RealIP(r); rip != "" {
			r.RemoteAddr = rip
		}
		next.ServeHTTP(w, r)
	})
}

// a middleware to limit request from certain address based on our configuration
func (app *APIServer) rateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if app.config.Ratelimiter.Enabled {
			if allow, retryAfter := app.ratelimiter.Allow(r.RemoteAddr); !allow {
				app.rateLimitExceededResponse(w, r, retryAfter.String())
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
