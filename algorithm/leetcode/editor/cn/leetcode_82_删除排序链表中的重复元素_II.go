package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 定义一个虚拟节点
	dummy := &ListNode{Val: 0, Next: head}
	tail := dummy
	for tail.Next != nil && tail.Next.Next != nil {
		// 判断两个节点是否相等
		if tail.Next.Val == tail.Next.Next.Val {
			// 相等的话，看一下后面的节点是否还有相等的
			// 保存一些节点的值
			v := tail.Next.Val
			// 直到找到一个空姐点或值不相等的节点
			for tail.Next != nil && tail.Next.Val == v {
				tail.Next = tail.Next.Next
			}
		} else {
			tail = tail.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
