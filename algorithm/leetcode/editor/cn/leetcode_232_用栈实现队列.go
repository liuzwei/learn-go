package cn

// leetcode submit region begin(Prohibit modification and deletion)
type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	// 向stack1中压入数据
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	// 先判断stack2是否为空，如果是空的，则将stack1放到stack2中
	if len(this.stack2) == 0 {
		// 将stack1中的数据放到stack2中
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		// 将stack1清空
		this.stack1 = make([]int, 0)
	}
	// 从stack2中取出元素
	bk := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return bk
}

func (this *MyQueue) Peek() int {
	// 先判断stack2是否为空，如果是空的，则将stack1放到stack2中
	if len(this.stack2) == 0 {
		// 将stack1中的数据放到stack2中
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		// 将stack1清空
		this.stack1 = make([]int, 0)
	}
	return this.stack2[len(this.stack2)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
