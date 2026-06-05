package twoPointers

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "базовый пример",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "ответ в конце массива",
			numbers: []int{1, 2, 3, 4, 5},
			target:  9,
			want:    []int{4, 5},
		},
		{
			name:    "ответ в начале и в конце",
			numbers: []int{1, 3, 5, 7, 9},
			target:  10,
			want:    []int{1, 5},
		},
		{
			name:    "два элемента — единственная пара",
			numbers: []int{3, 7},
			target:  10,
			want:    []int{1, 2},
		},
		{
			name:    "отрицательные числа",
			numbers: []int{-5, -3, -1, 2, 4},
			target:  -4,
			want:    []int{2, 3},
		},
		{
			name:    "отрицательное и положительное",
			numbers: []int{-3, 0, 1, 4, 7},
			target:  1,
			want:    []int{1, 4},
		},
		{
			name:    "оба числа одинаковые",
			numbers: []int{1, 3, 3, 7},
			target:  6,
			want:    []int{2, 3},
		},
		{
			name:    "нули в массиве",
			numbers: []int{0, 0, 3, 4},
			target:  0,
			want:    []int{1, 2},
		},
		{
			name:    "большой массив, ответ посередине",
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:  11,
			want:    []int{1, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum(%v, %d) = %v, want %v",
					tt.numbers, tt.target, got, tt.want)
			}
		})
	}
}
