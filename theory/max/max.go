package max

import "fmt"

func sliceMaxMin(slice []int) (int, int, error) {

	if len(slice) == 0 {
		return 0, 0, fmt.Errorf("null slice provided")
	}

	maxI := slice[0]
	minI := slice[0]
	for i, v := range slice[1:] {
		if v > maxI {
			maxI = i
		}
		if v < minI {
			minI = i
		}
	}

	return maxI, minI, nil
}
