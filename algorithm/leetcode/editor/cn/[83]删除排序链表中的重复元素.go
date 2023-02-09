package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	tail := dummy.Next
	for tail != nil {
		if tail.Next != nil && tail.Val == tail.Next.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
