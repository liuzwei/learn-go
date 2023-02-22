package cn

// leetcode submit region begin(Prohibit modification and deletion)
func fib(n int) int {
	if n == 0 {
		return 0
	}
	// 保存一个计算表
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	// 自底向上计算
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
