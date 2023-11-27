package cn

// leetcode submit region begin(Prohibit modification and deletion)
func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	dif := len(first) - len(second)
	if dif > 1 || dif < -1 {
		return false
	}
	// 操作次数
	op := 0

	i, j := 0, 0
	for i < len(first) || j < len(second) {
		if i == len(first) && j < len(second) {
			op++
			break
		} else if j == len(second) && i < len(first) {
			op++
			break
		}

		if first[i] == second[j] {
			i++
			j++
		} else {
			// 谁长谁往前走
			switch len(first) - len(second) {
			case 1:
				i++
				break
			case -1:
				j++
				break
			case 0:
				i++
				j++
				break
			}
			op++
			if op > 1 {
				break
			}
		}
	}

	return op < 2
}

//leetcode submit region end(Prohibit modification and deletion)
