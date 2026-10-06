package safestates

import (
	"reflect"
	"testing"
)

func TestEventualSafeNodes(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  []int
	}{
		{"example 1", [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}, []int{2, 4, 5, 6}},
		{"example 2", [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}, []int{4}},
		{"no edges all safe", [][]int{{}, {}, {}}, []int{0, 1, 2}},
		{"self loop unsafe", [][]int{{0}}, []int{}},
		{"simple chain to terminal", [][]int{{1}, {2}, {}}, []int{0, 1, 2}},
		{"two separate cycles", [][]int{{1}, {0}, {3}, {2}}, []int{}},
		{"empty graph", [][]int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EventualSafeNodes(tt.graph)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
