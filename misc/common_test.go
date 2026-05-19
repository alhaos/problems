package misc

import (
	"reflect"
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	testCases := []struct {
		desc     string
		a        []int
		b        []int
		expected []int
	}{
		{
			desc:     "common elements",
			a:        []int{23, 3, 1, 2},
			b:        []int{6, 2, 4, 23},
			expected: []int{2, 23},
		},
		{
			desc:     "no intersection",
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: []int{},
		},
		{
			desc:     "identical slices",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			desc:     "empty a",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: []int{},
		},
		{
			desc:     "empty b",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{},
		},
		{
			desc:     "both empty",
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
		{
			desc:     "single common element",
			a:        []int{1, 2, 3},
			b:        []int{3, 4, 5},
			expected: []int{3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := intersection(tC.a, tC.b)

			got := make([]int, len(result))
			copy(got, result)
			exp := make([]int, len(tC.expected))
			copy(exp, tC.expected)

			slices.Sort(got)
			slices.Sort(exp)

			if !reflect.DeepEqual(got, exp) {
				t.Errorf("intersection(%v, %v): expected %v, got %v", tC.a, tC.b, tC.expected, result)
			}
		})
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	testCases := []struct {
		desc     string
		graph    map[int][]int
		start    int
		expected []int
	}{
		{
			desc: "binary tree",
			graph: map[int][]int{
				0: {1, 2},
				1: {3, 4},
				2: {5, 6},
				3: {},
				4: {},
				5: {},
				6: {},
			},
			start:    0,
			expected: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			desc: "peterson graph",
			graph: map[int][]int{
				0: {1, 5},
				1: {2, 6},
				2: {3, 7},
				3: {4, 8},
				4: {0, 9},
				5: {7},
				7: {9},
				9: {6},
				6: {8},
				8: {5},
			},
			start:    0,
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			desc: "single node",
			graph: map[int][]int{
				0: {},
			},
			start:    0,
			expected: []int{0},
		},
		{
			desc: "linear chain",
			graph: map[int][]int{
				0: {1},
				1: {2},
				2: {3},
				3: {},
			},
			start:    0,
			expected: []int{0, 1, 2, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := breadthFirstSearch(tC.graph, tC.start)

			got := make([]int, len(result))
			copy(got, result)
			exp := make([]int, len(tC.expected))
			copy(exp, tC.expected)

			slices.Sort(got)
			slices.Sort(exp)

			if !reflect.DeepEqual(got, exp) {
				t.Errorf("breadthFirstSearch (%v, start=%d): expected %v, got %v", tC.graph, tC.start, tC.expected, result)
			}
		})
	}
}
