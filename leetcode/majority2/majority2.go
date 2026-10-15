package majority2

func majorityElement(nums []int) []int {

	var candidate1, candidate2 int
	var counter1, counter2 int

	for _, n := range nums {
		if n == candidate1 {
			counter1++
		} else if n == candidate2 {
			counter2++
		} else if counter1 == 0 {
			candidate1 = n
			counter1 = 1
		} else if counter2 == 0 {
			candidate2 = n
			counter2 = 1
		} else {
			counter1--
			counter2--
		}
	}

	counter1, counter2 = 0, 0

	for _, n := range nums {
		switch n {
		case candidate1:
			counter1++
		case candidate2:
			counter2++
		}
	}

	var result []int

	if counter1 > len(nums)/3 {
		result = append(result, candidate1)
	}

	if counter2 > len(nums)/3 {
		result = append(result, candidate2)
	}

	return result
}
