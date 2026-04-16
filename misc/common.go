package misc

func intersection(a, b []int) []int {

	m := make(map[int]struct{})
	var result []int

	for _, el := range a {
		m[el] = struct{}{}
	}

	for _, el := range b {
		if _, exist := m[el]; exist {
			result = append(result, el)
		}
	}

	return result
}
