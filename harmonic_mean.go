package statistics

import "errors"

func HarmonicMean(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("Empty list")
	}

	cumSum := 0.0
	for i := 0; i < len(list); i++ {
		if list[i] <= 0 {
			return 0.0, errors.New("Harmonic mean is only for positive numbers")
		}
		
	}
}