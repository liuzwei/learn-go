package cn

//leetcode submit region begin(Prohibit modification and deletion)
func peakIndexInMountainArray(arr []int) int {

	left := 0
	right := len(arr) - 1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

//leetcode submit region end(Prohibit modification and deletion)
