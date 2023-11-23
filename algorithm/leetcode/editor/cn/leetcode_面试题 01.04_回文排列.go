package cn

// leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	n := 0
	runeMap := make(map[rune]bool)
	for _, r := range s {
		if runeMap[r] {
			runeMap[r] = false
			n--
		} else {
			runeMap[r] = true
			n++
		}
	}

	return n < 2
}

//leetcode submit region end(Prohibit modification and deletion)
