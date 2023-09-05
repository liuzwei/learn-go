package cn

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		element := nums1[i]
		// 找到element在nums2的位置，并向后查找是否有更大元素
		optValue := -1
		optIndex := -1
		for j := 0; j < len(nums2); j++ {
			if element == nums2[j] {
				optIndex = j
			}
			if optIndex != -1 && j > optIndex && nums2[j] > element {
				optValue = nums2[j]
				break
			}
		}
		result = append(result, optValue)
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
