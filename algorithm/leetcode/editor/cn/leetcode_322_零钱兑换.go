package cn

import "math"

// leetcode submit region begin(Prohibit modification and deletion)

var memo []int

func coinChange(coins []int, amount int) int {
	memo = make([]int, amount+1)
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != 0 {
		return memo[amount]
	}

	res := math.MaxInt
	for _, coin := range coins {
		// 求子问题
		subProgress := dp(coins, amount-coin)
		if subProgress < 0 {
			continue
		}
		res = getMin(res, subProgress+1)
	}
	if res == math.MaxInt {
		memo[amount] = -1
	}
	memo[amount] = res
	return memo[amount]
}

// getMin 获取最小值
func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
