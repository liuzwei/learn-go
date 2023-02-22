package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	// 迭代法删除节点

	// 创建一个虚拟节点
	dummy := &ListNode{Val: 0, Next: head}
	temp := dummy
	for temp.Next != nil {
		// 判断节点是否等于val
		if temp.Next.Val == val {
			// 说明该节点需要删除
			temp.Next = temp.Next.Next
		}else {
			// 该节点不用删除，移动到下一节点
			temp = temp.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
