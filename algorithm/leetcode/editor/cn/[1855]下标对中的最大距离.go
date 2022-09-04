package cn

//leetcode submit region begin(Prohibit modification and deletion)
func maxDistance(nums1 []int, nums2 []int) int {
	max := 0
	for i := 0; i < len(nums1); i++ {
		target := nums1[i]

		left := i
		right := len(nums2)-1
		// 找到大于target 的最右边的数
		for left <= right {
			mid := left + (right-left)/2
			if nums2[mid] < target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if left-1-i > max {
			max = left - 1 - i
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
