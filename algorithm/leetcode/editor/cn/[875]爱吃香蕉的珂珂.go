package cn

//leetcode submit region begin(Prohibit modification and deletion)
func minEatingSpeed(piles []int, h int) int {
	// sum
	sum := 0
	for i := 0; i < len(piles); i++ {
		sum += piles[i]
	}
	if sum <= h {
		return 1
	}
	// 先排序，取香蕉堆的最大值，二分查找
	sPiles := minEatingSpeedSort(piles)
	l := 0
	r := sPiles[len(sPiles)-1]
	for l < r {
		m := l + (r-l)/2
		if minEatingSpeedCount(piles, m) > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

// minEatingSpeedCount
func minEatingSpeedCount(piles []int, speed int) int {
	c := 0
	for i := 0; i < len(piles); i++ {
		c += piles[i] / speed
		if piles[i]%speed != 0 {
			c += 1
		}
	}
	return c
}

func minEatingSpeedSort(nums []int) []int {
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
