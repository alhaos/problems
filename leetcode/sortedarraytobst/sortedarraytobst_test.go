package sortedarraytobst

import (
	"slices"
	"testing"
)

// inorder собирает значения дерева в порядке left, root, right.
// У валидного BST результат должен совпасть с отсортированным входом.
func inorder(n *TreeNode, out *[]int) {
	if n == nil {
		return
	}
	inorder(n.Left, out)
	*out = append(*out, n.Val)
	inorder(n.Right, out)
}

// height возвращает высоту дерева или -2, если где-то нарушен баланс.
// Пустое дерево имеет высоту -1.
func height(n *TreeNode) int {
	if n == nil {
		return -1
	}
	l, r := height(n.Left), height(n.Right)
	if l == -2 || r == -2 || l-r > 1 || r-l > 1 {
		return -2
	}
	return max(l, r) + 1
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"example 1", []int{-10, -3, 0, 5, 9}},
		{"example 2", []int{1, 3}},
		{"single", []int{42}},
		{"three", []int{1, 2, 3}},
		{"even length", []int{1, 2, 3, 4, 5, 6}},
		{"seven", []int{1, 2, 3, 4, 5, 6, 7}},
		{"negatives", []int{-9, -7, -5, -3, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := sortedArrayToBST(tt.nums)

			var got []int
			inorder(root, &got)
			if !slices.Equal(got, tt.nums) {
				t.Errorf("inorder = %v, want %v (не BST или потеряны узлы)", got, tt.nums)
			}

			if height(root) == -2 {
				t.Errorf("дерево не сбалансировано для %v", tt.nums)
			}
		})
	}
}
