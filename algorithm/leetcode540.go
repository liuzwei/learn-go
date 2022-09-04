package algorithm

func SingleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == nums[mid+1] {
			// 看右边是否是偶数个
			if (right-mid+1)%2 == 0 {
				// 右边不缺数，在左边
				right = mid - 1
			} else {
				// 左边不缺数，在右边
				left = mid + 2
			}
		} else if nums[mid] == nums[mid-1] {
			// 看右边是否是偶数个
			if (mid-left+1)%2 == 0 {
				// 左边不缺数，在右边
				left = mid + 1
			} else {
				// 左边不缺数，在右边
				right = mid - 2
			}
		} else {
			left = mid
		}
	}
	return nums[left]
}
