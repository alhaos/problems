package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "example 2",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "single element",
			nums: []int{7},
			want: 7,
		},
		{
			name: "two elements same",
			nums: []int{5, 5},
			want: 5,
		},
		{
			name: "majority at the start",
			nums: []int{1, 1, 1, 2, 3},
			want: 1,
		},
		{
			name: "majority at the end",
			nums: []int{4, 5, 6, 6, 6, 6},
			want: 6,
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, 2, 2},
			want: -1,
		},
		{
			name: "large majority block",
			nums: []int{9, 1, 9, 9, 2, 9, 9, 3, 9},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MajorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("MajorityElement(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
