package cn

//leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := l + (r-l)/2
		// 如果中间值大于右边值，说明峰值在左边，缩小右边界
		if arr[m] > arr[m+1] {
			r = m
		} else {
			// 说明峰值在右边，缩小左边界
			l = m + 1
		}
	}
	return l
}

//leetcode submit region end(Prohibit modification and deletion)
