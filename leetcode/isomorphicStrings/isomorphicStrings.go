package isomorphic

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)

	for i := range s {

		tValue, exist := m1[s[i]]
		if !exist {
			m1[s[i]] = t[i]
		} else {
			if tValue != t[i] {
				return false
			}
		}

		sValue, exist := m2[t[i]]
		if !exist {
			m2[s[i]] = s[i]
			continue
		}

		if sValue != s[i] {
			return false
		}

	}

	return true
}
