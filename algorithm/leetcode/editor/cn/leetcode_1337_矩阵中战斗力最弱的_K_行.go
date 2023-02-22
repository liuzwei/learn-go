package cn

//leetcode submit region begin(Prohibit modification and deletion)
func kWeakestRows(mat [][]int, k int) []int {
	mas := make(map[int]int)
	for i := 0; i < len(mat); i++ {
		temp := mat[i]
		// 找1的右边界
		left := 0
		right := len(temp) - 1
		for left <= right {
			mid := left + (right-left)/2
			if temp[mid] > 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left-1 >= 0 {
			mas[i] = left
		} else {
			mas[i] = 0
		}
	}
	// 排序 TODO
	rets := []int{}

	for i := 0; i < k; i++ {
		//找最小的
		min := 101
		tempIndex := 0
		for j:=0; j<len(mat); j++ {
			temp,ok := mas[j]
			if ok && temp < min {
				min = temp
				tempIndex = j
			}
		}
		rets = append(rets, tempIndex)
		delete(mas, tempIndex)
	}
	return rets
}

//leetcode submit region end(Prohibit modification and deletion)
