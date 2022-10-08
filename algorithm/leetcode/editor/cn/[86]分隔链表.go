package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{
		Val: -1,
	}
	big := &ListNode{
		Val: -1,
	}
	p1 := small
	p2 := big
	p := head
	for p != nil {
		if p.Val >= x {
			p2.Next = p
			p2 = p2.Next
		} else {
			p1.Next = p
			p1 = p1.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = big.Next
	return small.Next
}

//leetcode submit region end(Prohibit modification and deletion)
