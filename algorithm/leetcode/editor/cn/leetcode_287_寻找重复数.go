package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findDuplicate(nums []int) int {
	// 先排序，后二分查找
	l := 0
	r := len(nums) - 1
	ans := -1
	for l <= r {
		m := l + (r-l)/2
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= m {
				cnt++
			}
		}
		if cnt <= m {
			l = m + 1
		} else {
			r = m - 1
			ans = m
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
