package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		in   [][]int
		ni   []int
		want [][]int
	}{
		{"my", [][]int{{5, 7}, {8, 9}}, []int{1, 2}, [][]int{{1, 2}, {5, 7}, {8, 9}}},
		{"basic", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{"multiple merge", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{"empty input", [][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{"before all", [][]int{{3, 5}}, []int{1, 2}, [][]int{{1, 2}, {3, 5}}},
		{"after all", [][]int{{1, 2}}, []int{3, 4}, [][]int{{1, 2}, {3, 4}}},
		{"between, no overlap", [][]int{{1, 2}, {6, 7}}, []int{3, 4}, [][]int{{1, 2}, {3, 4}, {6, 7}}},
		{"touch right end", [][]int{{1, 5}}, []int{5, 7}, [][]int{{1, 7}}},
		{"touch left end", [][]int{{3, 5}}, []int{1, 3}, [][]int{{1, 5}}},
		{"inside existing", [][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{"same interval", [][]int{{1, 5}}, []int{1, 5}, [][]int{{1, 5}}},
		{"covers all", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{0, 10}, [][]int{{0, 10}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insert(tt.in, tt.ni)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert(%v, %v) = %v, want %v", tt.in, tt.ni, got, tt.want)
			}
		})
	}
}
