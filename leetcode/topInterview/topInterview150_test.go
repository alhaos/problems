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

func TestJump(t *testing.T) {
	data := []struct {
		mums     []int
		expected int
	}{
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
		},
	}

	for i, datum := range data {

		result := jump(datum.mums)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestHIndex(t *testing.T) {

	data := []struct {
		citations []int
		expected  int
	}{
		{
			[]int{3, 0, 6, 1, 5},
			3,
		},
		{
			[]int{1, 3, 1},
			1,
		},
	}

	for i, datum := range data {

		result := hIndex(datum.citations)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestProductExceptSelf(t *testing.T) {
	data := []struct {
		nums     []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}
	for i, datum := range data {

		result := productExceptSelf(datum.nums)

		if !reflect.DeepEqual(result, datum.expected) {
			t.Errorf("unexpected result for test index %d expected [%+v] got [%+v]", i, datum.expected, result)
		}
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	data := []struct {
		gas      []int
		cost     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}

	for i, datum := range data {
		result := canCompleteCircuit(datum.gas, datum.cost)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%d] got [%d]", i, datum.expected, result)
		}
	}
}

func TestCandy(t *testing.T) {
	data := []struct {
		ratings  []int
		expected int
	}{
		{
			[]int{1, 0, 2},
			5,
		},
		{
			[]int{1, 2, 2},
			4,
		},
		{
			[]int{1, 3, 2, 2, 1},
			7,
		},
		{
			[]int{1, 2, 87, 87, 87, 2, 1},
			13,
		}, {
			[]int{1, 3, 4, 5, 2},
			11,
		},
	}

	for i, datum := range data {
		result := candy(datum.ratings)

		if result != datum.expected {
			t.Errorf("unexpected result for test index %d expected [%d] got [%d]", i, datum.expected, result)
		}
	}
}

func TestTrap(t *testing.T) {

	data := []struct {
		height   []int
		expected int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for i, datum := range data {
		result := trap(datum.height)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%+v), expected [%d], but got [%d]",
				i,
				datum.height,
				datum.expected,
				result,
			)
		}
	}
}

func TestRomanToInt(t *testing.T) {

	data := []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, datum := range data {
		result := romanToInt(datum.s)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%strs), expected [%d], but got [%d]",
				i,
				datum.s,
				datum.expected,
				result,
			)
		}
	}
}

func TestIntToRoman(t *testing.T) {

	data := []struct {
		num      int
		expected string
	}{
		{3749, "MMMDCCXLIX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, datum := range data {
		result := intToRoman(datum.num)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%d), expected [%strs], but got [%strs]",
				i,
				datum.num,
				datum.expected,
				result,
			)
		}
	}
}

func TestLengthOfLastWord(t *testing.T) {
	data := []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for i, datum := range data {
		result := lengthOfLastWord(datum.s)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%strs), expected [%d], but got [%d]",
				i,
				datum.s,
				datum.expected,
				result,
			)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	data := []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for i, datum := range data {
		result := longestCommonPrefix(datum.strs)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%+v), expected [%s], but got [%s]",
				i,
				datum.strs,
				datum.expected,
				result,
			)
		}
	}
}

func TestReverseWords(t *testing.T) {
	data := []struct {
		s        string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"Привет мир こんにちは 世界 🚀", "🚀 世界 こんにちは мир Привет"},
	}

	for i, datum := range data {
		result := reverseWords(datum.s)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%s), expected [%s], but got [%s]",
				i,
				datum.s,
				datum.expected,
				result,
			)
		}
	}
}

func TestConvert(t *testing.T) {
	data := []struct {
		s        string
		rowNums  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
		{"AB", 1, "AB"},
	}

	for i, datum := range data {
		result := convert(datum.s, datum.rowNums)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (%s), expected [%s], but got [%s]",
				i,
				datum.s,
				datum.expected,
				result,
			)
		}
	}
}

func TestStrStr(t *testing.T) {
	data := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"abc", "c", 2},
		{"mississippi", "issip", 4},
	}

	for i, datum := range data {

		result := strStr(datum.haystack, datum.needle)

		if result != datum.expected {
			t.Errorf(
				"unexpected result for test data index %d (haystack:%s, needle:%s), expected [%d], but got [%d]",
				i,
				datum.haystack,
				datum.needle,
				datum.expected,
				result,
			)
		}
	}
}
