package statistics

import "testing"

func TestQuantileSorted(t *testing.T) {
	list := []float64{3, 6, 7, 8, 8, 9, 10, 13, 15, 16, 20}
	quantile, _ := QuantileSorted(list, 1)

	if quantile != 20 {
		t.Error("Quantile 100 must be 20(greater/last value)")
	}

	quantile, _ = QuantileSorted(list, 0)

	if quantile != 3 {
		t.Error("Quantile 0 must be 3(lower/first value)")
	}

	if quantile, _ = QuantileSorted(list, 0.5); quantile != 9 {
		t.Error("Quantile 50(median) must be equal to 9")
	}

	if _, err := QuantileSorted([]float64{}, 0.4); err == nil {
		t.Error("Empty list must be an error")
	}

	list = []float64{2, 3, 3, 4, 6, 6 , 6, 8, 9, 20}
	if quantile, _ = QuantileSorted(list, 0.4); quantile != 4 {
		t.Error("Quantile 0.4 must be 5.2")
	}
}