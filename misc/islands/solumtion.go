package islands

func NumIslands(grid [][]byte) int {
	NumIslands := 0

	for i := range len(grid) {
		for j := range len(grid[i]) {
			if grid[i][j] == '1' {
				NumIslands++
				sink(grid, i, j)
			}
		}
	}

	return NumIslands
}

func sink(grid [][]byte, i, j int) {

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	// Up
	if i > 0 {
		sink(grid, i-1, j)
	}

	// Down

	if i < len(grid)-1 {
		sink(grid, i+1, j)
	}

	// Left
	if j > 0 {
		sink(grid, i, j-1)
	}

	// Right
	if j < len(grid[i])-1 {
		sink(grid, i, j+1)
	}
}
