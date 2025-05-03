package api

import (
	"mirzaadr/calculator-api/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (app *APIServer) parseJSON(r *http.Request, payload any) error {
	// get JSON payload
	return utils.ParseJSON(r, payload)
}

func (app *APIServer) parseAndValidateJSON(r *http.Request, payload any) error {
	// get JSON payload
	if err := utils.ParseJSON(r, payload); err != nil {
		return err
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		return errors
	}

	return nil
}

func (app *APIServer) writeError(w http.ResponseWriter, status int, err error) {
	type envelope struct {
		Error string `json:"error"`
	}
	utils.WriteJSON(w, status, &envelope{Error: err.Error()})
}

func (app *APIServer) writeResponseJSON(w http.ResponseWriter, status int, result any) error {
	type envelope struct {
		Result any `json:"result"`
	}
	return utils.WriteJSON(w, status, &envelope{Result: result})
}
