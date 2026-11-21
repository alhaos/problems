package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	var l int

	var fn func(node *TreeNode, level int)

	fn = func(node *TreeNode, level int) {

		if node == nil {
			return
		}

		l = max(l, level)

		fn(node.Left, level+1)
		fn(node.Right, level+1)

	}

	fn(root, 1)

	return l
}

func maxDepth_(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
