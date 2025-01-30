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
			t.Errorf("unexpected expected for test index %d expected %+v 'got %+v", i, datum.expected, datum.num1)
		}
	}
}

func TestRemoveElement(t *testing.T) {

	data := []struct {
		nums     []int
		val      int
		expected struct {
			result    int
			numsAfter []int
		}
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			struct {
				result    int
				numsAfter []int
			}{
				2,
				[]int{2, 2},
			},
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			struct {
				result    int
				numsAfter []int
			}{
				5,
				[]int{0, 1, 3, 0, 4},
			},
		},
	}

	for i, datum := range data {
		result := removeElement(datum.nums, datum.val)

		if datum.expected.result != result || !reflect.DeepEqual(datum.nums[:len(datum.expected.numsAfter)], datum.expected.numsAfter) {
			t.Errorf("unexpected expected for test index %d expected %+v got [%+v, %+v ]", i, datum.expected, result, datum.nums[:len(datum.expected.numsAfter)])
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {

	data := []struct {
		nums      []int
		expected  int
		numsAfter []int
	}{
		{
			nums:      []int{1, 1, 2},
			expected:  2,
			numsAfter: []int{1, 2},
		},
		{
			nums:      []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:  5,
			numsAfter: []int{0, 1, 2, 3, 4},
		},
		{
			nums:      []int{},
			expected:  0,
			numsAfter: []int{},
		},
		{
			nums:      []int{1},
			expected:  1,
			numsAfter: []int{1},
		},
	}

	for i, datum := range data {
		result := removeDuplicates(datum.nums)

		if result != datum.expected || !reflect.DeepEqual(datum.nums[:datum.expected], datum.numsAfter) {
			t.Errorf("unexpected result for test index %d expected [%+v, %+v] got [%+v, %+v]", i, datum.expected, datum.numsAfter, result, datum.numsAfter[:datum.expected])
		}
	}
}

func TestRemoveDuplicatesII(t *testing.T) {

	data := []struct {
		nums      []int
		expected  int
		numsAfter []int
	}{
		{
			nums:      []int{},
			expected:  0,
			numsAfter: []int{},
		},
		{
			nums:      []int{1},
			expected:  1,
			numsAfter: []int{1},
		},
		{
			nums:      []int{-9, -9, -2, -2, -2, 0, 0, 0, 1},
			expected:  7,
			numsAfter: []int{-9, -9, -2, -2, 0, 0, 1},
		},
		{
			nums:      []int{1, 1, 1, 2, 2, 3},
			expected:  5,
			numsAfter: []int{1, 1, 2, 2, 3},
		},
		{
			nums:      []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected:  7,
			numsAfter: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}

	for i, datum := range data {
		result := removeDuplicatesII(datum.nums)
		if result != datum.expected || !reflect.DeepEqual(datum.nums[:datum.expected], datum.numsAfter) {
			t.Errorf("unexpected result for test index %d expected [%+v, %+v] got [%+v, %+v]", i, datum.expected, datum.numsAfter, result, datum.nums[:datum.expected])
		}
	}

}

func TestMajorityElement(t *testing.T) {
	data := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1},
			1,
		},

		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	for i, datum := range data {

		result := majorityElement(datum.nums)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestRotate(t *testing.T) {

	data := []struct {
		nums        []int
		k           int
		numsMutated []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			1,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			2,
			[]int{6, 7, 1, 2, 3, 4, 5},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			[]int{-1, -100, 3, 99},
			1,
			[]int{99, -1, -100, 3},
		},
		{
			[]int{-1, -100, 3, 99},
			0,
			[]int{-1, -100, 3, 99},
		},
		{
			[]int{0},
			99,
			[]int{0},
		},
	}

	for i, datum := range data {

		rotate(datum.nums, datum.k)

		if !reflect.DeepEqual(datum.nums, datum.numsMutated) {
			t.Errorf("unexpected result for test index %d key was: %d expected [%+v] got [%+v]", i, datum.k, datum.numsMutated, datum.nums)
		}
	}
}

func TestMaxProfit(t *testing.T) {

	data := []struct {
		prices   []int
		expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{0},
			0,
		},
	}

	for i, datum := range data {
		result := maxProfit(datum.prices)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestMaxProfitII(t *testing.T) {

	data := []struct {
		prices   []int
		expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for i, datum := range data {

		result := maxProfitII(datum.prices)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestCanJump(t *testing.T) {

	data := []struct {
		mums     []int
		expected bool
	}{
		{
			[]int{0},
			true,
		},
		{
			[]int{0, 0},
			false,
		},
		{
			[]int{},
			false,
		},
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
	}

	for i, datum := range data {

		result := canJump(datum.mums)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}
