package util

// Reverse 反转列表
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 递归反转
	last := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

var subNode *ListNode

// ReverseN 反转列表的前N个元素
func ReverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 存储后续节点
		subNode = head.Next
		return head
	}
	last := ReverseN(head.Next, n-1)
	head.Next.Next = head
	// 将剩下的节点接到反转后的节点后面
	head.Next = subNode
	return last
}

// ReverseBetween 反转链条中的一部分
func ReverseBetween(head *ListNode, start int, end int) *ListNode {
	if start == 1 {
		return ReverseN(head, end)
	}
	head.Next = ReverseBetween(head.Next, start-1, end-1)
	return head
}
