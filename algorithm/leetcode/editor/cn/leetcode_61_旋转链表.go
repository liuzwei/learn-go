package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	// 先确定链条的长度
	l := getNodeLength(head)
	// 取余找到旋转的节点
	m := k % l
	if m == 0 {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	tail := dummy
	var mid *ListNode
	for i := 0; i < l; i++ {
		tail = tail.Next
		if i == (l - m - 1) {
			mid = tail
		}
	}
	dummy.Next = mid.Next
	mid.Next = nil
	tail.Next = head
	return dummy.Next
}

func getNodeLength(head *ListNode) int {
	l := 0
	for head != nil {
		l += 1
		head = head.Next
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
