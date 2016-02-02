package statistics

import "errors"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Less than 0")
	}

	if n == 0 {
		return 1, nil
	}

	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}

	return total, nil
}