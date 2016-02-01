package statistics

// The Sum method calculate the float slice sum
// with a O(n)
func Sum(list []float64) float64 {
	total := 0.0
	for _, item := range list {
		total += item
	}

	return total
}