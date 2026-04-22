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

/*
 * 56. Merge Intervals
 Medium
 Topics
 premium lock icon
 Companies
 Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.



 Example 1:

 Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 Output: [[1,6],[8,10],[15,18]]
 Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
 Example 2:

 Input: intervals = [[1,4],[4,5]]
 Output: [[1,5]]
 Explanation: Intervals [1,4] and [4,5] are considered overlapping.
 Example 3:

 Input: intervals = [[4,7],[1,4]]
 Output: [[1,7]]
 Explanation: Intervals [1,4] and [4,7] are considered overlapping.


 Constraints:

 1 <= intervals.length <= 104
 intervals[i].length == 2
 0 <= starti <= endi <= 104
*/

/*

func Contains(r io.Reader, seq []byte) (bool, error) {

	buffer := make([]byte, len(seq))

	err := r.Read(buffer)

	return false, nil
}

*/

func merge(intervals [][]int) [][]int {
	merged, didMerge := mergeOnce(intervals)
	if !didMerge {
		return merged
	}
	return merge(merged)
}

// Один проход: пытаемся смержить первую найденную пару
func mergeOnce(intervals [][]int) ([][]int, bool) {
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			a, b := intervals[i], intervals[j]

			// Проверяем пересечение или касание
			if a[0] <= b[1] && b[0] <= a[1] {
				merged := []int{min(a[0], b[0]), max(a[1], b[1])}

				// Убираем оба, добавляем смерженный
				result := [][]int{merged}
				for k, iv := range intervals {
					if k != i && k != j {
						result = append(result, iv)
					}
				}
				return result, true
			}
		}
	}
	return intervals, false
}

// generate 118. Pascal's Triangle
func generate(n int) [][]int {
	var result [][]int
	for rowNumber := range n {
		switch rowNumber {
		case 0:
			result = append(result, []int{1})
		case 1:
			result = append(result, []int{1, 1})
		default:
			n := []int{1}
			for i := 0; i < len(result[rowNumber-1])-1; i++ {
				n = append(n, result[rowNumber-1][i]+result[rowNumber-1][i+1])
			}
			n = append(n, 1)
			result = append(result, n)
		}
	}
	return result
}
