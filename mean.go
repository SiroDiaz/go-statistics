package statistics

import "errors"

// Compute the mean or average and throws an error instance
// for empty slices.
func Mean(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("Empty list")
	}

	return Sum(list) / float64(len(list)), nil
}