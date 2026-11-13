package uniquenumberofoccurrences

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{"example 1", []int{1, 2, 2, 1, 1, 3}, true},
		{"example 2", []int{1, 2}, false},
		{"example 3", []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
		{"single element", []int{5}, true},
		{"all same", []int{4, 4, 4, 4}, true},
		{"all distinct", []int{1, 2, 3}, false},
		{"negative values collide", []int{-1, -1, 2, 2}, false},
		{"bounds", []int{-1000, 1000, 1000}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.arr); got != tt.want {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}
