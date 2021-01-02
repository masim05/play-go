package adder

// Adder interface
type Adder interface {
	Add(a, b int) int
}

// Add two integer numbers
func Add(a, b int) int {
	return a + b
}
