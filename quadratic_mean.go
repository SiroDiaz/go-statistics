package statistics

import "errors"
import "math"

// calculates the quadratic mean and throws an error
// if the list is empty.
func QuadraticMean(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("Empty list")
	}

	squaresSum := 0.0
	for _, item := range list {
		squaresSum += math.Pow(item, 2)
	}
	
	return math.Sqrt(squaresSum / float64(len(list))), nil
}