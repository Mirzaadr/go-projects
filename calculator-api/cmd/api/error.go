package api

import (
	"fmt"
	"net/http"
)

func (app *APIServer) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("internal error", "method", r.Method, "path", r.URL.Path, "error", err)
	app.writeError(w, http.StatusInternalServerError, fmt.Errorf("the server encountered a problem"))
}

func (app *APIServer) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warn("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	app.writeError(w, http.StatusBadRequest, err)
}

func (app *APIServer) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warn("not found error", "method", r.Method, "path", r.URL.Path, "error", err)
	app.writeError(w, http.StatusNotFound, fmt.Errorf("not found"))
}

func (app *APIServer) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warn("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err)
	app.writeError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}

func (app *APIServer) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.logger.Warn("rate limit exceeded", "method", r.Method, "path", r.URL.Path)
	w.Header().Set("Retry-After", retryAfter)
	app.writeError(w, http.StatusTooManyRequests, fmt.Errorf("rate limit exceeded, retry after: %s", retryAfter))
}
