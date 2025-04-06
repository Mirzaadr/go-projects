package types

type CalculationPayload struct {
	Number1 float64 `json:"number1" validate:"required"`
	Number2 float64 `json:"number2"`
}

type DivisionPayload struct {
	Dividend float64 `json:"dividend"`
	Divisor  float64 `json:"divisor"`
}

type ResponseData struct {
	Result float64 `json:"result"`
}
