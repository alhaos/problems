package TwoSum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},    // базовый
		{[]int{3, 2, 4}, 6, []int{1, 2}},         // не первые два
		{[]int{3, 3}, 6, []int{0, 1}},            // дубликаты
		{[]int{-1, -2, -3, -4}, -7, []int{2, 3}}, // отрицательные
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},      // нули
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v",
				tt.nums, tt.target, got, tt.want)
		}
	}
}
