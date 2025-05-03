package calculation

type ICalculationService interface {
	Add(number1, number2 float64) (float64, error)
	Substract(number1, number2 float64) (float64, error)
	Multiply(number1, number2 float64) (float64, error)
	Divide(dividend, divisor float64) (float64, error)
	Sum(numbers []float64) (float64, error)
}
