package courses

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{"simple chain ok", 2, [][]int{{1, 0}}, true},
		{"direct cycle", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"longer chain ok", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{"no prerequisites", 3, nil, true},
		{"longer cycle", 4, [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}}, false},
		{"disconnected components ok", 5, [][]int{{1, 0}, {3, 2}}, true},
		{"disconnected with cycle", 5, [][]int{{1, 0}, {3, 2}, {2, 3}}, false},
		{"single course no deps", 1, nil, true},
		{"diamond shape ok", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"self-loop-like via two nodes", 2, [][]int{{0, 1}, {1, 0}}, false},
		{"multiple prereqs one course", 3, [][]int{{2, 0}, {2, 1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
