package leetcode

import (
	"reflect"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "Example 1: обычное дерево глубиной 3",
			root:     buildTree([]int{3, 9, 20, -1, -1, 15, 7}), // -1 означает null
			expected: 3,
		},
		{
			name:     "Example 2: дерево глубиной 2",
			root:     buildTree([]int{1, -1, 2}),
			expected: 2,
		},
		{
			name:     "Пустое дерево",
			root:     nil,
			expected: 0,
		},
		{
			name:     "Дерево из одного узла",
			root:     &TreeNode{Val: 1},
			expected: 1,
		},
		{
			name:     "Сбалансированное дерево глубиной 3",
			root:     buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: 3,
		},
		{
			name:     "Вырожденное дерево (только левая ветка)",
			root:     buildTree([]int{1, 2, -1, 3, -1, 4}),
			expected: 4,
		},
		{
			name:     "Вырожденное дерево (только правая ветка)",
			root:     buildTree([]int{1, -1, 2, -1, 3, -1, 4}),
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxDepth(tt.root)
			if got != tt.expected {
				t.Errorf("maxDepth() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Левый ребёнок
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Правый ребёнок
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestGenerate(t *testing.T) {
	testCases := []struct {
		desc     string
		numRows  int
		expected [][]int
	}{
		{
			desc:     "case_1",
			numRows:  1,
			expected: [][]int{{1}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := generate(tC.numRows)
			if !reflect.DeepEqual(result, tC.expected) {
				t.Errorf("Unexpected result for test case: %s, expected: %+v, but got: %+v", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestTitleToNumber(t *testing.T) {

	testCases := []struct {
		desc        string
		columnTitle string
		expected    int
	}{
		{
			desc:        "A",
			columnTitle: "A",
			expected:    1,
		},
		{
			desc:        "AB",
			columnTitle: "AB",
			expected:    28,
		},
		{
			desc:        "ZY",
			columnTitle: "ZY",
			expected:    701,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := titleToNumber(tC.columnTitle)
			if result != tC.expected {
				t.Errorf("Unexpected result from test %s, expected: %d, but got: %d", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		desc     string
		s        string
		t        string
		expected bool
	}{
		{
			desc:     "anagram/nagaram",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			desc:     "ret/car",
			s:        "ret",
			t:        "car",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := isAnagram(tC.s, tC.t)
			if result != tC.expected {
				t.Errorf("Unexpected result from test %s, expected: %t, but got: %t", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		desc     string
		a        string
		b        string
		expected string
	}{
		{
			desc:     "",
			a:        "11",
			b:        "1",
			expected: "100",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := addBinary(tC.a, tC.b)
			if result != tC.expected {
				t.Errorf("Unexpected result for test %s, expected: %s, but got: %s", tC.desc, tC.expected, result)
			}
		})
	}
}

func TestRandomOddOrEven(t *testing.T) {
	t.Run("random test", func(t *testing.T) {
		result := RandomOddOrEven()
		if !((result[0]&1 == 1 && result[1]&1 == 1 && result[2]&1 == 1) || (result[0]&1 == 0 && result[1]&1 == 0 && result[2]&1 == 0)) {
			t.Errorf("Unexpected result for test: %s\n, result is: %+v", "random test", result)
		}
	})
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		str      string
		expected bool
	}{
		{
			desc:     "#1",
			str:      "()",
			expected: true,
		},
		{
			desc:     "#2",
			str:      "()[]{}",
			expected: true,
		},
		{
			desc:     "#3",
			str:      "(]",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := isValid(tC.str)
			if result != tC.expected {
				t.Errorf("Unexpected result for test: %s, expected: %t, but got: %t", tC.desc, tC.expected, result)
			}
		})
	}
}
