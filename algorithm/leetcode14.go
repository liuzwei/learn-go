package algorithm

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串 ""。
//
//

func LongestCommonPrefix(strs []string) string {
	// 纵向遍历
	l := len(strs[0])
	result := ""
	for i := 0; i < l; i++ {
		temp := string(strs[0][i])
		for _, s := range strs {
			if len(s) < i+1 || temp != string(s[i]) {
				return result
			}
		}
		result = result + temp
	}
	return result
}
