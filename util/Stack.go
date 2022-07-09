package util

import "container/list"

type Stack struct {
	list *list.List
}

// NewStack 创建一个新的栈
func NewStack() *Stack {
	clist := list.New()
	return &Stack{clist}
}

// Push 添加元素
func (stack *Stack) Push(item interface{}) {
	stack.list.PushBack(item)
}

// Pop 弹出元素
func (stack *Stack) Pop() interface{} {
	element := stack.list.Back()
	if element != nil {
		stack.list.Remove(element)
		return element
	}
	return nil
}

// 栈是否为空
func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

// 栈的长度
func (stack *Stack) Len() int {
	return stack.list.Len()
}
