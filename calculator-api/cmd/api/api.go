package api

import (
	"log/slog"
	service "mirzaadr/calculator-api/services/calculation"
	"mirzaadr/calculator-api/services/ratelimiter"
	"net/http"
)

type APIServer struct {
	config             APIConfig
	calculationService service.ICalculationService
	ratelimiter        ratelimiter.Limiter
	logger             *slog.Logger
}

type APIConfig struct {
	Addr        string
	Ratelimiter ratelimiter.Config
	Version     string
	Env         string
}

func (app *APIServer) Mount() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /health", app.handleMiddleware(app.handleHealthCheck))
	router.Handle("POST /add", app.handleMiddleware(app.handleAdd))
	router.Handle("POST /subtract", app.handleMiddleware(app.handleSubtract))
	router.Handle("POST /multiply", app.handleMiddleware(app.handleMultiply))
	router.Handle("POST /divide", app.handleMiddleware(app.handleDivide))
	router.Handle("POST /sum", app.handleMiddleware(app.handleSum))

	return router
}

func (s *APIServer) Run() error {
	router := s.Mount()

	s.logger.Info("Starting application", "address", s.config.Addr)
	return http.ListenAndServe(s.config.Addr, router)
}

func NewAPIServer(cfg APIConfig, logger *slog.Logger, service service.ICalculationService, ratelimiter ratelimiter.Limiter) *APIServer {
	return &APIServer{
		config:             cfg,
		logger:             logger,
		calculationService: service,
		ratelimiter:        ratelimiter,
	}
}
