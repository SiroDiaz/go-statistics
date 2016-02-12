package statistics

import "errors"

func Cumsum(list []float64) ([]float64, error) {
	if len(list) == 0 {
		return []float64{0}, errors.New("Empty list given")
	}

	total := 0.0
	cumsum := make([]float64, len(list))

	for index, val := range list {
		total += val
		cumsum[index] = total
	}

	return cumsum, nil
}