package api

import (
	"fmt"
	"mirzaadr/url-shortener/types"
	"net/http"
	"net/url"
)

func (app *Application) handlerRoot(w http.ResponseWriter, r *http.Request) {
	if err := app.tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Failed to render the page", http.StatusInternalServerError)
	}
}

func (app *Application) handlerShorten(w http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("url")
	if longUrl == "" {
		app.badRequestResponse(w, r, fmt.Errorf("no url is provided"))
		return
	}

	// Validate URL format
	parsedURL, err := url.ParseRequestURI(longUrl)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		app.badRequestResponse(w, r, fmt.Errorf("url is not valid"))
		return
	}

	slug, err := app.store.Urls.Create(longUrl)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	err = app.tmpl.ExecuteTemplate(w, "result.html", types.ResultData{
		Link: app.cfg.ApiURL + "/" + slug,
	})
	if err != nil {
		http.Error(w, "Failed to render the page", http.StatusInternalServerError)
	}
}

func (app *Application) handlerRedirectURL(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	original_url, err := app.store.Urls.GetBySlug(slug)
	if err != nil {
		// TODO: separate response for not found and internal error
		app.logger.Error("an error occured", "error", err)
		app.notFoundResponse(w, r, fmt.Errorf("sorry, we couldn’t find the page you’re looking for"))
		return
	}
	http.Redirect(w, r, original_url, http.StatusFound)
}

func (app *Application) handlerError(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("status")
	msg := r.URL.Query().Get("msg")
	desc := r.URL.Query().Get("desc")

	data := types.ErrorData{
		Status:      500,
		Message:     "Unknown error",
		Description: "",
	}

	if code != "" {
		fmt.Sscanf(code, "%d", &data.Status)
	}
	if msg != "" {
		data.Message = msg
	}
	if desc != "" {
		data.Description = desc
	}

	if err := app.tmpl.ExecuteTemplate(w, "error.html", data); err != nil {
		http.Error(w, "Failed to render error page", http.StatusInternalServerError)
	}
}
