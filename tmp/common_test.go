package leetCode

import (
	"reflect"
	"testing"
)

func Test_maxProfit(t *testing.T) {

	data := []struct {
		ints     []int
		expected int
	}{
		{
			ints:     []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
	}

	for _, datum := range data {

		result := maxProfit(datum.ints)
		if result != datum.expected {
			t.Errorf("unexpected result for input: %v expected: %d, but got %d", datum.ints, datum.expected, result)
		}
	}
}

func Test_mergeAlternately(t *testing.T) {

	data := []struct {
		world1   string
		world2   string
		expected string
	}{
		{
			world1:   "abc",
			world2:   "pqr",
			expected: "apbqcr",
		},
	}

	for _, datum := range data {

		result := mergeAlternately(datum.world1, datum.world1)
		if result != datum.expected {
			t.Errorf("unexpected result for input: %s, %s expected: %s, but got %s", datum.world1, datum.world2, datum.expected, result)
		}
	}
}

func Test_kidsWithCandies(t *testing.T) {

	data := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expected:     []bool{true, true, true, false, true},
		},
	}

	for _, datum := range data {

		result := kidsWithCandies(datum.candies, datum.extraCandies)
		if !reflect.DeepEqual(result, datum.expected) {
			t.Errorf("unepected result for %v and %d, expected %v, but got %v", datum.candies, datum.extraCandies, datum.expected, result)
		}
	}
}

func Test_canPlaceFlowers(t *testing.T) {

	data := []struct {
		flowerbed []int
		n         int
		expected  bool
	}{
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
			n:         2,
			expected:  true,
		},
	}

	for i, datum := range data {

		result := canPlaceFlowers(datum.flowerbed, datum.n)
		t.Logf("test #%d pass for %v and %v, got %v", i, datum.flowerbed, datum.n, result)
		if result != datum.expected {
			t.Errorf("unepected result for %v and %v, expected %v, but got %v", datum.flowerbed, datum.n, datum.expected, result)
		}
	}
}

func Test_lengthOfLastWord(t *testing.T) {
	data := []struct {
		input    string
		expected int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
		{"a", 1},
	}

	for _, datum := range data {
		result := lengthOfLastWord(datum.input)
		if result != datum.expected {
			t.Errorf("unexpected relust for [%s] expected [%d], but got [%d]", datum.input, datum.expected, result)
		}
	}
}
