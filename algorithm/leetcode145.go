package algorithm

import "learn-go/util"

func PostorderTraversal(root *util.TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	queue := []*CustomNode{
		{root, 0},
	}

	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if cur.root == nil {
			continue
		}
		if cur.color == 0 {
			queue = append(queue, &CustomNode{cur.root, 1})
			queue = append(queue, &CustomNode{cur.root.Right, 0})
			queue = append(queue, &CustomNode{cur.root.Left, 0})
		} else {
			result = append(result, cur.root.Val)
		}
	}

	return result
}
