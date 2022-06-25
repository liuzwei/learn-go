package algorithm

func RemoveElement(nums []int, val int) int {
	// 双指针，一个从左开始，一个从右开始，左指针遍历到val时，将右指针的赋值给左指针，直到两个指针相遇
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
