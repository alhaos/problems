package triangle

func MinimumTotal(triangle [][]int) int {
	for rowIndex := len(triangle) - 2; rowIndex >= 0; rowIndex-- {
		for valueIndex := range triangle[rowIndex] {
			triangle[rowIndex][valueIndex] += min(
				triangle[rowIndex+1][valueIndex],
				triangle[rowIndex+1][valueIndex+1],
			)
		}
	}
	return triangle[0][0]
}
