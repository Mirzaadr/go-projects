package api

import (
	"mirzaadr/url-shortener/types"
	"net/http"
	"net/url"
	"strconv"
)

func redirectToError(w http.ResponseWriter, r *http.Request, errData types.ErrorData) {
	params := url.Values{}
	params.Set("status", strconv.Itoa(errData.Status))
	params.Set("msg", errData.Message)
	params.Set("desc", errData.Description)

	http.Redirect(w, r, "/error?"+params.Encode(), http.StatusSeeOther)
}

func (app *Application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("internal error", "method", r.Method, "path", r.URL.Path, "error", err)
	redirectToError(w, r, types.ErrorData{
		Status:      http.StatusInternalServerError,
		Message:     "Internal Server Error",
		Description: "The server encountered a problem.",
	})
}

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warn("bad request", "method", r.Method, "path", r.URL.Path, "error", err)
	redirectToError(w, r, types.ErrorData{
		Status:      http.StatusBadRequest,
		Message:     "Invalid Input",
		Description: err.Error(),
	})
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warn("not found error", "method", r.Method, "path", r.URL.Path, "error", err)
	redirectToError(w, r, types.ErrorData{
		Status:      http.StatusNotFound,
		Message:     "Page Not Found",
		Description: err.Error(),
	})
}
