package cn

import "strconv"

// leetcode submit region begin(Prohibit modification and deletion)
func calPoints(operations []string) int {
	// 结果暂存
	var trans = make([]int, 0)
	// 翻译
	for _, p := range operations {
		switch p {
		case "D":
			trans = append(trans, trans[len(trans)-1]*2)
			break
		case "C":
			trans = trans[:len(trans)-1]
			break
		case "+":
			trans = append(trans, trans[len(trans)-1]+trans[len(trans)-2])
			break
		default:
			t, _ := strconv.Atoi(p)
			trans = append(trans, t)
			break
		}
	}
	// 求和
	sum := 0
	for i, _ := range trans {
		sum += trans[i]
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)
