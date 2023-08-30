package cn

// leetcode submit region begin(Prohibit modification and deletion)
type MyStack struct {
	q          []int
	topElement int
}

func Constructor() MyStack {
	return MyStack{q: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.q = append(this.q, x)
	this.topElement = x
}

func (this *MyStack) Pop() int {
	/**
	解法2：
	*/
	// 如何开始和结束节点相等，说明只有一个元素
	if len(this.q) == 1 {
		bk := this.q[0]
		this.q = this.q[1:]
		return bk
	}
	// 尾元素位置
	end := len(this.q) - 1
	// 根据栈的规则，先进后出，则尾元素是要出栈的第一个元素
	bk := this.q[end]
	// 得到新的栈
	this.q = this.q[0:end]
	// 尾元素出栈后，则前一元素变为尾元素
	this.topElement = this.q[len(this.q)-1]
	return bk
	/** 解法1：
	size := len(this.q)
	for size > 1 {
		this.q = append(this.q, this.q[0])
		this.q = this.q[1:]
		size--
	}
	this.topElement = this.q[len(this.q)-1]

	result := this.q[0]
	this.q = this.q[1:]
	return result
	*/
}

func (this *MyStack) Top() int {
	return this.topElement
}

func (this *MyStack) Empty() bool {
	return len(this.q) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
//leetcode submit region end(Prohibit modification and deletion)
