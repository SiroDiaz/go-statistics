package statistics

import(
	"testing"
	"strconv"
)

func TestGeometricMean(t *testing.T) {
	if gm, _ := GeometricMean([]float64{2, 8}); gm != 4 {
		t.Error("The geometric mean must be equal to 4")
	}
	
	if gm, _ := GeometricMean([]float64{4, 1, 1.0 / 32}); strconv.FormatFloat(gm, 'g', 1, 64) != "0.5" {
		t.Error("The geometric mean must be equal to 0.5")
	}

	if _, err := GeometricMean([]float64{}); err == nil {
		t.Error("Empty list for geometric mean should return an error")
	}
}