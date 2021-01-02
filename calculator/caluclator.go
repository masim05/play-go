package calculator

// Adder interface
type Adder interface {
	Add(a, b int) int
}

// Calculator type
type Calculator struct{}

// NewCalculator creates a new Calculator
func NewCalculator() Calculator {
	return Calculator{}
}

// Add two integer numbers
func (c Calculator) Add(a, b int) int {
	return a + b
}
