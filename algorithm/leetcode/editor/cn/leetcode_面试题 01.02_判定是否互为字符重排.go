package cn

import (
	"sort"
	"strings"
)

// leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// 对s1进行排序
	sa1 := strings.Split(s1, "")
	sort.Strings(sa1)
	ns1 := strings.Join(sa1, "")

	// 对s2进行排序
	sa2 := strings.Split(s2, "")
	sort.Strings(sa2)
	ns2 := strings.Join(sa2, "")
	// 判断两个排序后的字符串是否相等
	return ns1 == ns2
}

//leetcode submit region end(Prohibit modification and deletion)
