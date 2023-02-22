package cn

//leetcode submit region begin(Prohibit modification and deletion)
func mySqrt(x int) int {
	l := 0
	r := x
	for l <= r {
		m := l + (r-l)/2
		if m*m <= x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l - 1
}

//leetcode submit region end(Prohibit modification and deletion)
