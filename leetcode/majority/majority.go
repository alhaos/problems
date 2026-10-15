package majority

func majorityElement(nums []int) int {

	candidate := nums[0]
	counter := 1

	for _, newCandidate := range nums[1:] {

		if newCandidate == candidate {
			counter++
			continue
		}

		counter--

		if counter == 0 {
			candidate = newCandidate
			counter++
		}

	}

	return candidate
}
