package api

import (
	"log/slog"
	"mirzaadr/url-shortener/internal/store"
	"net/http"
	"text/template"
)

type Application struct {
	cfg    Config
	tmpl   *template.Template
	static http.Handler
	store  store.Storage
	logger *slog.Logger
}

type Config struct {
	Port    string
	ApiURL  string
	Env     string
	Version string
	Db      DBConfig
}

type DBConfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

func NewApplication(
	config Config,
	tmpl *template.Template,
	static http.Handler,
	store store.Storage,
	logger *slog.Logger) *Application {
	return &Application{
		cfg:    config,
		tmpl:   tmpl,
		static: static,
		store:  store,
		logger: logger,
	}
}

func (a *Application) Mount() *http.Server {
	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", a.static))

	router.HandleFunc("GET /{$}", a.handlerRoot)
	router.HandleFunc("POST /shorten", a.handlerShorten)
	router.HandleFunc("GET /error", a.handlerError)
	router.HandleFunc("GET /{slug}", a.handlerRedirectURL)

	srv := http.Server{
		Addr:    a.cfg.Port,
		Handler: router,
	}

	return &srv
}
