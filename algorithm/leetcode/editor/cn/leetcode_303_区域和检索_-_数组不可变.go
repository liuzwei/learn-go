package cn

// leetcode submit region begin(Prohibit modification and deletion)
type NumArray struct {
	preNums []int
}

func Constructor3(nums []int) NumArray {
	nr := NumArray{
		preNums: make([]int, len(nums)+1),
	}
	for i := 1; i <= len(nums); i++ {
		nr.preNums[i] = nr.preNums[i-1] + nums[i-1]
	}
	return nr
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preNums[right+1] - this.preNums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
//leetcode submit region end(Prohibit modification and deletion)
