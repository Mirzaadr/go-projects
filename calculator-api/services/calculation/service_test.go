package calculation

import (
	"math"
	"testing"
)

func TestAddition(t *testing.T) {
	calcService := CalculationService{}
	tests := []struct {
		name     string
		number1  float64
		number2  float64
		expected float64
	}{
		{
			name:     "positive Number",
			number1:  20,
			number2:  15,
			expected: 35,
		},
		{
			name:     "negative Number",
			number1:  -40,
			number2:  -15,
			expected: -55,
		},
		{
			name:     "zero Number",
			number1:  0,
			number2:  15,
			expected: 15,
		},
		{
			name:     "zero Number 2",
			number1:  25,
			number2:  0,
			expected: 25,
		},
		{
			name:     "positive with negative",
			number1:  -25,
			number2:  30,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calcService.Add(tt.number1, tt.number2)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.number1, tt.number2, got, tt.expected)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {
	calcService := CalculationService{}
	tests := []struct {
		name     string
		number1  float64
		number2  float64
		expected float64
	}{
		{
			name:     "positive Number",
			number1:  20,
			number2:  15,
			expected: 5,
		},
		{
			name:     "negative Number",
			number1:  -40,
			number2:  -15,
			expected: -25,
		},
		{
			name:     "zero Number",
			number1:  0,
			number2:  15,
			expected: -15,
		},
		{
			name:     "zero Number 2",
			number1:  25,
			number2:  0,
			expected: 25,
		},
		{
			name:     "positive with negative",
			number1:  -25,
			number2:  30,
			expected: -55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calcService.Substract(tt.number1, tt.number2)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.number1, tt.number2, got, tt.expected)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	calcService := CalculationService{}
	tests := []struct {
		name     string
		number1  float64
		number2  float64
		expected float64
	}{
		{
			name:     "positive Number",
			number1:  20,
			number2:  15,
			expected: 300,
		},
		{
			name:     "negative Number",
			number1:  -40,
			number2:  -15,
			expected: 600,
		},
		{
			name:     "zero Number",
			number1:  0,
			number2:  15,
			expected: 0,
		},
		{
			name:     "zero Number 2",
			number1:  25,
			number2:  0,
			expected: 0,
		},
		{
			name:     "positive with negative",
			number1:  -25,
			number2:  30,
			expected: -750,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calcService.Multiply(tt.number1, tt.number2)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.number1, tt.number2, got, tt.expected)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	calcService := CalculationService{}
	tests := []struct {
		name     string
		dividend float64
		divisor  float64
		expected float64
	}{
		{
			name:     "positive Number",
			dividend: 20,
			divisor:  5,
			expected: 4,
		},
		{
			name:     "negative Number",
			dividend: -40,
			divisor:  -10,
			expected: 4,
		},
		{
			name:     "zero Number",
			dividend: 0,
			divisor:  15,
			expected: 0,
		},
		{
			name:     "zero Number 2",
			dividend: 25,
			divisor:  0,
			expected: math.Inf(1),
		},
		{
			name:     "positive with negative",
			dividend: -25,
			divisor:  5,
			expected: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calcService.Divide(tt.dividend, tt.divisor)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.dividend, tt.divisor, got, tt.expected)
			}
		})
	}
}

func TestSumArray(t *testing.T) {
	calcService := CalculationService{}
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Mixed numbers",
			numbers:  []int{-3, 4, -1, 0, 2},
			expected: 2,
		},
		{
			name:     "Single number",
			numbers:  []int{10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := calcService.Sum(tt.numbers)
			if got != tt.expected {
				t.Errorf("sumArray(%v) = %v, want %v", tt.numbers, got, tt.expected)
			}
		})
	}
}
