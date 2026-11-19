package lengthoflongest

func lengthOfLongestSubstring(s string) int {
	l, r, length := 0, 0, 0
	m := make(map[byte]struct{})
	for r < len(s) {
		if _, exist := m[s[r]]; !exist {
			m[s[r]] = struct{}{}
			r++
			length = max(length, r-l)
		} else {
			delete(m, s[l])
			l++
		}
	}
	return length
}
