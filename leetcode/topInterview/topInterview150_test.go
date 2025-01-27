package topInterview

import (
	"reflect"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	data := []struct {
		num1     []int
		m        int
		num2     []int
		n        int
		expected []int
	}{
		{
			num1:     []int{1, 2, 3, 0, 0, 0},
			m:        3,
			num2:     []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			num1:     []int{1},
			m:        1,
			num2:     []int{},
			n:        0,
			expected: []int{1},
		},
		{
			num1:     []int{4, 5, 6, 0, 0, 0},
			m:        3,
			num2:     []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i, datum := range data {

		merge(datum.num1, datum.m, datum.num2, datum.n)

		if !reflect.DeepEqual(datum.num1, datum.expected) {
			t.Errorf("unexpected result for test index %d expected %+v 'got %+v", i, datum.expected, datum.num1)
		}
	}
}
