package algorithm

import (
	"math"

	"learn-go/util"
)

func MergeKLists2(lists []*util.ListNode) *util.ListNode {
	if len(lists) == 0 {
		return nil
	}
	root := &util.ListNode{
		Val:  -10001,
		Next: nil,
	}

	for _, l := range lists {
		root = mergeTwo(root, l)
	}
	return root.Next
}

func mergeTwo(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	// 合并两个链表
	// 创建一个虚拟节点
	dummy := &util.ListNode{Val: -10002, Next: nil}
	head := dummy
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

func MergeKLists(lists []*util.ListNode) *util.ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 创建一个虚拟节点
	dump := &util.ListNode{
		Val:  -int(math.Pow(10, 4)) - 1,
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
