package algorithm

func ConvertToTitle(columnNumber int) string {
	// 26进制
	var result string
	for columnNumber > 0 {
		columnNumber--
		result = string(rune(columnNumber%26+'A')) + result
		columnNumber = columnNumber / 26
	}
	return result
}
