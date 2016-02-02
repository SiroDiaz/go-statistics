package statistics

import "errors"

func Variance(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("Empty list")
	}

	sumNthPowerDeviation, err := SumNthPowerDeviations(list, 2)
	return  sumNthPowerDeviation / float64(len(list)), err
}