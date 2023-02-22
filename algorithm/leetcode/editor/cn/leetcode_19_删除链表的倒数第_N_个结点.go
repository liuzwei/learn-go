package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	current := dummy
	l := getLength(head)
	// 要删除的节点就是 第(l-n+1)个
	for i := 0; i < (l - n); i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0
	// 遍历得到链条的长度
	for head != nil {
		head = head.Next
		length += 1
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)
