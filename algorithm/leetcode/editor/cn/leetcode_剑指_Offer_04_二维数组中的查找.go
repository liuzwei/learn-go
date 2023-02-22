package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	// 从左上角开始找，如果大了就左移，如果小了就下移
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1

	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
