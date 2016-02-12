package statistics

import "testing"

func TestCumsum(t *testing.T) {
	cumsum, _ := Cumsum([]float64{1, 2, 3})
	if len(cumsum) != 3 {
		t.Error("Cumsum must have a length of 3")
	}

	if cumsum[0] != 1 {
		t.Error()
	}

	if cumsum[1] != 3 {
		t.Error()
	}

	if cumsum[2] != 6 {
		t.Error()
	}

	cumsum, err := Cumsum([]float64{})
	if err == nil {
		t.Error("Cumsum must return an error")
	}
}