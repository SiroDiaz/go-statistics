package statistics

import(
	"testing"
)

func TestSumNthPowerDeviations(t *testing.T) {
	if tspd, _ := SumNthPowerDeviations([]float64{0,0,0}, 2); tspd != 0 {
		t.Error("Sum nth power deviation must be equal to 0")
	}

	if tspd, _ := SumNthPowerDeviations([]float64{0,1,2}, 2); tspd != 2 {
		t.Error("Sum nth power deviation must be equal to 2")
	}

	if _, err := SumNthPowerDeviations([]float64{}, 2); err == nil {
		t.Error("Empty list")
	}
}