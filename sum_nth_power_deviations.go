package statistics

import "math"

// The sum of deviations to the Nth power.
// Power n=2 it's the sum of squared deviations.
// Power n=3 it's the sum of cubed deviations.
func SumNthPowerDeviations(list []float64, power float64) (float64, error) {
	mean, err := Mean(list)
	if err != nil {
		return mean , err
	}
	sum := 0.0

	for _, val := range list {
		sum += math.Pow(val - mean, power)
	}

	return sum, nil
}