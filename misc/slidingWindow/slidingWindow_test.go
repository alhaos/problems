package slidingWindow

import "testing"

func TestMaxSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "базовый пример",
			nums: []int{2, 1, 5, 1, 3, 2},
			k:    3,
			want: 9,
		},
		{
			name: "максимум в начале массива",
			nums: []int{9, 8, 1, 2, 3},
			k:    2,
			want: 17,
		},
		{
			name: "максимум в конце массива",
			nums: []int{1, 2, 3, 8, 9},
			k:    2,
			want: 17,
		},
		{
			name: "все элементы одинаковые",
			nums: []int{5, 5, 5, 5, 5},
			k:    3,
			want: 15,
		},
		{
			name: "k равно длине массива",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 15,
		},
		{
			name: "k равно 1",
			nums: []int{3, 1, 4, 1, 5, 9},
			k:    1,
			want: 9,
		},
		{
			name: "отрицательные числа",
			nums: []int{-3, -1, -4, -1, -5},
			k:    2,
			want: -4, // -3 + (-1) = -4
		},
		{
			name: "смесь отрицательных и положительных",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			k:    3,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSum(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("maxSum(%v, %d) = %d, want %d",
					tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
