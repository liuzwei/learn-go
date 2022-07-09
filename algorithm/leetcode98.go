package algorithm

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
//有效 二叉搜索树定义如下：
//
//节点的左子树只包含 小于 当前节点的数。
//节点的右子树只包含 大于 当前节点的数。
//所有左子树和右子树自身必须也是二叉搜索树。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/validate-binary-search-tree
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：
// 中序遍历二叉树，左--根--右  为有序，当当前节点<小于前一个节点则认为搜索二叉树无效

// IsValidBST 是否有效二叉搜索树
func IsValidBST(root *TreeNode) bool {
	// 颜色标记法
	type ColorNode struct {
		color int
		node  *TreeNode
	}
	stack := []ColorNode{
		{0, root},
	}

	result := []int{}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.node == nil {
			continue
		}
		if current.color == 0 {
			stack = append(stack, ColorNode{0, current.node.Right})
			stack = append(stack, ColorNode{1, current.node})
			stack = append(stack, ColorNode{0, current.node.Left})
		} else {
			if len(result) > 0 && current.node.Val <= result[len(result)-1] {
				return false
			}
			result = append(result, current.node.Val)
		}
	}
	return true
}
