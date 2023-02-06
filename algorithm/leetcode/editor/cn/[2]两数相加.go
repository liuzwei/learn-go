package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟节点
	dummy := &ListNode{Val: 0}
	tail := dummy
	carry := 0
	// 直到所有的链遍历完为止
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		// 获取求和之后的值和进位值
		sum, carry = sum%10, sum/10
		// 将计算后的节点挂到新链上
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
