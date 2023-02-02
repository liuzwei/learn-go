package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 递归反转
	if head == nil || head.Next == nil {
		return head
	}
	// 递归寻找最后一个节点
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

//leetcode submit region end(Prohibit modification and deletion)
