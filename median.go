package statistics

import(
	"sort"
	"errors"
)

func Median(list []float64) (float64, error) {
	if len(list) == 0 {
		return 0.0, errors.New("")
	}

	if !sort.Float64sAreSorted(list) {
		sort.Float64s(list)
	}

	if len(list) % 2 == 1 {
		return list[(len(list) / 2)], nil
	}
	
	return (list[len(list) / 2] + list[len(list) / 2 - 1]) / 2, nil
}