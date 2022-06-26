package algorithm

func PivotIndex(nums []int) int {
	// 遍历数组，左侧元素之和-右侧元素之和=0
	for index, _ := range nums {
		leftSum := sum(nums[:index])
		rightSum := sum(nums[index+1:])
		if leftSum == rightSum {
			return index
		}
	}

	return -1
}

func sum(nums []int) int {
	sum := 0
	for _, item := range nums {
		sum += item
	}
	return sum
}
