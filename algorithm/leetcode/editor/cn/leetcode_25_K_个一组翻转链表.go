package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	// 获取节点的数量
	l := getLen(head)
	i := 0
	// 1 2 | 3 4 | 5 6
	for i < l/k {
		// 翻转
		start := i*k + 1
		end := start + k - 1
		head = reverseBetween(head, start, end)
		i += 1
	}

	return head
}

func getLen(head *ListNode) int {
	s := 0
	for head != nil {
		s += 1
		head = head.Next
	}
	return s
}

var subNode *ListNode

// reverseN 反转前N个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		subNode = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = subNode
	return last
}

// reverseBetween 反转链表中开始和结束之间的节点
func reverseBetween(head *ListNode, start int, end int) *ListNode {
	if start == 1 {
		return reverseN(head, end)
	}
	head.Next = reverseBetween(head.Next, start-1, end-1)
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
