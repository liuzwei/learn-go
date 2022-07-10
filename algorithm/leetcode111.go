package algorithm

import "math"

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	minD := math.MaxInt32
	if root.Left != nil {
		minD = minNum(MinDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = minNum(MinDepth(root.Right), minD)
	}
	return minD + 1
}

func minNum(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
