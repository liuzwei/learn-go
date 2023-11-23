package cn

// leetcode submit region begin(Prohibit modification and deletion)
func isUnique(astr string) bool {
	// 双指针遍历
	for i := 0; i < len(astr); i++ {
		for j := i + 1; j < len(astr); j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
