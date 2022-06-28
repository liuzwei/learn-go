package algorithm

//给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

func FindDiagonalOrder(mat [][]int) []int {
	m, n := len(mat)-1, len(mat[0])-1
	x, y := 0, 0
	points := [][]int{}
	for {
		points = append(points, []int{x, y})
		if x == m && y == n {
			break
		}
		isOdd := (x + y) % 2
		if isOdd == 0 {
			// 偶数上升
			x = x - 1
			y = y + 1
			if x < 0 || y > n {
				x = x + 1
			}
			// 校验，x,y是否合法

		} else {
			//奇数下降
			x = x + 1
			y = y - 1
			if y < 0 || x > m {
				y = y + 1
			}
		}
	}
	result := []int{}
	// 根据坐标取值
	for _, point := range points {
		if point[0] < 0 || point[1] < 0 || point[0] > m || point[1] > n {
			continue
		}
		x, y = point[0], point[1]
		result = append(result, mat[point[0]][point[1]])
	}
	return result
}

func FindDiagonalOrder2(mat [][]int) []int {
	m, n := len(mat)-1, len(mat[0])-1
	x, y := 0, 0
	points := [][]int{}
	for {
		points = append(points, []int{x, y})
		if x == m && y == n {
			break
		}
		isOdd := (x + y) % 2
		if isOdd == 0 {
			// 偶数上升
			x = x - 1
			y = y + 1
			if x < 0 || y > n {
				x = x + 1
			}
		} else {
			//奇数下降
			x = x + 1
			y = y - 1
			if y < 0 || x > m {
				y = y + 1
			}
		}
	}
	result := []int{}
	// 根据坐标取值
	for _, point := range points {
		if point[0] < 0 || point[1] < 0 || point[0] > m || point[1] > n {
			continue
		}
		x, y = point[0], point[1]
		result = append(result, mat[point[0]][point[1]])
	}
	return result
}
