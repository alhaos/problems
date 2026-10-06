package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		want [][]int
	}{
		{"basic", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"touching", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"single", [][]int{{1, 2}}, [][]int{{1, 2}}},
		{"unsorted", [][]int{{8, 10}, {1, 3}, {15, 18}, {2, 6}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"unsorted reversed pair", [][]int{{1, 4}, {0, 4}}, [][]int{{0, 4}}},
		{"nested", [][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		{"all merge into one", [][]int{{1, 10}, {2, 3}, {4, 5}, {6, 7}}, [][]int{{1, 10}}},
		{"no overlap", [][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{"chain", [][]int{{1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 4}}},
		{"same intervals", [][]int{{2, 5}, {2, 5}, {2, 5}}, [][]int{{2, 5}}},
		{"point intervals", [][]int{{3, 3}, {1, 1}, {3, 3}}, [][]int{{1, 1}, {3, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
