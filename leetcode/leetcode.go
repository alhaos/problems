package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func inorderTraversal(root *TreeNode) []int {
	var values []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		values = append(values, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return values
}

func maxDepth(root *TreeNode) int {

	maxLevel := 0

	var dfs func(node *TreeNode, level int)

	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		maxLevel = max(maxLevel, level)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 1)

	return maxLevel
}

func minDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	depth := 1

	for len(q) > 0 {

		sz := len(q)

		for _ = range sz {

			currentNode := q[0]
			q = q[1:]

			if currentNode.Left == nil && currentNode.Right == nil {
				return depth
			}

			if currentNode.Left != nil {
				q = append(q, currentNode.Left)
			}

			if currentNode.Right != nil {
				q = append(q, currentNode.Right)
			}

		}

		depth++
	}
	return depth
}
