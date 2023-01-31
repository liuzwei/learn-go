package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 创建一个虚拟节点
	dump := &ListNode{
		Val:  -10001,
		Next: nil,
	}
	for _, l := range lists {
		// 取两个指针，一个是主指针，一个是动态的每个链条的指针
		p := dump
		// 每个链条的指针
		lp := l
		// 比较链条指针和主指针的值哪个值大
		for lp != nil {
			if (lp.Val >= p.Val && p.Next == nil) || (lp.Val >= p.Val && lp.Val <= p.Next.Val) {
				// 将lp的Node挂到主链
				current := lp
				lp = lp.Next
				// 断开主链，添加lp节点
				subNode := p.Next
				p.Next = current
				p = p.Next
				p.Next = subNode
			} else {
				p = p.Next
			}
		}
	}

	return dump.Next
}

//leetcode submit region end(Prohibit modification and deletion)
