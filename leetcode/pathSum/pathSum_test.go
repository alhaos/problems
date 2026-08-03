package pathSum

import "testing"

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name       string
		root       *TreeNode
		targetSum  int
		want       bool
	}{
		{
			name:      "nil root",
			root:      nil,
			targetSum: 0,
			want:      false,
		},
		{
			name:      "single node, sum matches",
			root:      &TreeNode{Val: 5},
			targetSum: 5,
			want:      true,
		},
		{
			name:      "single node, sum does not match",
			root:      &TreeNode{Val: 5},
			targetSum: 1,
			want:      false,
		},
		{
			name: "classic leetcode example - true",
			// 5 -> 4 -> 11 -> (7,2), 5 -> 8 -> (13, 4->1)
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:  11,
						Left: &TreeNode{Val: 7},
						Right: &TreeNode{Val: 2},
					},
				},
				Right: &TreeNode{
					Val:  8,
					Left: &TreeNode{Val: 13},
					Right: &TreeNode{
						Val:   4,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22, // 5+4+11+2
			want:      true,
		},
		{
			name: "classic leetcode example - false (sum exists only at non-leaf)",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum: 5, // 1+... no leaf path sums to 5 (1+2=3, 1+3=4)
			want:      false,
		},
		{
			name: "non-leaf node sum equals target must not short-circuit",
			// root=1 -> left child=2 (leaf). Path 1 alone is not root-to-leaf.
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			targetSum: 1,
			want:      false,
		},
		{
			name: "path only on right subtree",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			targetSum: 4, // 1+3
			want:      true,
		},
		{
			name: "negative values",
			root: &TreeNode{
				Val: -2,
				Right: &TreeNode{
					Val: -3,
				},
			},
			targetSum: -5,
			want:      true,
		},
		{
			name:      "nil root with nonzero target",
			root:      nil,
			targetSum: 5,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasPathSum(tt.root, tt.targetSum)
			if got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
