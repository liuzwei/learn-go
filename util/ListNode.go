package util

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ListNodeInstance(nums []int) *ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	listNode := ListNode{Val: nums[0], Next: nil}
	currNode := &listNode
	for i := 1; i < len(nums); i++ {
		tempNode := ListNode{Val: nums[i], Next: nil}
		currNode.Next = &tempNode
		currNode = &tempNode
	}
	return &listNode
}

func (head *ListNode) toString() string {
	var builder = strings.Builder{}
	builder.WriteString("[")
	for head != nil {
		builder.WriteString(strconv.Itoa(head.Val))
		builder.WriteString(" ")
		head = head.Next
	}
	builder.WriteString("]")
	return builder.String()
}
