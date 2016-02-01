package statistics

import "testing"
import "strconv"

func TestMean(t *testing.T) {
	list := []float64{1, 2, 3}
	if mean, _ := Mean(list); mean != 2 {
		t.Error("The mean must be 2 no: "+
			strconv.FormatFloat(mean, 'f', 2, 64),
		)
	}

	if mean, err := Mean([]float64{}); mean != 0.0 && err != nil {
		t.Error(err)
	} 
}