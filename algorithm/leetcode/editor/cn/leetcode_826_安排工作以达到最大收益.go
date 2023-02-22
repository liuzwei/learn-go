package cn

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	// 先排序，从小到大
	// 难度和收益是意义对应的，所以难度排序后要和收益排序一致
	sDiff, sProfit := maxProfitAssignmentSort(difficulty, profit)
	sum := 0
	// 遍历查找工人能干的最高收益
	for i := 0; i < len(worker); i++ {
		l := 0
		r := len(sDiff) - 1
		for l <= r {
			m := l + (r-l)/2
			if sDiff[m] <= worker[i] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if l-1 >= 0 {
			// 取能干的活的最大的收益那个
			max := 0
			for m := 0; m <= l-1; m++ {
				if sProfit[m] > max {
					max = sProfit[m]
				}
			}
			sum += max
		}
	}
	return sum
}

func maxProfitAssignmentSort(difficulty []int, profit []int) ([]int, []int) {
	for i := len(difficulty) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if difficulty[j] > difficulty[j+1] {
				difficulty[j], difficulty[j+1] = difficulty[j+1], difficulty[j]
				profit[j], profit[j+1] = profit[j+1], profit[j]
			}
		}
	}
	return difficulty, profit
}

//leetcode submit region end(Prohibit modification and deletion)
