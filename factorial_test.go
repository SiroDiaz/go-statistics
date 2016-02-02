package statistics

import "testing"

func TestFactorial(t *testing.T) {
	if fact, _ := Factorial(3); fact != 6 {
		t.Error("The factorial must be 6")
	}
	
	if fact, _ := Factorial(0); fact != 1 {
		t.Error("By default the 0 factorial is 1")
	}

	if _, err := Factorial(-1); err == nil {
		t.Error("There is not factorial for negative numbers")
	}
}