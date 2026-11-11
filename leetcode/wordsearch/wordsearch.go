package wordsearch

func exist(board [][]byte, word string) bool {

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {

		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
			return false
		}

		if board[r][c] != word[i] {
			return false
		}

		buffer := board[r][c]
		board[r][c] = 0

		found := dfs(r-1, c, i+1) || dfs(r+1, c, i+1) || dfs(r, c-1, i+1) || dfs(r, c+1, i+1)

		board[r][c] = buffer

		return found

	}

	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
