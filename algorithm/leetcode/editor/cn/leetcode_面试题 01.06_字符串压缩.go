package cn

import (
	"strconv"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	// 遍历压缩
	cnt := 0
	currentStr := S[0]
	var pressStr strings.Builder

	for i := 0; i < len(S); i++ {
		if S[i] != currentStr {
			pressStr.WriteString(string(currentStr))
			pressStr.WriteString(strconv.Itoa(cnt))
			currentStr = S[i]
			cnt = 1
		} else {
			cnt++
		}
	}
	pressStr.WriteString(string(currentStr))
	pressStr.WriteString(strconv.Itoa(cnt))

	if len(pressStr.String()) >= len(S) {
		return S
	}

	return pressStr.String()
}

//leetcode submit region end(Prohibit modification and deletion)
