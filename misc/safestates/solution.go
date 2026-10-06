package safestates

func EventualSafeNodes(graph [][]int) []int {

	graphMap := make(map[int][]int)

	for i, dep := range graph {
		graphMap[i] = dep
	}

	return nil
}
