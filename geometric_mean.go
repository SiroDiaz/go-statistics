package statistics

import(
	"math"
	"errors"
)

// Obtains the geometric mean of a given slice.
// Returns an error if the list is empty.
func GeometricMean(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("Empty list")
	}

	cumMulti := 1.0
	for _, val := range list {
		if val == 0 {
			return 0.0, errors.New("0 multiplication")
		}
		cumMulti *= val
	}

	return math.Pow(cumMulti, 1 / float64(len(list))), nil
}