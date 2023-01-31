package algorithm

import (
	"math"

	"learn-go/util"
)

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
