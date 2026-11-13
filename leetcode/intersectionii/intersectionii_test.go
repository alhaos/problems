package intersectii

import (
	"slices"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"example 1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"example 2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{"no intersection", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"identical", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"more in first", []int{1, 1, 1, 2}, []int{1, 1}, []int{1, 1}},
		{"more in second", []int{1}, []int{1, 1, 1}, []int{1}},
		{"min of counts", []int{1, 1, 1, 2, 2}, []int{1, 1, 2, 2, 2}, []int{1, 1, 2, 2}},
		{"zeros", []int{0, 0}, []int{0}, []int{0}},
		{"single miss", []int{0}, []int{1}, []int{}},
		{"max values", []int{1000, 1000, 0}, []int{1000, 1000, 1000}, []int{1000, 1000}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// копии, чтобы тест ловил мутацию входных данных не молча
			n1 := slices.Clone(tt.nums1)
			n2 := slices.Clone(tt.nums2)

			got := intersect(n1, n2)
			slices.Sort(got)
			slices.Sort(tt.want)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
