package main

import (
	"log/slog"
	"mirzaadr/calculator-api/cmd/api"
	service "mirzaadr/calculator-api/services/calculation"
	"mirzaadr/calculator-api/services/ratelimiter"
	"os"
	"time"
)

func main() {
	cfg := api.APIConfig{
		Addr:    ":4000",
		Version: "1.0.0",
		Env:     "development",
		Ratelimiter: ratelimiter.Config{
			RequestsPerTimeFrame: 20,
			TimeFrame:            time.Second * 5,
			Enabled:              false,
		},
	}

	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	calculationService := service.CalculationService{}
	// Rate limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.Ratelimiter.RequestsPerTimeFrame,
		cfg.Ratelimiter.TimeFrame,
	)

	server := api.NewAPIServer(cfg, jsonLogger, &calculationService, rateLimiter)
	if err := server.Run(); err != nil {
		jsonLogger.Error(err.Error())
	}
}
