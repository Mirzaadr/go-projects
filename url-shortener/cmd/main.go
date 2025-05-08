package main

import (
	"errors"
	"log/slog"
	"mirzaadr/url-shortener/cmd/api"
	"mirzaadr/url-shortener/internal/db"
	"mirzaadr/url-shortener/internal/env"
	"mirzaadr/url-shortener/internal/store"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	// logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// loading .env file
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	// configuration
	cfg := api.Config{
		Port:   env.GetString("PORT", ":3000"),
		ApiURL: env.GetString("APP_URL", "localhost:3000"),
		Db: api.DBConfig{
			Addr:         env.GetString("DB_ADDR", "./storage.db"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		Env:     "development",
		Version: "1.0.0",
	}

	db, err := db.New(
		cfg.Db.Addr,
		cfg.Db.MaxOpenConns,
		cfg.Db.MaxIdleConns,
		cfg.Db.MaxIdleTime,
	)
	if err != nil {
		logger.Error("error establishing db connection", "error", err)
		os.Exit(1)
	}

	defer db.Close()
	logger.Info("database connection established")

	storage := store.NewStorage(db)

	tmpl := template.Must(template.New("").ParseGlob("./static/templates/*"))

	files := http.FileServer(http.Dir("./static"))

	app := api.NewApplication(
		cfg,
		tmpl,
		files,
		storage,
		logger,
	)

	srv := app.Mount()

	logger.Info("Starting website at port " + cfg.Port)

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("an error occured", "error", err)
		os.Exit(1)
	}
}
