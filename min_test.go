package statistics

import "testing"

func TestMin(t *testing.T) {
	list := []float64{3, 5, 1, 4, 6, 2}

	if min, _ := Min(list); min != 1 {
		t.Error("The minimum value must be 1")
	}

	list = []float64{}
	
	if _, err := Min(list); err == nil {
		t.Error("Error must be thrown because it is an empty list")
	}
}