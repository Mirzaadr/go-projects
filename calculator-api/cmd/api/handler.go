package api

import (
	"mirzaadr/calculator-api/types"
	"net/http"
)

func (app *APIServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.Env,
		"version": app.config.Version,
	}

	if err := app.writeResponseJSON(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *APIServer) handleAdd(w http.ResponseWriter, r *http.Request) {
	// parse body and validate it
	var payload types.CalculationPayload
	if err := app.parseAndValidateJSON(r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// function
	result, _ := app.calculationService.Add(payload.Number1, payload.Number2)

	if err := app.writeResponseJSON(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *APIServer) handleSubtract(w http.ResponseWriter, r *http.Request) {
	// parse body and validate it
	var payload types.CalculationPayload
	if err := app.parseAndValidateJSON(r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// function
	result, _ := app.calculationService.Substract(payload.Number1, payload.Number2)

	if err := app.writeResponseJSON(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *APIServer) handleMultiply(w http.ResponseWriter, r *http.Request) {
	// parse body and validate it
	var payload types.CalculationPayload
	if err := app.parseAndValidateJSON(r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// function
	result, _ := app.calculationService.Multiply(payload.Number1, payload.Number2)

	if err := app.writeResponseJSON(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *APIServer) handleDivide(w http.ResponseWriter, r *http.Request) {
	// parse body and validate it
	var payload types.DivisionPayload
	if err := app.parseAndValidateJSON(r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// function
	result, _ := app.calculationService.Divide(payload.Dividend, payload.Divisor)

	if err := app.writeResponseJSON(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *APIServer) handleSum(w http.ResponseWriter, r *http.Request) {
	// parse body and validate it
	var payload []float64
	if err := app.parseJSON(r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// function
	result, _ := app.calculationService.Sum(payload)

	if err := app.writeResponseJSON(w, http.StatusOK, result); err != nil {
		app.internalServerError(w, r, err)
	}
}
