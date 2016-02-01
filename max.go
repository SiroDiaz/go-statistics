package statistics

import "errors"

// This function obtain the maximum number in a slice
// or returns and error if the list is empty.
func Max(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("List is empty")
	}

	maxValue := list[0]
	for i := 1 ; i < len(list); i++ {
		if maxValue < list[i] {
			maxValue = list[i]
		}
	}

	return maxValue, nil
}