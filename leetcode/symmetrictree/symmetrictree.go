package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		a, b := queue[0], queue[1]
		queue = queue[2:]

		if a == nil && b == nil {
			continue
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}

		queue = append(queue, a.Left, b.Right, a.Right, b.Left)
	}

	return true
}
