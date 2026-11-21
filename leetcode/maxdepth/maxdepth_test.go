package maxdepth

import (
	"math"
	"testing"
)

// null - маркер отсутствующего узла в списке уровней (как в LeetCode).
const null = math.MinInt

// build строит дерево из списка по уровням: [3,9,20,null,null,15,7].
func build(vals []int) *TreeNode {
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

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		tree []int
		want int
	}{
		{"example 1", []int{3, 9, 20, null, null, 15, 7}, 3},
		{"example 2", []int{1, null, 2}, 2},
		{"empty", []int{}, 0},
		{"single node", []int{1}, 1},
		{"left skew", []int{1, 2, null, 3, null, 4}, 4},
		{"right skew", []int{1, null, 2, null, 3, null, 4}, 4},
		{"balanced", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"deep left, shallow right", []int{1, 2, 3, 4, null, null, null, 5}, 4},
		{"deep right, shallow left", []int{1, 2, 3, null, null, null, 4}, 3},
		{"negative values", []int{-1, -2, -3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(build(tt.tree)); got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
