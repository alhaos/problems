package symmetrictree

import (
	"math"
	"testing"
)

// null обозначает пустое место в level-order массиве.
// Значения в задаче от -100 до 100, так что коллизий нет.
const null = math.MinInt

// buildTree строит дерево из level-order массива в формате LeetCode.
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}

	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"example 1", []int{1, 2, 2, 3, 4, 4, 3}, true},
		{"example 2", []int{1, 2, 2, null, 3, null, 3}, false},
		{"single node", []int{1}, true},
		{"only left child", []int{1, 2}, false},
		{"different children values", []int{1, 2, 3}, false},
		{"same shape inside, mirrored", []int{1, 2, 2, null, 3, 3}, true},
		{"same values but wrong shape", []int{1, 2, 2, 2, null, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSymmetric(buildTree(tt.vals))
			if got != tt.want {
				t.Errorf("isSymmetric(%v) = %v, want %v", tt.vals, got, tt.want)
			}
		})
	}
}
