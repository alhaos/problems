package solvit

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	charIndexMap := make(map[byte]int)
	leftIndex := 0
	maxLength := 0
	rightIndex := 0

	for {

		if rightIndex == l {
			break
		}

		currentRightChar := s[rightIndex]

		idx, ok := charIndexMap[currentRightChar]

		if ok && idx >= leftIndex {
			leftIndex = idx + 1
		}

		charIndexMap[currentRightChar] = rightIndex

		if rightIndex-leftIndex+1 > maxLength {
			maxLength = rightIndex - leftIndex + 1
		}

		rightIndex++
	}

	return maxLength
}
