package cn

//leetcode submit region begin(Prohibit modification and deletion)
func specialArray(nums []int) int {

	return 0
}

func bubbleSort(nums []int) []int {
	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
