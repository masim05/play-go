package calculator_test

import (
	"testing"

	"github.com/masim05/play-go/calculator"
)

func TestAddHappyPath(t *testing.T) {
	c := calculator.NewCalculator()
	res := c.Add(1, 2)
	if res != 3 {
		t.Error("Expected 3, got ", res)
	}
}

func TestAddBigNumbers(t *testing.T) {
	c := calculator.NewCalculator()
	res := c.Add(1000000, 20000000000000000)
	if res != 1000000+20000000000000000 {
		t.Error("Got ", res)
	}
}
