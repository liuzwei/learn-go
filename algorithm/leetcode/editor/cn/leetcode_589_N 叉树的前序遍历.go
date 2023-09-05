package cn

type Node struct {
	Val      int
	Children []*Node
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var result []int

func preorder(root *Node) []int {
	result = make([]int, 0)
	// 前序遍历    根 -> 左 —> 右
	traverse(root)
	return result
}

func traverse(root *Node) {
	if root == nil {
		return
	}
	result = append(result, root.Val)
	for _, child := range root.Children {
		traverse(child)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
