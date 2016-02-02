package statistics

import(
	"testing"
	"strconv"
)

func TestVariance(t *testing.T) {
	if variance, _ := Variance([]float64{1}); variance != 0 {
		t.Error("The variance of one number is zero")
	}

	if _, err := Variance([]float64{}); err == nil {
		t.Error("Variance must return an error")
	}

	list := []float64{1,2,3,4,5,6}
	variance, _ := Variance(list)
	if strconv.FormatFloat(variance, 'f', 3, 64) != "2.917" {
		t.Error("Variance must be equal to 2.917")
	}
}