package statistics

import(
	"errors"
	"math"
)

// Calculates the quantile 
func QuantileSorted(list []float64, p float64) (float64, error) {
	index := float64(len(list)) * p
	if len(list) == 0 {
		return 0.0, errors.New("List empty")
	}

	if p < 0 || p > 1 {
		return 0.0, errors.New("Quantile must be between 0 and 1")
	} else if p == 1 {
		return list[len(list) - 1], nil
	} else if p == 0 {
		return list[0], nil
	} else if math.Remainder(p, 1) != 0 {
		return list[int(math.Ceil(index)) - 1], nil
	} else if len(list) % 2 == 0 {
		return (list[int(index) - 1] + list[int(index)]) / 2, nil
	} else {
		return list[int(index)], nil
	}
}