package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var afterResult []int

func postorder(root *Node) []int {
	afterResult = make([]int, 0)
	afterTraverse(root)
	return afterResult
}

func afterTraverse(root *Node) {
	if root == nil {
		return
	}
	for _, n := range root.Children {
		afterTraverse(n)
	}
	afterResult = append(afterResult, root.Val)
}

//leetcode submit region end(Prohibit modification and deletion)
