package cn

//leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {
	l := 0
	r := len(letters)

	for l < r {
		m := l + (r-l)/2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	if l == len(letters) {
		return letters[0]
	}
	return letters[l]
}

//leetcode submit region end(Prohibit modification and deletion)
