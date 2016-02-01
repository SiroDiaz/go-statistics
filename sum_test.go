package statistics

import "testing"

func TestSum(t *testing.T) {
	list := []float64{1.2, 2, 3}
	result := Sum(list)
	if result != 6.2 {
		t.Error("Sum result must be equal to 6.2")
	}
}