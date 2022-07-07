package algorithm

// FinalValueAfterOperations 执行操作后的变量值
// @param operations
func FinalValueAfterOperations(operations []string) int {
	sum := 0
	for _, op := range operations {
		switch op {
		case "++X", "X++":
			sum++
		case "--X", "X--":
			sum--
		}
	}
	return sum
}
