package cn

//leetcode submit region begin(Prohibit modification and deletion)
func isPerfectSquare(num int) bool {

	l := 0
	r := num
	for l <= r {
		m := l + (r-l)/2
		if m*m <= num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if (l-1)*(l-1) == num {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
