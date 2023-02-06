package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		// 2 -> 1
		temp := pre.Next
		pre.Next = pre.Next.Next
		sub := pre.Next.Next
		pre.Next.Next = temp
		temp.Next = sub

		pre = pre.Next.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
