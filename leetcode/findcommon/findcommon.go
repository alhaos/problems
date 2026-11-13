package findcommon

func commonChars(words []string) []string {
	var counts [26]int
	for _, r := range words[0] {
		counts[r-'a']++
	}

	for _, word := range words[1:] {
		var cur [26]int
		for _, r := range word {
			cur[r-'a']++
		}
		for i := range counts {
			counts[i] = min(counts[i], cur[i])
		}
	}

	var result []string
	for i, c := range counts {
		for range c {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}
