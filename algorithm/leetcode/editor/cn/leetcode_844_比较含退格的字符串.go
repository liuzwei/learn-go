package cn

// leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare(s string, t string) bool {
	i := len(s) - 1
	j := len(t) - 1

	iSkp, jSkp := 0, 0
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				iSkp++
				i--
			} else if iSkp > 0 {
				iSkp--
				i--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				jSkp++
				j--
			} else if jSkp > 0 {
				jSkp--
				j--
			} else {
				break
			}
		}

		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
