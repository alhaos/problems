package pathSum

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// HasPathSum returns true if the tree has a root-to-leaf path
// such that adding up all the values along the path equals targetSum.
func HasPathSum(root *TreeNode, targetSum int) bool {

	var sum int

	var DFS func(node *TreeNode) bool

	DFS = func(node *TreeNode) bool {

		if node == nil {
			return false
		}

		sum += node.Val
		if DFS(node.Left) || DFS(node.Right) {
			return true
		}

		if sum == targetSum && node.Left == nil && node.Right == nil {
			return true
		}

		sum -= node.Val
		return false
	}

	return DFS(root)
}
