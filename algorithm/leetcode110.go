package algorithm

// 给定一个二叉树，判断它是否是高度平衡的二叉树。
//
//本题中，一棵高度平衡二叉树定义为：
//
//一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。

// IsBalanced 是否为平衡二叉树
func IsBalanced(root *TreeNode) bool {
	// 如果是一个平衡二叉树，则左子树的高度和右子树的高度 绝对值查 <= 1 ,并且左子树是平衡二叉树，右子树也是平衡二叉树
	if root == nil {
		return true
	} else {
		return treeAbs(height(root.Left), height(root.Right)) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
	}
}

// height 计算树的高度
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxNum(height(root.Left), height(root.Right)) + 1
}

//maxNum 取最大的
func maxNum(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func treeAbs(left int, right int) int {
	if left > right {
		return left - right
	}
	return right - left
}
