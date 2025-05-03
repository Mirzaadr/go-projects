package api

import (
	"encoding/json"
	"log/slog"
	service "mirzaadr/calculator-api/services/calculation"
	"mirzaadr/calculator-api/services/ratelimiter"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func newTestApplication(t *testing.T, cfg APIConfig) *APIServer {
	t.Helper()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	calculationService := service.CalculationService{}

	// Rate Limiter
	rateLimiter := ratelimiter.NewFixedWindowLimiter(
		cfg.Ratelimiter.RequestsPerTimeFrame,
		cfg.Ratelimiter.TimeFrame,
	)

	return &APIServer{
		config:             cfg,
		ratelimiter:        rateLimiter,
		logger:             logger,
		calculationService: &calculationService,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}

func checkResult(t *testing.T, rr *httptest.ResponseRecorder, expected float64) {
	var response map[string]float64
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response["result"] != expected {
		t.Errorf("Expected result to be %v, got %v", expected, response["result"])
	}
}
