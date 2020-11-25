package adder_test

import "testing"
import "github.com/masim05/play-go/adder"

func TestAdderHappyPath(t *testing.T) {
	res := adder.Add(1, 2)
	if res != 3 {
		t.Error("Expected 3, got ", res)
	}
}
