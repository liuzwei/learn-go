package algorithm

import "learn-go/util"

/**
 * 二叉搜索树中第K小的元素
 *
 * 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
 */
func kthSmallest(root *util.TreeNode, k int) int {
	// 二叉树迭代中序遍历，取第K个最小元素
	stack := []*util.TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
