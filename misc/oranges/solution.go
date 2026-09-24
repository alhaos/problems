package oranges

func OrangesRotting(grid [][]int) int {

	minutes := 0

	for {

		infected := false

		for i := range len(grid) {
			for j := range len(grid[i]) {
				if grid[i][j] == 2 {
					if i > 0 && grid[i-1][j] == 1 {
						grid[i-1][j] = 3
						infected = true
					}

					if i < len(grid)-1 && grid[i+1][j] == 1 {
						grid[i+1][j] = 3
						infected = true
					}

					if j > 0 && grid[i][j-1] == 1 {
						grid[i][j-1] = 3
						infected = true
					}

					if j < len(grid[i])-1 && grid[i][j+1] == 1 {
						grid[i][j+1] = 3
						infected = true
					}
				}
			}
		}

		for i := range len(grid) {
			for j := range len(grid[i]) {
				if grid[i][j] == 3 {
					grid[i][j] = 2
				}
			}
		}

		if !infected {
			break
		}

		minutes++

	}

	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
