package oranges

type point struct {
	i, j int
}

func OrangesRottingBFS(grid [][]int) int {

	var points []point
	var freshCount int

	// collect start data
	for i := range len(grid) {
		for j := range len(grid[i]) {

			if grid[i][j] == 2 {
				points = append(points, point{i, j})
			}

			if grid[i][j] == 1 {
				freshCount++
			}

		}
	}

	// Process points
	minutes := 0
	for len(points) > 0 && freshCount > 0 {

		var next []point

		for _, p := range points {

			if p.i > 0 && grid[p.i-1][p.j] == 1 {
				grid[p.i-1][p.j] = 2
				next = append(next, point{p.i - 1, p.j})
				freshCount--
			}

			if p.i < len(grid)-1 && grid[p.i+1][p.j] == 1 {
				grid[p.i+1][p.j] = 2
				next = append(next, point{p.i + 1, p.j})
				freshCount--
			}

			if p.j > 0 && grid[p.i][p.j-1] == 1 {
				grid[p.i][p.j-1] = 2
				next = append(next, point{p.i, p.j - 1})
				freshCount--
			}

			if p.j < len(grid[p.i])-1 && grid[p.i][p.j+1] == 1 {
				grid[p.i][p.j+1] = 2
				next = append(next, point{p.i, p.j + 1})
				freshCount--
			}
		}

		points = next
		minutes++

	}

	if freshCount > 0 {
		return -1
	}
	return minutes
}
