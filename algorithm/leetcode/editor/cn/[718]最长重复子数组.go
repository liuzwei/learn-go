package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findLength(nums1 []int, nums2 []int) int {
	l := 0
	for i := 0; i < len(nums1); i++ {
		for j := i; j < len(nums1); j++ {
			common := nums1[i : j+1]
			for m := 0; m < len(nums2); m++ {
				e := m + len(common)
				if e > len(nums2) {
					e = len(nums2) - 1
				}
				if compareArray(common, nums2[m:e]) {
					if len(common) > l {
						l = len(common)
					}
					break
				}
			}
		}
	}
	return l
}

func compareArray(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
