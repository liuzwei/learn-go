package cn

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n == 0 {
		return 0
	}
	// 保存一个计算表
	//dp := make([]int, n+1)
	dp0 := 0
	dp1 := 1
	// 自底向上计算
	for i := 2; i <= n; i++ {
		dpi := dp0 + dp1
		dp0, dp1 = dp1, dpi
	}
	return dp1
}

//leetcode submit region end(Prohibit modification and deletion)
