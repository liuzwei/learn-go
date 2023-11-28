package cn

// leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int) {
	l := len(matrix)
	x := l / 2
	y := (l + 1) / 2

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			tmp := matrix[i][j]
			// 旋转
			matrix[i][j] = matrix[l-j-1][i]
			matrix[l-j-1][i] = matrix[l-i-1][l-j-1]
			matrix[l-i-1][l-j-1] = matrix[j][l-i-1]
			matrix[j][l-i-1] = tmp
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
