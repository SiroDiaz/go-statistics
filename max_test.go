package statistics

import "testing"

func TestMax(t *testing.T) {
	list := []float64{3, 5, 1, 4, 6, 2}

	if max, _ := Max(list); max != 6 {
		t.Error("The maximum value must be 6")
	}

	list = []float64{}
	
	if _, err := Max(list); err == nil {
		t.Error("Error must be thrown because it is an empty list")
	}
}