package statistics

import "errors"

func Min(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("List is empty")
	}

	minValue := list[0]
	for i := 1; i < len(list); i++ {
		if minValue > list[i] {
			minValue = list[i]
		}
	}

	return minValue, nil
}