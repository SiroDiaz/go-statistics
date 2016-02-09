package statistics

import "testing"

func TestMedian(t *testing.T) {
	median, _ := Median([]float64{1,2,3})
	if median != 2 {
		t.Error("Median must be 2")
	}

	median, _ = Median([]float64{1,2})
	if median != 1.5 {
		t.Error("Median must be 1 / 2")
	}

	median, _ = Median([]float64{1,2,3,4})
	if median != 2.5 {
		t.Error("Median must be equal to 2.5")
	}

	median, err := Median([]float64{})
	if err == nil {
		t.Error("Empty list. Must return an error")
	}
}