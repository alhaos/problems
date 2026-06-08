package validParentheses

func isValid(s string) bool {
	m := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	var opening []rune

	for _, c := range s {
		switch c {
		case '[', '{', '(':
			opening = append(opening, c)
		default:
			if len(opening) == 0 || opening[len(opening)-1] != m[c] {
				return false
			} else {
				opening = opening[:len(opening)-1]
			}
		}

	}
	return len(opening) == 0
}
