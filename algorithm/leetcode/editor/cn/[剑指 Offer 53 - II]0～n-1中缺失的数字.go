package cn

//leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		// 如果大于中间值，说明左边缺元素
		if nums[mid] > mid {
			// 缩小右边界
			right = mid - 1
		} else {
			// 元素和下标对应，右边缺元素
			left = mid + 1
		}
	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)
