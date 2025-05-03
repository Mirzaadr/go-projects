package calculation

type CalculationService struct{}

func (s *CalculationService) Add(number1, number2 float64) (float64, error) {
	return number1 + number2, nil
}

func (s *CalculationService) Substract(number1, number2 float64) (float64, error) {
	return number1 - number2, nil
}

func (s *CalculationService) Multiply(number1, number2 float64) (float64, error) {
	return number1 * number2, nil
}

func (s *CalculationService) Divide(dividend, divisor float64) (float64, error) {
	return dividend / divisor, nil
}

func (s *CalculationService) Sum(numbers []int) (int, error) {
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return result, nil
}
