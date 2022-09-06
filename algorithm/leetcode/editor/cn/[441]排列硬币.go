package cn

//leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	l := 0
	r := n
	for l < r {
		m := l + (r-l)/2
		if (m*(m+1))/2 <= n {
			l = m + 1
		}else {
			r = m
		}
	}
	return l - 1
}

//leetcode submit region end(Prohibit modification and deletion)
