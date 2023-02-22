package cn

//leetcode submit region begin(Prohibit modification and deletion)
func countNegatives(grid [][]int) int {
	n := 0
	for i := 0; i < len(grid); i++ {
		temp := grid[i]
		if temp[0] < 0 {
			n += len(temp)
			continue
		}
		left := 0
		right := len(temp)
		for left < right {
			mid := left + (right-left)/2
			if temp[mid] >= 0 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		n += len(temp) - left
	}

	return n
}

//leetcode submit region end(Prohibit modification and deletion)
