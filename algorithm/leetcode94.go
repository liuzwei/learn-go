package algorithm

import "learn-go/util"

func InorderTraversal(root *util.TreeNode) []int {
	// 递归
	result := []int{}
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	if root.Left != nil {
		result = append(result, InorderTraversal(root.Left)...)
	}

	result = append(result, root.Val)

	if root.Right != nil {
		result = append(result, InorderTraversal(root.Right)...)
	}

	return result
}

// InorderTraversalColorTag 中序遍历，颜色标记法
func InorderTraversalColorTag(root *util.TreeNode) []int {
	type CustomNode struct {
		// 0白色 1灰色
		color int
		node  *util.TreeNode
	}
	stack := []CustomNode{
		{
			0, root,
		},
	}
	result := []int{}
	for len(stack) > 0 {
		currNode := stack[len(stack)-1]
		if currNode.node == nil {
			continue
		}
		stack = stack[:len(stack)-1]
		if currNode.color == 0 {
			stack = append(stack, CustomNode{0, currNode.node.Right})
			stack = append(stack, CustomNode{1, currNode.node})
			stack = append(stack, CustomNode{0, currNode.node.Left})
		} else {
			result = append(result, currNode.node.Val)
		}
	}
	return result
}
