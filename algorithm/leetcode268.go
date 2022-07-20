package algorithm

import "learn-go/util"

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
//

// MissingNumber
func MissingNumber(nums []int) int {
	// 1. sort
	nums = util.BubbleSort(nums)
	// 2. iterate and calculate difference
	cur := -1
	for _, i := range nums {
		if i-cur == 2 {
			return i - 1
		}
		cur = i
	}
	return len(nums)
}
