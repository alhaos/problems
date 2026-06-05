package twoPointers

func twoSum(numbers []int, target int) []int {

	leftIndex, rightIndex := 0, len(numbers)-1

	for {
		sum := numbers[leftIndex] + numbers[rightIndex]
		switch {
		case sum > target:
			rightIndex--
			continue
		case sum < target:
			leftIndex++
			continue
		default:
			return []int{leftIndex + 1, rightIndex + 1}
		}
	}
}
