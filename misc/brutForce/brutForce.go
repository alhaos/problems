package brutForce

import "math"

type current []int

func (c current) sum() int {
	var sum int
	for _, v := range c {
		sum += v
	}
	return sum
}

// FindGraterThan finds all elements grater than x
func FindGraterThan(nums []int, x int) []int {
	var result []int
	for _, n := range nums {
		if n > x {
			result = append(result, x)
		}
	}
	return result
}

// TwoSum returns all value pairs that sum to target
func TwoSum(nums []int, target int) [][2]int {

	var result [][2]int

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			r := nums[i] * nums[j]
			if r == target {
				result = append(result, [2]int{nums[i], nums[j]})
			}
		}
	}
	return result
}

func maxPairSum(nums []int) int {
	var maxValue = math.MinInt
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			maxValue = max(maxValue, nums[i]+nums[j])
		}
	}
	return maxValue
}

func maxTripleSum(nums []int) int {
	var maxValue = math.MinInt
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				maxValue = max(maxValue, nums[i]+nums[j]+nums[k])
			}
		}
	}
	return maxValue
}

// Найти подмножество с максимальной суммой
func MaxSubsetSum(nums []int) int {
	l := len(nums)
	var maxSum int
	for mask := range 1 << l {
		var sum int
		for i := range nums {
			if mask&(1<<i) != 0 {
				sum += nums[i]
			}
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func CombinationSum(nums []int, target int) [][]int {

	var result [][]int
	var current current
	var backtrack func(start int)
	backtrack = func(start int) {

		s := current.sum()

		if s == target {
			tmp := make([]int, len(current))
			copy(tmp, current)
			result = append(result, tmp)
			return
		}

		if s > target {
			return
		}

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i]) // шаг вперёд
			backtrack(i + 1)                   // рекурсия
			current = current[:len(current)-1] // откат
		}
	}

	backtrack(0)

	return result
}

// Найти все перестановки и вернуть только те,
// где первый элемент меньше последнего
func permutationsAscEdge(nums []int) [][]int {

	result := [][]int{}
	var backtrack func(start int)

	backtrack = func(start int) {

		if start == len(nums) {

			tmp := make([]int, len(nums))
			copy(tmp, mums)
			result = append(result, tmp)
			return
		}

		for i := start; i < len(nums); i++ {

			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}

		backtrack(0)


	}

	backtrack(0)

}

// permutationsAscEdge([]int{1, 2, 3}) →
// [1,2,3], [1,3,2], [2,3,1]  -- первый < последний
// [2,1,3], [3,1,2], [3,2,1]  -- отбрасываем
