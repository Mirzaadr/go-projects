package service

import (
	"fmt"
	"mirzaadr/calculator-api/middleware"
	"mirzaadr/calculator-api/types"
	"mirzaadr/calculator-api/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name string `json:"name"`
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /example", middleware.RecoveryMiddleware(h.handleExample))
	router.HandleFunc("POST /add", h.handleAdd)
	router.HandleFunc("POST /subtract", h.handleSubtract)
	router.HandleFunc("POST /multiply", h.handleMultiply)
	router.HandleFunc("POST /divide", middleware.RecoveryMiddleware(h.handleDivide))
	router.HandleFunc("POST /sum", h.handleSum)
}

func (h *Handler) handleExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}

func (h *Handler) handleAdd(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CalculationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// function
	result := payload.Number1 + payload.Number2

	utils.WriteJSON(w, http.StatusOK, types.ResponseData{
		Result: result,
	})
}
func (h *Handler) handleSubtract(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CalculationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// function
	result := payload.Number1 - payload.Number2

	utils.WriteJSON(w, http.StatusOK, types.ResponseData{
		Result: result,
	})
}
func (h *Handler) handleMultiply(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.CalculationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// function
	result := payload.Number1 * payload.Number2

	utils.WriteJSON(w, http.StatusOK, types.ResponseData{
		Result: result,
	})
}
func (h *Handler) handleDivide(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.DivisionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errors)
		return
	}

	// function
	result := payload.Dividend / payload.Divisor

	utils.WriteJSON(w, http.StatusOK, types.ResponseData{
		Result: result,
	})
}
func (h *Handler) handleSum(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload []float64
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	// if err := utils.Validate.Struct(payload); err != nil {
	// 	errors := err.(validator.ValidationErrors)
	// 	utils.WriteError(w, http.StatusBadRequest, errors)
	// 	return
	// }

	// function
	result := 0.0
	for i := 0; i < len(payload); i++ {
		result += payload[i]
	}

	utils.WriteJSON(w, http.StatusOK, types.ResponseData{
		Result: result,
	})
}

func sumArray(numbers []int) int {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result
}
