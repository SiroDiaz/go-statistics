package statistics

import(
	"testing"
	"strconv"
)

func TestQuadraticMean(t *testing.T) {
	if mean, _ := QuadraticMean([]float64{-1, 1, -1, 1}); mean != 1 {
		t.Error("Quadratic mean should return 1")
	}

	if mean, _ := QuadraticMean([]float64{3, 4, 5}); strconv.FormatFloat(mean, 'f', 3, 64) != "4.082" {
		t.Error("Quadratic mean should return 4.082")
	}


	if _, err := QuadraticMean([]float64{}); err == nil {
		t.Error("Quadratic mean should recieve a slice whit length higher than 0")
	}
}