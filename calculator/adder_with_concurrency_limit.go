package calculator

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using concurrencylimit template

//go:generate gowrap gen -p github.com/masim05/play-go/calculator -i Adder -t concurrencylimit -o adder_with_concurrency_limit.go

// AdderWithConcurrencyLimit implements Adder
type AdderWithConcurrencyLimit struct {
	_base  Adder
	_burst chan int
}

// NewAdderWithConcurrencyLimit instruments an implementation of the Adder with rate limiting
func NewAdderWithConcurrencyLimit(base Adder, concurrentCalls int) *AdderWithConcurrencyLimit {
	d := &AdderWithConcurrencyLimit{
		_base:  base,
		_burst: make(chan int, concurrentCalls),
	}

	return d
}

// Add implements Adder
func (_d *AdderWithConcurrencyLimit) Add(a int, b int) (i1 int) {

	_d._burst <- 1

	defer func() {
		<-_d._burst
	}()

	return _d._base.Add(a, b)
}
