package cn

import "strings"

// leetcode submit region begin(Prohibit modification and deletion)
func replaceSpaces(S string, length int) string {
	ns := S[:length]
	return strings.Replace(ns, " ", "%20", -1)
}

//leetcode submit region end(Prohibit modification and deletion)
