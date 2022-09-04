package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	l := 0
	r := len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m]-m-1 < k {
			l = m + 1
		} else {
			r = m
		}
	}
	// (arr[l-1] - (l - 1) - 1) 计算的是 l-1位置缺失的个数
	// k - 缺失的个数 计算的是还剩下几个缺失的数
	// 剩下缺失的数 + arr[l-1] 即为缺失的第K个数
	return k - (arr[l-1] - (l - 1) - 1) + arr[l-1]
}

//leetcode submit region end(Prohibit modification and deletion)
