package binarySearch

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "базовый пример — элемент найден",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "элемент не найден",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "элемент в начале массива",
			nums:   []int{1, 3, 5, 7, 9},
			target: 1,
			want:   0,
		},
		{
			name:   "элемент в конце массива",
			nums:   []int{1, 3, 5, 7, 9},
			target: 9,
			want:   4,
		},
		{
			name:   "массив из одного элемента — найден",
			nums:   []int{42},
			target: 42,
			want:   0,
		},
		{
			name:   "массив из одного элемента — не найден",
			nums:   []int{42},
			target: 7,
			want:   -1,
		},
		{
			name:   "чётное количество элементов",
			nums:   []int{1, 2, 3, 4},
			target: 3,
			want:   2,
		},
		{
			name:   "отрицательные числа",
			nums:   []int{-10, -5, -3, -1, 0},
			target: -3,
			want:   2,
		},
		{
			name:   "целевой элемент — минимальный в массиве",
			nums:   []int{2, 4, 6, 8, 10},
			target: 2,
			want:   0,
		},
		{
			name:   "целевой элемент — максимальный в массиве",
			nums:   []int{2, 4, 6, 8, 10},
			target: 10,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d",
					tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
