package cn

//leetcode submit region begin(Prohibit modification and deletion)
func intersect(nums1 []int, nums2 []int) []int {
	// 排序，双指针，谁小谁走，相等一块走
	sort1 := intersectSort(nums1)
	sort2 := intersectSort(nums2)

	result := []int{}
	p1 := 0
	p2 := 0
	for p1 <= len(sort1)-1 && p2 <= len(sort2)-1 {
		if sort1[p1] < sort2[p2] {
			p1 += 1
		} else if sort1[p1] > sort2[p2] {
			p2 += 1
		} else {
			result = append(result, nums1[p1])
			p1 += 1
			p2 += 1
		}
	}
	return result
}

// 冒泡排序
func intersectSort(nums []int) []int {

	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

//leetcode submit region end(Prohibit modification and deletion)
