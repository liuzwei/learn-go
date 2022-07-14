package algorithm

import "learn-go/util"

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

// LevelOrder 层序遍历
func LevelOrder(root *util.TreeNode) [][]int {
	// 利用队列的特性，并且记住每一层的元素数
	result := [][]int{}
	queue := []*util.TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		levelNum := len(queue)
		levelResult := []int{}
		for i := 0; i < levelNum; i++ {
			item := queue[0]
			queue = queue[1:]
			levelResult = append(levelResult, item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}

		result = append(result, levelResult)
	}

	return result
}
