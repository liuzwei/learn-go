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
	root := &ListNode{
		Val:  -10001,
		Next: nil,
	}

	for _, l := range lists {
		root = mergeTwo(root, l)
	}
	return root.Next
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	// 合并两个链表
	// 创建一个虚拟节点
	dummy := &ListNode{Val: -10002, Next: nil}
	// 新生成的链的头节点
	head := dummy
	// 要合并的两条链的双指针
	p1 := l1
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
		head = head.Next
	}
	if p1 == nil {
		head.Next = p2
	} else {
		head.Next = p1
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
