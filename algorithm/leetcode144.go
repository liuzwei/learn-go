package algorithm

import "learn-go/util"

// 给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
func preorderTraversal(root *util.TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []CustomNode{
		{root, 0},
	}

	for len(stack) > 0 {
		cn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if cn.root == nil {
			continue
		}

		if cn.color == 0 {
			stack = append(stack, CustomNode{cn.root.Right, 0})
			stack = append(stack, CustomNode{cn.root.Left, 0})
			stack = append(stack, CustomNode{cn.root, 1})
		} else {
			result = append(result, cn.root.Val)
		}
	}

	return result
}

type CustomNode struct {
	root  *util.TreeNode
	color int
}
