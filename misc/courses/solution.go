package courses

func CanFinish(numCourses int, prerequisites [][]int) bool {

	graph := make(map[int][]int)

	for _, p := range prerequisites {
		graph[p[0]] = append(graph[p[0]], p[1])
	}

	state := make(map[int]int)

	for course := range numCourses {
		if hasCycle(course, graph, state) {
			return false
		}
	}

	return true

}

func hasCycle(course int, graph map[int][]int, state map[int]int) bool {
	if state[course] == 1 {
		return true // встретили курс, который уже в процессе — цикл
	}
	if state[course] == 2 {
		return false // уже проверяли, циклов нет
	}

	state[course] = 1 // помечаем "в процессе"

	for _, dep := range graph[course] {
		if hasCycle(dep, graph, state) {
			return true
		}
	}

	state[course] = 2 // помечаем "завершён"
	return false
}
