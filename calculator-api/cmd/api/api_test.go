package api

import (
	"bytes"
	"mirzaadr/calculator-api/services/ratelimiter"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRateLimiterMiddleware(t *testing.T) {
	cfg := APIConfig{
		Ratelimiter: ratelimiter.Config{
			RequestsPerTimeFrame: 20,
			TimeFrame:            time.Second * 5,
			Enabled:              true,
		},
		Addr: ":8080",
	}

	app := newTestApplication(t, cfg)
	ts := httptest.NewServer(app.Mount())
	defer ts.Close()

	client := &http.Client{}
	mockIP := "192.168.1.1"
	marginOfError := 2

	for i := 0; i < cfg.Ratelimiter.RequestsPerTimeFrame+marginOfError; i++ {
		req, err := http.NewRequest("GET", ts.URL+"/health", nil)

		if err != nil {
			t.Fatalf("could not create request: %v", nil)
		}

		req.Header.Set("X-Forwarded-For", mockIP)

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}

		defer resp.Body.Close()

		if i < cfg.Ratelimiter.RequestsPerTimeFrame {
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", resp.Status)
			}
		} else {
			if resp.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status Too Many Request; got %v", resp.Status)
			}
		}
	}
}

func TestAdditionAPI(t *testing.T) {
	app := newTestApplication(t, APIConfig{
		Addr: ":8080",
	})
	mux := app.Mount()

	t.Run("should return bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/add", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		validData := []byte(`{"number1": 1, "number2": 20}`)
		req, err := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		checkResult(t, rr, 21)
	})
}

func TestSubtractionAPI(t *testing.T) {
	app := newTestApplication(t, APIConfig{
		Addr: ":8080",
	})
	mux := app.Mount()

	t.Run("should return bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/subtract", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		validData := []byte(`{"number1": 5, "number2": 20}`)
		req, err := http.NewRequest(http.MethodPost, "/subtract", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		checkResult(t, rr, -15)
	})
}

func TestMultiplicationAPI(t *testing.T) {
	app := newTestApplication(t, APIConfig{
		Addr: ":8080",
	})
	mux := app.Mount()

	t.Run("should return bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/multiply", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		validData := []byte(`{"number1": 5, "number2": 20}`)
		req, err := http.NewRequest(http.MethodPost, "/multiply", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		checkResult(t, rr, 100)
	})
}

func TestDivisionAPI(t *testing.T) {
	app := newTestApplication(t, APIConfig{
		Addr: ":8080",
	})
	mux := app.Mount()

	t.Run("should return bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/divide", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		validData := []byte(`{"dividend": 5, "divisor": 20}`)
		req, err := http.NewRequest(http.MethodPost, "/divide", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		checkResult(t, rr, 0.25)
	})
}

func TestSumAPI(t *testing.T) {
	app := newTestApplication(t, APIConfig{
		Addr: ":8080",
	})
	mux := app.Mount()

	t.Run("should return bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/sum", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return success", func(t *testing.T) {
		validData := []byte(`[1,2,3,4,5]`)
		req, err := http.NewRequest(http.MethodPost, "/sum", bytes.NewBuffer(validData))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			t.Fatal(err)
		}

		rr := executeRequest(req, mux)

		checkResponseCode(t, http.StatusOK, rr.Code)
		checkResult(t, rr, 15)
	})
}
