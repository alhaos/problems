package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {

	m1 := make(map[int]int)

	for _, n := range arr {
		m1[n]++
	}

	seen := make(map[int]struct{}, len(m1))

	for _, v := range m1 {
		if _, exist := seen[v]; exist {
			return false
		}
		seen[v] = struct{}{}
	}

	return true
}
