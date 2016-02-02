package statistics

import "testing"

func TestHarmonicMean(t *testing.T) {
	if hm, _ := HarmonicMean([]float64{1, 1}); hm != 1 {
		t.Error("The harmonic mean must be equal to 1")
	}
	
	if hm, err := HarmonicMean([]float64{0}); hm != 0.0 || err == nil {
		t.Error(err)
	}

	if _, err := HarmonicMean([]float64{}); err == nil {
		t.Error(err)
	}
}