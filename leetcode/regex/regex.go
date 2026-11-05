package regex

func isMatch(s, p string) bool {

	if len(s) != len(p) {
		return false
	}

	for i := range min(len(s), len(p)) {
		if s[i] != p[i] {
			return false
		}
	}

	return true
}
