package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	// 创建一个新链表
	n := reverse(createNew(head))
	// 对链表反转

	// 判断两个链表是否一样，一样则是回文链表，不一样则不是
	for head != nil {
		if head.Val != n.Val {
			return false
		}
		head = head.Next
		n = n.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func createNew(head *ListNode) *ListNode {
	prev := &ListNode{
		Val: head.Val,
	}
	cur := prev
	for head.Next != nil {
		cur.Next = &ListNode{
			Val: head.Next.Val,
		}
		cur = cur.Next
		head = head.Next
	}
	return prev
}

//leetcode submit region end(Prohibit modification and deletion)
