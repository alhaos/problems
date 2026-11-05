package palindrome

func longestPalindrome(s string) string {
	var left, right int
	var result string
	for i := range s {
		left, right = expand(s, i, i)
		if right-left+1 > len(result) {
			result = s[left : right+1]
		}
	}
	for i := range s {
		left, right = expand(s, i, i+1)
		if right-left+1 > len(result) {
			result = s[left : right+1]
		}
	}
	return result
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	left++
	right--
	return left, right
}
