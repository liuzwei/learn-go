package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		}else {
			left = mid + 1
		}
	}
	return left
}

//leetcode submit region end(Prohibit modification and deletion)
