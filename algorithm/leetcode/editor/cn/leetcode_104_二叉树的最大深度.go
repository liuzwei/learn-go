package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	if leftMax > rightMax {
		return leftMax + 1
	} else {
		return rightMax + 1
	}
}

//leetcode submit region end(Prohibit modification and deletion)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
